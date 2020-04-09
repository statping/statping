<template>
  <div id="app">
    <router-view :app="app" :loaded="loaded"/>
      <Footer :logged_in="logged_in" :version="version" v-if="$route.path !== '/setup'"/>
  </div>
</template>

<script>
  import Api from './API';
  import Footer from "./components/Index/Footer";

  export default {
  name: 'app',
  components: {
    Footer
  },
  data () {
    return {
      loaded: false,
        version: "",
        logged_in: false,
        app: null
    }
  },
      computed: {
          core() {
            return this.$store.getters.core
          }
      },
  async created() {
      this.app = await this.$store.dispatch('loadRequired')

      this.app = {...this.$store.state}

    if (this.core.logged_in) {
      await this.$store.dispatch('loadAdmin')
    }
      this.loaded = true
      if (!this.core.setup) {
        this.$router.push('/setup')
      }
  },
    async mounted() {
          if (this.$route.path !== '/setup') {
              const tk = localStorage.getItem("statping_user")
              if (this.core.logged_in) {
                this.logged_in = true
                  await this.$store.dispatch('loadAdmin')
              }
          }
    }
}
</script>

<style lang="scss">
    @import "./assets/css/bootstrap.min.css";
    @import "./assets/scss/main";
</style>
