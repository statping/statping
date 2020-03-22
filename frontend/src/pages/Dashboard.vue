<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <TopNav :admin="$store.state.admin"/>
        <router-view :admin="$store.state.admin"/>
    </div>
</template>

<script>
  import Api from "../API"
  import TopNav from "../components/Dashboard/TopNav";

  export default {
  name: 'Dashboard',
  components: {
    TopNav,
  },
  data () {
      return {
          authenticated: false,
      }
  },
  async mounted() {
        const core = await Api.core()
      this.$store.commit('setAdmin', core.admin)
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
