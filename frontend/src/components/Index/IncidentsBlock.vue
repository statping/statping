<template>
    <div class="row mt-4">
        <div v-for="(incident, i) in incidents" class="col-12">
            <h6><span class="badge badge-secondary">New</span>
                {{incident.title}}
            <span class="font-2">{{incident.created_at}}</span>
            </h6>
            {{incident.description}}
            {{incident}}
        </div>
    </div>
</template>

<script>
import Api from '../../API';

export default {
  name: 'IncidentsBlock',
    props: {
        service: {
            type: Object
        }
    },
    data() {
        return {
            incidents: null,
        }
    },
    mounted () {
        this.getIncidents()
    },
    methods: {
        async getIncidents() {
            this.incidents = await Api.incidents_service(this.service)
        },
        async getIncidentsUpdates() {
            this.incidents = await Api.incidents_service(this.service)
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
