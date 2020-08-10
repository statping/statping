<template>
    <div class="container col-md-7 col-sm-12 mt-md-5">
        <TopNav :admin="admin"/>
        <router-view :admin="admin"/>
    </div>
</template>

<script>
  const TopNav = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/TopNav')

  export default {
  name: 'Dashboard',
  components: {
    TopNav,
  },
  data () {
      return {
          authenticated: false,
          loaded: false,
      }
  },
    computed: {
      admin() {
        return this.$store.getters.admin
      },
      user() {
        return this.$store.getters.user
      }
    },
    mounted() {
      // if (!this.user || !this.admin) {
      //   this.$router.push('/login')
      // }
    }
  }
</script>

