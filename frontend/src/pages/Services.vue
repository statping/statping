<template>
    <div v-show="core" class="container col-md-7 col-sm-12 mt-2 sm-container">

    </div>
</template>

<script>
  import Api from "../components/API"

  export default {
  name: 'Services',
  components: {

  },
  data () {
    return {
      services: null,
      groups: null,
      core: null,
      auth: null
    }
  },
  beforeMount() {
    this.loadAll()
  },
  methods: {
    async loadAll () {
      this.auth = Api.authToken()
      this.core = await Api.core()
      this.groups = await Api.groups()
      this.services = await Api.services()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
