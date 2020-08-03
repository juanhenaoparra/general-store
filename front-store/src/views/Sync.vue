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

        <b-button size="lg" variant="success" @click="syncData2()">Sync!</b-button>
      </div>
    </div>
    <div class="row">
      <div v-if="syncResponse">
        <code>{{syncResponse}}</code>
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
    }
  },
  computed: {
    ...mapState(['syncResponse']),
  },
  methods: {
    ...mapActions(['syncData']),
    syncData2: function () {
      if(this.syncThisDate){
        return this.syncData(new Date(this.syncThisDate));
      }

      this.syncData();
    }
  }
}
</script>
