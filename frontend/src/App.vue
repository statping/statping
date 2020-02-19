<template>
  <div id="app">
    <router-view :loaded="loaded"/>
      <Footer :logged_in="logged_in" :version="version" v-if="$route.path !== '/setup'"/>
  </div>
</template>

<script>
  import Api from './components/API';
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
    }
  },
      async created() {
        await this.$store.dispatch('loadRequired')

        if (this.$store.getters.core.logged_in) {
          await this.$store.dispatch('loadAdmin')
        }
          this.loaded = true
          if (!this.$store.getters.core.setup) {
            this.$router.push('/setup')
          }
          window.console.log('finished loadRequired')
      },
    async mounted() {
          if (this.$route.path !== '/setup') {
              const tk = localStorage.getItem("statping_user")
              if (this.$store.getters.core.logged_in) {
                this.logged_in = true
                  await this.$store.dispatch('loadAdmin')
              }
          }
    },
    methods: {

    }
}
</script>

<style lang="scss">
    @import "./assets/css/bootstrap.min.css";
    @import "./assets/scss/variables";
    @import "./assets/scss/base";
    @import "./assets/scss/mobile";
</style>
