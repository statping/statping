<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">
        {{service}}
    </div>
</template>

<script>
import Api from "../components/API"

export default {
  name: 'Service',
  components: {

  },
  data () {
    return {
      id: null,
      service: null,
    }
  },
  beforeMount() {
    this.id = this.$route.params.id
    this.getService()
  },
  methods: {
    async getService() {
        this.service = await Api.services()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
