<template>
    <div class="col-12">
    <div class="card">
    <div class="card-body">
        <FormService :in_service="service"/>
    </div>
    </div>
    </div>
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
      service: {}
    }
  },
  computed: {

  },
  async beforeCreate() {
    const id = this.$route.params.id
    if (id) {
      this.service = await Api.service(id)
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
