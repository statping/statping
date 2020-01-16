<template>
    <div class="col-12 mt-3">

        <div class="row stats_area mb-5">
            <div class="col-4">
                <span class="lg_number">{{services.length}}</span>
                Total Services
            </div>
            <div class="col-4">
                <span class="lg_number">48</span>
                Failures last 24 Hours
            </div>
            <div class="col-4">
                <span class="lg_number">7</span>
                Online Services
            </div>
        </div>

        <div v-for="(service, index) in services" v-bind:key="index">
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
  props: {
    services: Array
  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
