<template>
    <div class="card-body bg-light pt-3">

        <div v-if="updates.length===0" class="alert alert-link text-danger">
            No updates found, create a new Incident Update below.
        </div>

        <div v-for="update in updates" :key="update.id">
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
        components: {},
        props: {
            incident: {
                type: Object,
                required: true
            }
        },
        data () {
            return {
                updates: [],
                incident_update: {
                    incident: this.incident.id,
                    message: "",
                    type: "Investigating" // TODO: default to something.. theres is no error checking for blank submission...
                }
            }
        },

        async mounted() {
            await this.loadUpdates()
        },

        methods: {

            async delete_update(update) {
                this.res = await Api.incident_update_delete(update)
                if (this.res.status === "success") {
                    this.updates = this.updates.filter(obj => obj.id !== update.id); // this is better in terms of not having to querry the db to get a fresh copy of all updates
                    //await this.loadUpdates()
                }
            },

            async createIncidentUpdate() {
                this.res = await Api.incident_update_create(this.incident_update)
                if (this.res.status === "success") {
                    this.updates.push(this.res.output) // this is better in terms of not having to querry the db to get a fresh copy of all updates
                    //await this.loadUpdates()
                } // TODO: further error checking here... maybe alert user it failed with modal or so

                // reset the form data
                this.incident_update = {
                    incident: this.incident.id,
                    message: "",
                    type: "Investigating"
                }

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
