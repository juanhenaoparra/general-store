<template>
<div class="container my-4" v-if="currentProfile">
  <b-jumbotron class="text-center col-md-7 mx-auto">
    <template v-slot:header>{{ currentProfile.name }}</template>

    <template v-slot:lead>
          ID: <u>{{currentProfile.id}}</u> <br>
          Age: <u>{{currentProfile.age}}</u>
    </template>

  </b-jumbotron>

  <div class="row text-center">
    <div class="col-md-10 mx-auto">
      <h3>We highly recommend based in your transactions history this products</h3>
      <p class="recommends p-3" v-if="currentProfile.recommends.length > 0">
        {{joinTextProducts(currentProfile.recommends)}}
      </p>
    </div>
  </div>

  <div class="row">
    <div class="col-md-6 text-center">
      <h3>Transactions</h3>
      <b-list-group>
        <b-list-group-item v-for="(transaction, i) in currentProfile['~by_buyer']" :key="i">
          ID:  <strong>{{transaction.id}}</strong> <br>
          it costs: ${{getTotalPrice(transaction)}} <br>
          at {{timeToHuman(transaction.time.timestamp)}} <br>
          <div>
            <b-list-group>
              <b-list-group-item v-for="(prod, i) in transaction.have_products" :key="i" class="text-right">
                <small>
                  {{prod.name}} | <strong>${{prod.price}}</strong>
                </small>
              </b-list-group-item>
            </b-list-group>
          </div>
        </b-list-group-item>
      </b-list-group>
    </div>
    <div class="col-md-6 text-center">
      <h3>Other users using this IP</h3>

      <b-list-group>
        <b-list-group-item v-for="(ip, i) in getIpsByCurrent" :key="i">
          <strong>IP Address: {{ip.address}}</strong> <br>
          <small>
            {{ joinText(ip.buyers) }}
          </small>
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</div>
</template>

<script>
import {mapState, mapGetters, mapActions} from 'vuex';

export default {
  name: 'Profile',
  data(){
    return {
    }
  },
  computed: {
    ...mapState(['currentProfile']),
    ...mapGetters(['getIpsByCurrent']),
  },
  methods: {
    ...mapActions(['seeMyProfile', 'getRecommendedProducts']),
    timeToHuman: function(unix_timestamp) {
      return new Date(unix_timestamp * 1000).toLocaleString();
    },
    getTotalPrice(transaction) {
      let total = transaction.have_products.reduce((acc, itemPrice) => {
        return acc + itemPrice.price;
      }, 0)

      return total;
    },
    joinText(buyers) {
      return buyers.map(b => {
        return `Name: ${b.name} ID: ${b.id}`;
      }).join(', ');
    },
    joinTextProducts(recommends) {
      return recommends.map(r => {
        return r.name;
      }).join(', ');
    },
  },
  created(){
    this.seeMyProfile([this.$route.params.id, 30, 0]).then(() => {
      this.getRecommendedProducts();
    });
  },
  destroyed() {
    this.deleteCurrentProfile();
  }
}
</script>

<style scoped>
  .recommends {
    color: gray;
  }
</style>