<template>
<div>

</div>
</template>

<script>
  import Api from "../API";
  import flatPickr from 'vue-flatpickr-component';
  import 'flatpickr/dist/flatpickr.css';
  import FormIncidentUpdates from './IncidentUpdates';

  export default {
  name: 'FormIncident',
  components: {
      FormIncidentUpdates
  },
  props: {
    service: {
      type: Object
    }
  },
  data () {
    return {
      incident: {
        title: "",
        description: "",
        service: this.service.id,
      },
        incidents: [],
    }
  },
      async mounted () {
          await this.loadIncidents()
      },
      methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
