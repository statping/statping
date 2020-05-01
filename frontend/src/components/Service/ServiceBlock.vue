<template>
    <div class="mb-md-4 mb-4">
        <div class="card index-chart" :class="{'expanded-service': expanded}">
            <div class="card-body">
                <div class="col-12">
                    <h4 class="mt-2">
                        <router-link :to="serviceLink(service)" class="d-inline-block text-truncate font-4" style="max-width: 65vw;" :in_service="service">{{service.name}}</router-link>
                        <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online}">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    </h4>

                    <ServiceTopStats :service="service"/>

                        <div v-if="expanded" class="row">
                            <Analytics title="Last Failure" :func="stats.total_failures"/>
                            <Analytics title="Total Failures" :func="stats.total_failures"/>
                            <Analytics title="Highest Latency" :func="stats.high_latency"/>
                            <Analytics title="Lowest Latency" :func="stats.lowest_latency"/>
                            <Analytics title="Total Uptime" :func="stats.high_ping"/>
                            <Analytics title="Total Downtime" :func="stats.low_ping"/>

                            <div class="col-12">
                            <router-link :to="serviceLink(service)" class="btn btn-block btn-outline-success mt-4" :class="{'btn-outline-success': service.online, 'btn-outline-danger': !service.online}">
                                View More Details
                            </router-link>
                            </div>
                         </div>

                </div>
            </div>

            <div v-show="!expanded" v-observe-visibility="visibleChart" class="chart-container">
                <ServiceChart :service="service" :visible="visible"/>
            </div>

            <div class="row lower_canvas full-col-12 text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                <div class="col-md-8 col-6">
                        <div class="dropup" :class="{show: dropDownMenu}">
                              <button style="font-size: 10pt;" @focusout="dropDownMenu = false"  @click="dropDownMenu = !dropDownMenu" type="button" class="d-none col-4 float-left btn btn-sm float-right btn-block text-white dropdown-toggle service_scale" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                  24 Hours
                              </button>
                              <div class="dropdown-menu" :class="{show: dropDownMenu}">
                                <a v-for="(timeframe, i) in timeframes" @click="timeframe.picked = true" class="dropdown-item" href="#">{{timeframe.text}}</a>
                              </div>

                            <span class="d-none float-left d-md-inline">
                                {{smallText(service)}}
                            </span>
                        </div>
                </div>

                <div class="col-md-4 col-6 float-right">
                    <button v-if="!expanded" @click="setService" class="btn btn-sm float-right dyn-dark text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                        View Service
                    </button>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
import Api from '../../API';
import Analytics from './Analytics';
import ServiceChart from "./ServiceChart";
import ServiceTopStats from "@/components/Service/ServiceTopStats";
import Graphing from '../../graphing'

export default {
    name: 'ServiceBlock',
    components: { Analytics, ServiceTopStats, ServiceChart},
    props: {
        in_service: {
            type: Object,
            required: true
        },
    },
    data() {
        return {
          timer_func: null,
            expanded: false,
            visible: false,
            dropDownMenu: false,
            timeframes: [
                {value: "72h", text: "3 Days", picked: true },
                {value: "24h", text: "Since Yesterday" },
                {value: "3", text: "3 Hours" },
                {value: "1m", text: "1 Month" },
                {value: "3", text: "Last 3 Months" },
            ],
            stats: {
                total_failures: {
                    title: "Total Failures",
                    subtitle: "Last 7 Days",
                    value: 0,
                },
                high_latency: {
                    title: "Highest Latency",
                    subtitle: "Last 7 Days",
                    value: 0,
                },
                lowest_latency: {
                    title: "Lowest Latency",
                    subtitle: "Last 7 Days",
                    value: 0,
                },
                high_ping: {
                    title: "Highest Ping",
                    subtitle: "Last 7 Days",
                    value: 0,
                },
                low_ping: {
                    title: "Lowest Ping",
                    subtitle: "Last 7 Days",
                    value: 0,
                }
            },
            track_service: null,
        }
    },
  beforeDestroy() {
    clearInterval(this.timer_func)
  },
  computed: {
      service() {
        return this.track_service
      }
  },
  async created() {
      this.track_service = this.in_service
  },
    methods: {
      async setService() {
        await this.$store.commit('setService', this.service)
        this.$router.push('/service/'+this.service.id, {props: {in_service: this.service}})
      },
        async showMoreStats() {
            this.expanded = !this.expanded;

            const failData = await Graphing.failures(this.service, 7)
            this.stats.total_failures.chart = failData.data;
            this.stats.total_failures.value = failData.total;

            const hitsData = await Graphing.hits(this.service, 7)

            this.stats.high_latency.chart = hitsData.chart;
            this.stats.high_latency.value = this.humanTime(hitsData.high);

            this.stats.lowest_latency.chart = hitsData.chart;
            this.stats.lowest_latency.value = this.humanTime(hitsData.low);

            const pingData = await Graphing.pings(this.service, 7)
            this.stats.high_ping.chart = pingData.chart;
            this.stats.high_ping.value = this.humanTime(pingData.high);

            this.stats.low_ping.chart = pingData.chart;
            this.stats.low_ping.value = this.humanTime(pingData.low);
        },
        smallText(s) {
          const incidents = s.incidents
            if (s.online) {
                return `Checked ${this.ago(s.last_success)} ago and responded in ${this.humanTime(s.latency)}`
            } else {
                const last = s.last_failure
                if (last) {
                    return `Offline, last error: ${last} ${this.ago(last.created_at)}`
                }
              if (!this.toUnix(s.last_success)) {
                return `Service has never been online`
              }
                return `Offline`
            }
        },
        visibleChart(isVisible, entry) {
                if (isVisible && !this.visible) {
                    this.visible = true

                  if (!this.timer_func) {
                    this.timer_func = setInterval(async () => {
                      this.track_service = await Api.service(this.service.id)
                    }, this.track_service.check_interval * 1000)
                  }
                }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
