<template>
  <div id="app">
    <router-view/>
      <Footer version="DEV" v-if="$route.path !== '/setup'"/>
  </div>
</template>

<script>
  import Footer from "./components/Footer";

  export default {
  name: 'app',
  components: {
    Footer
  },
  data () {
    return {
      loaded: false
    }
  },
    async mounted() {
      if (this.$route.path !== '/setup') {
        // const tk = JSON.parse(localStorage.getItem("statping_user"))
        if (!this.$store.getters.hasPublicData) {
          await this.$store.dispatch('loadRequired')
        }
      }
      this.loaded = true
    },
    methods: {
      async setAllObjects () {
        this.loaded = true
      }
    }
}
</script>

<style lang="scss">
    @import "./assets/css/bootstrap.min.css";
    @import "./assets/scss/variables";
    @import "./assets/scss/base";
    @import "./assets/scss/mobile";
</style>
