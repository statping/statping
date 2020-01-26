<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <TopNav/>
        <router-view></router-view>
    </div>
</template>

<script>
  import Login from "./Login";
  import TopNav from "../components/Dashboard/TopNav";
  import Api from "../components/API";

  export default {
  name: 'Dashboard',
  components: {
    TopNav,
    Login,
  },
  data () {
    return {
      authenticated: false
    }
  },
    mounted() {
    if (this.$store.getters.token !== null) {
      this.authenticated = true
    }
  },
  methods: {
    async setServices () {
      const services = await Api.services()
      this.$store.commit('setServices', services)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
