<template>
<div class="container mt-4 text-center">
  <div class="row">
    <div class="col-md-12">
      <h1>View All Buyers</h1>
    </div>
  </div>
  <div class="row">
    <b-table id="buyers-table" striped hover :fields="fields" :items="allBuyers">
      <template v-slot:cell(actions)="data">
        <b-button
              :to="{name: 'Profile', params: {id: data.item.id}}"
              size="sm"
              variant="info"
              class="m-1"
          > Profile
          </b-button>
      </template>
    </b-table>
  </div>
</div>
</template>

<script>
import {mapState, mapActions} from 'vuex';

export default {
  name: 'BuyersView',
  data() {
    return {
      first: 50,
      fields: [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Name' },
        { key: 'age', label: 'Age' },
        { key: 'actions', label: 'Actions' },
      ]
    }
  },
  computed: {
    ...mapState(['allBuyers'])
  },
  methods: {
    ...mapActions(['getAllBuyers']),
  },
  created(){
    this.getAllBuyers([this.first, 0])
  }
}
</script>

<style>

</style>