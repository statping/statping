<template>
    <div class="mb-4" id="service_id_1">
        <div class="card">
            <div class="card-body">
                <div class="col-12">
                    <h4 class="mt-3">
                        <router-link :to="`/service/${service.id}`">{{service.name}}</router-link>
                        <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online}">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    </h4>

                    <div class="row stats_area mt-5">
                        <div class="col-4">
                            <span class="lg_number">{{service.avg_response}}ms</span>
                            Average Response
                        </div>
                        <div class="col-4">
                            <span class="lg_number">{{service.online_24_hours}}%</span>
                            Uptime last 24 Hours
                        </div>
                        <div class="col-4">
                            <span class="lg_number">{{service.online_7_days}}%</span>
                            Uptime last 7 Days
                        </div>
                    </div>

                </div>
            </div>

            <div class="chart-container">
                <ServiceChart :service="service"/>
            </div>

            <div class="row lower_canvas full-col-12 text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                <div class="col-10 text-truncate">
                    <span class="d-none d-md-inline">
                        {{smallText(service)}}
                    </span>
                </div>
                <div class="col-sm-12 col-md-2">
                    <router-link :to="serviceLink(service)" class="btn btn-sm float-right dyn-dark btn-block" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                        View Service</router-link>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
  import ServiceChart from "./ServiceChart";

  export default {
  name: 'ServiceBlock',
  components: {ServiceChart},
  props: {
    service: {
      type: Object,
      required: true
    },
  },
      methods: {
        smallText(s) {
            if (s.online) {
                return `Online, last checked ${this.ago(s.last_success)}`
            } else {
                return `Offline, last error: ${s.last_failure.issue} ${this.ago(s.last_failure.created_at)}`
            }
          },
          ago(t1) {
            const tm = this.parseTime(t1)
            return this.duration(this.$moment().utc(), tm)
          }
      }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
