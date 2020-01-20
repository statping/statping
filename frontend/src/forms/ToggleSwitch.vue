<template>
    <font-awesome-icon @click="toggleChecking" class="toggle-service text-success" icon="toggle-on" />
</template>

<script>
import Api from "../components/API";

export default {
  name: 'ToggleSwitch',
  props: {
    next: {
      type: Function,
      required: true
    },
    service: {
      type: Object,
      required: true
    }
  },
  data () {
    return {

    }
  },
  mounted() {

  },
  methods: {
    toggleChecking() {
      let c = confirm(`Are you sure you want to stop monitoring '${this.service.name}'?`)
      if (c) {

      }
      this.props.next(this.props.service)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
