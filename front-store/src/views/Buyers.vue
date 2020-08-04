<template>
<div class="container mt-4 text-center">
  <div class="row my-5">
    <div class="col-md-12">
      <h1>View All Buyers</h1>
    </div>
  </div>
  <div class="row my-3">
    <b-paginate>
      <div id="table-page">
            <b-pagination
                v-model="currentPage"
                :total-rows="rows+1"
                :per-page="perPage"
                @input="prepareNextPage()"
                aria-controls="users-table"
            ></b-pagination>
        </div>
    </b-paginate>
  </div>
  <div class="row">
    <b-table
      id="buyers-table"
      striped hover
      :fields="fields"
      :items="allBuyers"
      :perPage="perPage"
      :currentPage="currentPage"
    >
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
      perPage: 50,
      currentPage: 1,
      fields: [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Name' },
        { key: 'age', label: 'Age' },
        { key: 'actions', label: 'Actions' },
      ]
    }
  },
  computed: {
    ...mapState(['allBuyers']),
    rows: function(){
      return this.allBuyers.length;
    },
  },
  methods: {
    ...mapActions(['getAllBuyers']),
    prepareNextPage(){
      if(this.allBuyers.length < this.perPage*this.currentPage){
        this.getAllBuyers([this.perPage, this.allBuyers.length]);
      }
    },
  },
  created(){
    if(this.allBuyers.length == 0){
      this.getAllBuyers([this.perPage, 0])
    }
  }
}
</script>

<style>

</style>