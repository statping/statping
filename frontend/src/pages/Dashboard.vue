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
      async mounted() {
          if (this.$route.path !== "/login") {
              try {
                  const u = await Api.users()
              } catch (e) {
                  this.$router.push('/logout')
              }
          }
      },
  data () {
    return {
      authenticated: false
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
