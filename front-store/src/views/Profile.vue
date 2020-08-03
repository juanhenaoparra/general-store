<template>
<div class="container my-4">
  <b-jumbotron>
    <template v-slot:header>{{ currentProfile.name }}</template>

    <template v-slot:lead>
          ID: <u>{{currentProfile.id}}</u> <br>
          Age: <u>{{currentProfile.age}}</u>
    </template>


  </b-jumbotron>
  <b-list-group>
    <b-list-group-item v-for="(transaction, i) in currentProfile['~by_buyer']" :key="i">
      {{transaction.id}} - {{timeToHuman(transaction.time.timestamp)}}
    </b-list-group-item>
  </b-list-group>
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
    }
  },
  created(){
    this.seeMyProfile([this.$route.params.id, 30, 0]);
  }
}
</script>

<style>

</style>