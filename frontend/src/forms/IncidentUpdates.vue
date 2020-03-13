<template>
<div>

    <div v-for="(update, i) in updates" class="col-12 bg-active card pt-2 pb-2 mt-3 pl-3 pr-3">
        <span class="font-4">
            <font-awesome-icon v-if="update.type === 'Resolved'" icon="check-circle" class="mr-2"/>
            <font-awesome-icon v-if="update.type === 'Update'" icon="asterisk" class="mr-2"/>
            <font-awesome-icon v-if="update.type === 'Investigating'" icon="lightbulb" class="mr-2"/>
            <font-awesome-icon v-if="update.type === 'Unknown'" icon="question" class="mr-2"/>

            {{update.type}}
        </span>
        <span class="font-3 mt-3">{{update.message}}</span>
    </div>

        <div class="col-12 bg-active card pt-2 pb-2 mt-3 pl-3 pr-3">

            <form @submit.prevent="createIncidentUpdate">

            <div class="form-group row">
                <label class="col-sm-4 col-form-label">Update Type</label>
                <div class="col-sm-8">
                    <select v-model="incident_update.type" class="form-control">
                        <option value="Investigating">Investigating</option>
                        <option value="Update">Update</option>
                        <option value="Unknown">Unknown</option>
                        <option value="Resolved">Resolved</option>
                    </select>
                </div>
            </div>

            <div class="form-group row">
                <label class="col-sm-4 col-form-label">New Update</label>
                <div class="col-sm-8">
                    <textarea v-model="incident_update.message" rows="5" name="description" class="form-control" id="description" required></textarea>
                </div>
            </div>

            <div class="form-group row">
                <div class="col-sm-12">
                    <button @click.prevent="createIncidentUpdate"
                            :disabled="!incident.title || !incident.description"
                            type="submit" class="btn btn-block btn-primary">
                        Add Update
                    </button>
                </div>
            </div>

            </form>

        </div>


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
      type: Object
    }
  },
  data () {
    return {
      updates: this.incident.updates,
        incident_update: {
            incident: this.incident,
            message: "",
            type: ""
        }
    }
  },
      async mounted () {
          this.updates = await Api.incident_updates(this.incident)
      },
      methods: {
          async createIncidentUpdate(incident) {
              await Api.incident_update_create(incident, this.incident_update)
              const updates = await Api.incident_updates()
          }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
