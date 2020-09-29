<template>
    <div class="col-12">
        <div v-if="!ready" class="row mt-5">
            <div class="col-12 text-center">
                <font-awesome-icon icon="circle-notch" size="3x" spin/>
            </div>
            <div class="col-12 text-center mt-3 mb-3">
                <span class="text-muted">Loading Service</span>
            </div>
        </div>
        <FormService v-if="ready" :in_service="service"/>
    </div>
</template>

<script>
  const FormGroup = () => import(/* webpackChunkName: "dashboard" */ "../../forms/Group");
  import Api from "../../API";
  const ToggleSwitch = () => import(/* webpackChunkName: "dashboard" */ "../../forms/ToggleSwitch");
  const draggable = () => import(/* webpackChunkName: "dashboard" */ 'vuedraggable')
  const FormService = () => import(/* webpackChunkName: "dashboard" */ "../../forms/Service");

  export default {
    name: 'EditService',
    components: {
      FormService,
      ToggleSwitch,
      FormGroup,
      draggable
    },
    created() {
        this.fetchData()
    },
    watch: {
      '$route': 'fetchData'
    },
    data () {
      return {
        service: null,
        ready: false
      }
    },
    methods: {
      async fetchData () {
        if (!this.$route.params.id) {
          this.ready = true
          return
        }
        this.service = await Api.service(this.$route.params.id)
        this.ready = true
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
