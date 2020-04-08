<template>
    <form class="row" @submit.prevent="createIncidentUpdate">

        <div class="col-3">
            <select v-model="incident_update.type" class="form-control">
                <option value="Investigating">Investigating</option>
                <option value="Update">Update</option>
                <option value="Unknown">Unknown</option>
                <option value="Resolved">Resolved</option>
            </select>
        </div>
        <div class="col-7">
            <input v-model="incident_update.message" rows="5" name="description" class="form-control" id="message" required>
        </div>

        <div class="col-2">
            <button @click.prevent="createIncidentUpdate"
                    :disabled="!incident_update.message"
                    type="submit" class="btn btn-block btn-primary">
                Add
            </button>
        </div>
    </form>
</template>

<script>
  import Api from "../API";
  import flatPickr from 'vue-flatpickr-component';
  import 'flatpickr/dist/flatpickr.css';

  export default {
  name: 'FormIncidentUpdates',
  components: {

  },
  props: {
    incident: {
      type: Object
    }
  },
  data () {
    return {
      updates: [],
        incident_update: {
            incident: this.incident.id,
            message: "",
            type: ""
        }
    }
  },
      async mounted () {
          await this.loadUpdates()
      },
      methods: {
            async loadUpdates() {
              this.updates = await Api.incident_updates(this.incident)
            },
          async createIncidentUpdate() {
              await Api.incident_update_create(this.incident_update)
                await this.loadUpdates()
          }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
