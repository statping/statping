<template>
    <div class="card-body pt-3">

        <div v-if="updates.length===0" class="alert alert-link text-danger">
            No updates found, create a new Incident Update below.
        </div>

        <div v-for="update in updates" :key="update.id">
            <IncidentUpdate :update="update" :onUpdate="loadUpdates" :admin="true"/>
        </div>

        <form class="row" @submit.prevent="createIncidentUpdate">
            <div class="col-12 col-md-3 mb-3 mb-md-0">
                <select v-model="incident_update.type" class="form-control">
                    <option value="Investigating">Investigating</option>
                    <option value="Update">Update</option>
                    <option value="Unknown">Unknown</option>
                    <option value="Resolved">Resolved</option>
                </select>
            </div>
            <div class="col-12 col-md-7 mb-3 mb-md-0">
                <input v-model="incident_update.message" name="description" class="form-control" id="message" required>
            </div>

            <div class="col-12 col-md-2">
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
    const IncidentUpdate = () => import(/* webpackChunkName: "index" */ "@/components/Elements/IncidentUpdate");

    export default {
        name: 'FormIncidentUpdates',
        components: {IncidentUpdate},
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
