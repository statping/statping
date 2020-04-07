<template v-if="service">
    <div class="col-12 card mb-4" style="min-height: 280px;" :class="{'offline-card': !service.online}">
        <div class="card-body p-3 p-md-1 pt-md-3 pb-md-1">
            <h4 class="card-title mb-4">
                <router-link :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge float-right" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h4>

            <transition name="fade">
            <div v-if="loaded && service.online" class="row pb-3">
                <div class="col-md-6 col-sm-12 mt-2 mt-md-0 mb-3">
                    <ServiceSparkLine :title="set2_name" subtitle="Latency Last 24 Hours" :series="set2"/>
                </div>
                <div class="col-md-6 col-sm-12 mt-4 mt-md-0 mb-3">
                    <ServiceSparkLine :title="set1_name" subtitle="Latency Last 7 Days" :series="set1"/>
                </div>

                <div class="d-none row col-12 mt-4 pt-1 mb-3 align-content-center">

                    <StatsGen :service="service"
                              title="Since Yesterday"
                              :start="this.toUnix(this.nowSubtract(86400 * 2))"
                              :end="this.toUnix(this.nowSubtract(86400))"
                              group="24h" expression="latencyPercent"/>

                    <StatsGen :service="service"
                              title="7 Day Change"
                              :start="this.toUnix(this.nowSubtract(86400 * 7))"
                              :end="this.toUnix(this.now())"
                              group="24h" expression="latencyPercent"/>

                    <StatsGen :service="service"
                              title="Max Latency"
                              :start="this.toUnix(this.nowSubtract(86400 * 2))"
                              :end="this.toUnix(this.nowSubtract(86400))"
                              group="24h" expression="latencyPercent"/>

                    <StatsGen :service="service"
                              title="Uptime"
                              :start="this.toUnix(this.nowSubtract(86400 * 2))"
                              :end="this.toUnix(this.nowSubtract(86400))"
                              group="24h" expression="latencyPercent"/>
                </div>

                    <div class="col-4">
                        <button @click.prevent="Tab('incident')" class="btn btn-block btn-outline-secondary incident" :class="{'text-white btn-secondary': openTab==='incident'}" >Incidents</button>
                    </div>
                    <div class="col-4">
                        <button @click.prevent="Tab('message')" class="btn btn-block btn-outline-secondary message" :class="{'text-white btn-secondary': openTab==='message'}">Announcements</button>
                    </div>
                    <div class="col-4">
                        <button @click.prevent="Tab('failures')" class="btn btn-block btn-outline-secondary failures" :disabled="service.stats.failures === 0" :class="{'text-white btn-secondary': openTab==='failures'}">
                            Failures <span class="badge badge-danger float-right mt-1">{{service.stats.failures}}</span></button>
                    </div>

                <div v-if="openTab === 'incident'" class="col-12 mt-4">
                    <FormIncident :service="service" />
                </div>

                <div v-if="openTab === 'message'" class="col-12 mt-4">
                    <FormMessage :service="service"/>
                </div>

                <div v-if="openTab === 'failures'" class="col-12 mt-4">
                    <ServiceFailures :service="service"/>
                </div>

            </div>
            </transition>

        </div>

        <span v-for="(failure, index) in failures" v-bind:key="index" class="alert alert-light">
            Failed {{failure.created_at}}<br>
            {{failure.issue}}
        </span>

    </div>
</template>

<script>
  import FormIncident from '../../forms/Incident';
  import FormMessage from '../../forms/Message';
  import ServiceFailures from './ServiceFailures';
  import ServiceSparkLine from "./ServiceSparkLine";
  import Api from "../../API";
  import StatsGen from "./StatsGen";

  export default {
      name: 'ServiceInfo',
      components: {
          ServiceFailures,
          FormIncident,
          FormMessage,
          StatsGen,
          ServiceSparkLine
      },
      props: {
          service: {
              type: Object,
              required: true
          }
      },
      data() {
          return {
              openTab: "",
              set1: [],
              set2: [],
              loaded: false,
              set1_name: "",
              set2_name: "",
              failures: null
          }
      },
      async mounted() {
          this.set1 = await this.getHits(24 * 7, "6h")
          this.set1_name = this.calc(this.set1)
          this.set2 = await this.getHits(24, "1h")
          this.set2_name = this.calc(this.set2)
          this.loaded = true
      },
      methods: {
          Tab(name) {
              if (this.openTab === name) {
                  this.openTab = ''
                  return
              }
              this.openTab=name;
          },
        sinceYesterday(data) {
            window.console.log(data)
          let total = 0
          data.forEach((f) => {
            total += parseInt(f.y)
          });
          total = total / data.length
        },
          async getHits(hours, group) {
              const start = this.nowSubtract(3600 * hours)
              const fetched = await Api.service_hits(this.service.id, this.toUnix(start), this.toUnix(this.now()), group, false)

              const data = this.convertToChartData(fetched, 0.001, true)

              return [{name: "Latency", ...data}]

          },
          calc(s) {
              let data = s[0].data

              if (data) {
                  let total = 0
                  data.forEach((f) => {
                      total += parseInt(f.y)
                  });
                  total = total / data.length
                  return total.toFixed(0) + "ms"
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
