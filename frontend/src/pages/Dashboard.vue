<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <TopNav/>
        <router-view/>
    </div>
</template>

<script>
    import Api from "../components/API"
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
  },
      async mounted() {
          const core = await Api.core()
          if (!core.logged_in) {
              this.$router.push('/login')
          }
      },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
