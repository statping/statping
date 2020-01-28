<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <TopNav/>
        <router-view></router-view>
    </div>
</template>

<script>
  import TopNav from "../components/Dashboard/TopNav";

  export default {
  name: 'Dashboard',
  components: {
    TopNav,
  },
  data () {
    return {
      authenticated: false
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
