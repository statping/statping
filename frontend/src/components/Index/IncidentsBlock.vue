<template>
    <div class="row">
        <div v-for="(incident, i) in incidents" class="col-12 mt-2">
            <span class="braker mt-1 mb-3"></span>
            <h6>{{incident.title}}
                <span class="font-2 float-right">{{niceDate(incident.created_at)}}</span>
            </h6>
            <div class="font-2 mb-3" v-html="incident.description"></div>
                <IncidentUpdate v-for="(update, i) in incident.updates" v-bind:key="i" :update="update" :admin="false"/>
        </div>
    </div>
</template>

<script>
import Api from '../../API';
import IncidentUpdate from "@/components/Elements/IncidentUpdate";

export default {
  name: 'IncidentsBlock',
  components: {
    IncidentUpdate
  },
  props: {
        service: {
            type: Object,
            required: true
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
        badgeClass(val) {
          switch (val.toLowerCase()) {
            case "resolved":
              return "badge-success"
            case "update":
              return "badge-info"
            case "investigating":
              return "badge-danger"
          }
        },
      async getIncidents() {
        this.incidents = await Api.incidents_service(this.service.id)
      },
      async incident_updates(incident) {
        return await Api.incident_updates(incident)
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
