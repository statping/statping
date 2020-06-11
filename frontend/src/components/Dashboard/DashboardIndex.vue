<template>
    <div class="col-12 mt-4 mt-md-3">

        <div class="row stats_area mb-5">
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{$store.getters.services.length}}</span>
                <span class="font-2">Total Services</span>
            </div>
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{failuresLast24Hours()}}</span>
                <span class="font-2">Failures last 24 Hours</span>
            </div>
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{$store.getters.onlineServices(true).length}}</span>
                <span class="font-2">Online Services</span>
            </div>
        </div>

        <div v-for="(service, index) in services" class="service_block" v-bind:key="index">
            <ServiceInfo :service=service />
        </div>
    </div>
</template>

<script>
  const ServiceInfo = () => import('@/components/Service/ServiceInfo')

  export default {
      name: 'DashboardIndex',
      components: {
          ServiceInfo
      },
    data() {
        return {
          visible: false
        }
    },
      computed: {
          services() {
              return this.$store.getters.services
          }
      },
      methods: {

          failuresLast24Hours() {
              let total = 0;
              this.services.map((s) => {
                  total += s.failures_24_hours
              })
              return total
          },

      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
