<template>
  <div id="app" v-if="loaded">
    <router-view/>
      <Footer version="DEV" />
  </div>
</template>

<script>
  import Api from './components/API';
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
    async created () {
     await this.setAllObjects()
      this.loaded = true
      this.$store.commit('setHasPublicData', true)
   },
    methods: {
      async setAllObjects () {
          await this.setCore()
          await this.setServices()
          await this.setGroups()
          await this.setMessages()
          await this.setToken()
          this.$store.commit('setHasPublicData', true)
        this.loaded = true
      },
          async setCore () {
              const core = await Api.core()
              this.$store.commit('setCore', core)
          },
          async setToken () {
              const token = await Api.token()
              this.$store.commit('setToken', token)
          },
          async setServices () {
              const services = await Api.services()
              this.$store.commit('setServices', services)
          },
          async setGroups () {
              const groups = await Api.groups()
              this.$store.commit('setGroups', groups)
          },
          async setMessages () {
              const messages = await Api.messages()
              this.$store.commit('setMessages', messages)
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
