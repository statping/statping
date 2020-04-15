<template>
    <div class="card-body bg-light pt-3">

        <div v-if="updates.length===0" class="alert alert-link text-danger">
            No updates found, create a new Incident Update below.
        </div>

        <div v-for="(update, i) in updates">
            <div class="alert alert-light" role="alert">
                <span class="badge badge-pill badge-info text-uppercase">{{update.type}}</span>
                <span class="float-right font-2">{{ago(update.created_at)}} ago</span>
                <span class="d-block mt-2">{{update.message}}
                        <button @click="delete_update(update)" type="button" class="close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        </span>
            </div>
        </div>
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
    </div>
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
      type: Object,
      required: true
    }
  },
  data () {
    return {
      updates: this.incident.updates,
        incident_update: {
            incident: this.incident.id,
            message: "",
            type: ""
        }
    }
  },
    beforeRouteUpdate (to, from, next) {

    },
    async mounted() {
      await this.loadUpdates()
    },
    methods: {
      async delete_update(update) {
        await Api.incident_update_delete(update)
        await this.loadUpdates()
      },
          async createIncidentUpdate() {
              await Api.incident_update_create(this.incident_update)
                await this.loadUpdates()
          },
          async loadUpdates() {
                this.updates = await Api.incident_updates(this.incident)
          }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
