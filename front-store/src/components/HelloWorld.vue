<template>
  <div class="hello p-4">
    <h1>{{ msg }}</h1>
    <p>
      Here you can sync up your data bussiness data and view it.
    </p>
    <h3>About</h3>
    <p>
      Hi! My name is <a href="http://github.com/juanhenaoparra">Juan Sebastian Henao</a>, I come from 2000's and I'm from Manizales.
    </p>
    <p>The infrastructure used here: </p>
    <ul>
      <li><strong>GO and Chi</strong> for backend</li>
      <li><strong>DGraph</strong> as database</li>
      <li><strong>Vue.js</strong> for the frontend</li>
    </ul>
    <h2>Enjoy it!</h2>
  </div>
</template>

<script>
import {mapActions} from 'vuex';

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  methods: {
    ...mapActions(['getAllBuyers', 'getBuyerProfile']),
  },
  mounted(){
    // this.getAllBuyers([10, 0]);
    // this.getBuyerProfile(["1fe89491", 10, 0]);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
