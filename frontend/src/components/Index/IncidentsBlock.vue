<template>
    <div class="row">
        <div v-for="(incident, i) in incidents" class="col-12 mt-4 mb-3">
            <span class="braker mt-1 mb-3"></span>
            <h6>Incident: {{incident.title}}
                <span class="font-2 float-right">{{niceDate(incident.created_at)}}</span>
            </h6>
            <span class="font-2" v-html="incident.description"></span>

            <UpdatesBlock :incident="incident"/>

        </div>
    </div>
</template>

<script>
import Api from '../../API';
import UpdatesBlock from "@/components/Index/UpdatesBlock";

export default {
  name: 'IncidentsBlock',
  components: {UpdatesBlock},
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
        await Api.incident_updates(incident).then((d) => {return d})
        return o
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
