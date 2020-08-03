<template>
  <div class="container my-4">
    <div class="row">
      <div class="col-md-5">
        <img class="logo" alt="Front Store Logo" src="../assets/market_research.svg">
      </div>
      <div class="col-md-7 text-center">
        <HelloWorld msg="Welcome to Your Store Administration"/>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'Home',
  components: {
    HelloWorld
  }
}
</script>

<style scoped>
  .logo{
    width: 500px;
  }
</style>
