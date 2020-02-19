<template v-if="service">
    <div class="col-12 card mb-3" style="min-height: 260px" :class="{'offline-card': !service.online}">
        <div class="card-body">
            <h5 class="card-title"><router-link :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge float-right" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h5>
            <div v-if="loaded && service.online" class="row">
                <div class="col-md-6 col-sm-12">
                    <ServiceSparkLine :title="set1_name" subtitle="Last Day Latency" :series="set1"/>
                </div>
                <div class="col-md-6 col-sm-12">
                    <ServiceSparkLine :title="set2_name" subtitle="Last 7 Days Latency" :series="set2"/>
                </div>
            </div>
        </div>
        <span v-for="(failure, index) in failures" v-bind:key="index" class="alert alert-light">
            Failed {{duration(current(), failure.created_at)}}<br>
            {{failure.issue}}
        </span>
    </div>
</template>

<script>
  import ServiceSparkLine from "./ServiceSparkLine";
  import Api from "../API";

  export default {
      name: 'ServiceInfo',
      components: {
          ServiceSparkLine
      },
      props: {
          service: {
              type: Object,
              required: true
          }
      },
      data () {
          return {
              set1: [],
              set2: [],
              loaded: false,
              set1_name: "",
              set2_name: "",
              failures: null
          }
      },
      async mounted () {
          this.set1 = await this.getHits(24 * 2, "hour")
          this.set1_name = this.calc(this.set1)
          this.set2 = await this.getHits(24 * 7, "hour")
          this.set2_name = this.calc(this.set2)
          this.loaded = true
      },
      methods: {
          async getHits (hours, group) {
              const start = this.ago(3600 * hours)
              if (!this.service.online) {
                  this.failures = await Api.service_failures(this.service.id, this.now()-360, this.now(), 5)
                  return [ { name: "None", data: [] } ]
              }
              const data = await Api.service_hits(this.service.id, start, this.now(), group)
              if (!data) {
                  return [ { name: "None", data: [] } ]
              }
              return [ { name: "Latency", data: data.data } ]
          },
          calc (s) {
              let data = s[0].data
              if (data.length > 1) {
                  let total = 0
                  data.forEach((f) => {
                      total += f.y
                  });
                  total = total / data.length
                  return total.toFixed(0) + "ms Average"
              } else {
                  return "Offline"
              }
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .offline-card {
        background-color: #fff5f5;
    }
    .fade-enter-active, .fade-leave-active {
        transition: opacity .75s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
        opacity: 0;
    }
</style>
