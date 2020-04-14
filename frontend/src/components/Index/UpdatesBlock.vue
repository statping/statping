<template>
    <div class="row">
        <div v-for="(update, i) in updates" v-bind:key="i" class="col-12 mt-3">
            <div class="col-md-2 col-12">
                <span class="badge text-uppercase" :class="badgeClass(update.type)">{{update.type}}</span>
            </div>
            <div class="col-md-12 col-12 mt-2 font-3">{{update.message}}</div>
            <div class="col-12 font-1 float-right text-black-50 mt-2">{{ago(update.created_at)}} ago</div>
        </div>
    </div>
</template>

<script>
import Api from '../../API';

export default {
  name: 'UpdatesBlock',
    props: {
        incident: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            updates: null,
        }
    },
    mounted () {
        this.getIncidentUpdates()
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
      async getIncidentUpdates() {
        this.updates = await Api.incident_updates(this.incident)
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
