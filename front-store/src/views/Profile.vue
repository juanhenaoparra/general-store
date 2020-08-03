<template>
<div class="container my-4">
  <b-jumbotron>
    <template v-slot:header>{{ currentProfile.name }}</template>

    <template v-slot:lead>
          ID: <u>{{currentProfile.id}}</u> <br>
          Age: <u>{{currentProfile.age}}</u>
    </template>


  </b-jumbotron>

  <div class="row">
    <div class="col-md-6">
      <h3>Transactions</h3>
      <b-list-group>
        <b-list-group-item v-for="(transaction, i) in currentProfile['~by_buyer']" :key="i">
          ID:  <strong>{{transaction.id}}</strong> <br>
          it costs: ${{getTotalPrice(transaction)}} <br>
          at {{timeToHuman(transaction.time.timestamp)}} <br>
          <div>
            <b-list-group>
              <b-list-group-item v-for="(prod, i) in transaction.have_products" :key="i">
                <small>
                  {{prod.name}} | <strong>${{prod.price}}</strong>
                </small>
              </b-list-group-item>
            </b-list-group>
          </div>
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</div>
</template>

<script>
import {mapState, mapActions} from 'vuex';

export default {
  name: 'Profile',
  data(){
    return {
    }
  },
  computed: {
    ...mapState(['currentProfile']),
  },
  methods: {
    ...mapActions(['seeMyProfile']),
    timeToHuman: function(unix_timestamp) {
      return new Date(unix_timestamp * 1000).toLocaleString();
    },
    getTotalPrice(transaction) {
      let total = transaction.have_products.reduce((acc, itemPrice) => {
        return acc + itemPrice.price;
      }, 0)

      return total;
    }
  },
  created(){
    this.seeMyProfile([this.$route.params.id, 30, 0]);
  }
}
</script>

<style>

</style>