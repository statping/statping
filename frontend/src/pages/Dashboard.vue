<template>
    <div>
        <Login v-if="!authenticated"/>

        <div v-if="authenticated" class="container col-md-7 col-sm-12 mt-md-5 bg-light">
            <TopNav/>
            <router-view></router-view>
         </div>
    </div>
</template>

<script>
  import Login from "./Login";
  import TopNav from "../components/Dashboard/TopNav";

  export default {
  name: 'Dashboard',
  components: {
    TopNav,
    Login,
  },
  data () {
    return {
      view: "DashboardIndex",
      authenticated: false
    }
  },
  mounted() {
    if (this.$store.getters.token !== null) {
      this.authenticated = true
    }
  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
