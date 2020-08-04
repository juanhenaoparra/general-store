<template>
  <div class="sync container my-4 text-center">

    <div class="row">
      <div class="col-md-12">
        <h1>Sync Up All Your Store Data</h1>
      </div>
      <div class="col-md-12">
        <div class="row">
          <div class="col-md-4 mx-auto my-5">
            <b-form-group label="Select a date to sync: " label-for="input-title">
                <b-form-input
                    id="input-title"
                    v-model="syncThisDate"
                    type="date"
                    required
                ></b-form-input>
            </b-form-group>
          </div>
        </div>

        <b-button size="lg" variant="success" @click="handleSyncButton()" :disabled="charging">Sync!</b-button>
      </div>
    </div>
    <div class="row">
      <div v-if="syncResponse">
        <div v-if="typeof syncResponse == 'string'">
          <p class="error">
            {{syncResponse}}
          </p>
        </div>
        <div v-else>
          Synchronized Data: <br>
          <ul class="text-left">
            <li>Date: {{timeToHuman(syncResponse.date)}}</li>
            <li>Buyers: {{syncResponse.buyers}}</li>
            <li>Products: {{syncResponse.products}}</li>
            <li>Ips: {{syncResponse.ips}}</li>
            <li>Devices: {{syncResponse.devices}}</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapState, mapActions} from 'vuex';

export default {
  name: 'SyncView',
  data(){
    return {
      syncThisDate: undefined,
      charging: undefined,
    }
  },
  computed: {
    ...mapState(['syncResponse']),
  },
  methods: {
    ...mapActions(['syncData']),
    handleSyncButton: function () {
      this.charging = true;
      if(this.syncThisDate){
        return this.syncData(new Date(this.syncThisDate)).then(() => {
          this.charging = undefined;
        });
      }

      this.syncData().then(() => {
        this.charging = undefined;
      });
    },
    timeToHuman: function(unix_timestamp) {
      return new Date(unix_timestamp * 1000).toLocaleString();
    },
  }
}
</script>
<style scoped>
  .error {
    color: red;
  }
</style>
