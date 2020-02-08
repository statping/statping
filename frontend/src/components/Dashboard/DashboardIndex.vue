<template>
    <div class="col-12 mt-3">

        <div class="row stats_area mb-5">
            <div class="col-4">
                <span class="lg_number">{{$store.getters.services.length}}</span>
                Total Services
            </div>
            <div class="col-4">
                <span class="lg_number">{{failuresLast24Hours()}}</span>
                Failures last 24 Hours
            </div>
            <div class="col-4">
                <span class="lg_number">{{$store.getters.onlineServices(true).length}}</span>
                Online Services
            </div>
        </div>

        <div v-for="(service, index) in $store.getters.services" v-bind:key="index">
            <ServiceInfo :service=service />
        </div>
    </div>
</template>

<script>
  import ServiceInfo from "./ServiceInfo";

  export default {
  name: 'DashboardIndex',
  components: {
    ServiceInfo
  },
    data () {
      return {

      }
    },
      computed: {

      },
      async created() {

      },
    methods: {
        failuresLast24Hours() {
            let total = 0;
            this.$store.getters.services.map((s) => { total += s.failures_24_hours })
            return total
        },

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
