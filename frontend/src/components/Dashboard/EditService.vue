<template>
    <FormService v-if="ready" :in_service="service"/>
</template>

<script>
  import FormGroup from "../../forms/Group";
  import Api from "../../components/API";
  import ToggleSwitch from "../../forms/ToggleSwitch";
  import draggable from 'vuedraggable'
  import FormService from "../../forms/Service";

  export default {
  name: 'EditService',
  components: {
    FormService,
    ToggleSwitch,
    FormGroup,
    draggable
  },
  props: {

  },
  data () {
    return {
      ready: false,
      service: null
    }
  },
  computed: {

  },
  async beforeCreate() {
    if (this.$route.params.id) {
      this.service = await Api.service(this.$route.params.id)
    }
    this.ready = true
  },
  beforeMount() {

  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
