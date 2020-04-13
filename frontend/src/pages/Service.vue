<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <div class="col-12 mb-4">

            <span class="mt-3 mb-3 text-white d-md-none btn d-block d-md-none" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                {{service.online ? "ONLINE" : "OFFLINE"}}
            </span>

            <h4 class="mt-2">
                <router-link to="/" class="text-black-50 text-decoration-none">{{core.name}}</router-link> - <span class="text-muted">{{service.name}}</span>
                <span class="badge float-right d-none d-md-block" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h4>

            <ServiceTopStats :service="service"/>

            <div v-for="(message, index) in $store.getters.serviceMessages(service.id)" v-if="messageInRange(message)">
                <MessageBlock :message="message"/>
            </div>

            <div class="row mt-5 mb-4">
                <span class="col-6 font-2">
                    <flatPickr v-model="start_time" type="text" name="start_time" class="form-control form-control-plaintext" id="start_time" value="0001-01-01T00:00:00Z" required />
                </span>
                <span class="col-6 font-2">
                    <flatPickr v-model="end_time" type="text" name="end_time" class="form-control form-control-plaintext" id="end_time" value="0001-01-01T00:00:00Z" required />
                </span>
            </div>

            <div v-if="series" class="service-chart-container">
                <apexchart width="100%" height="420" type="area" :options="chartOptions" :series="series"></apexchart>
            </div>

            <div class="service-chart-heatmap mt-5 mb-4">
                <ServiceHeatmap :service="service"/>
            </div>

            <div v-if="load_timedata" class="col-12">

                <apexchart width="100%" height="420" type="rangeBar" :options="timeRangeOptions" :series="rangeSeries"></apexchart>
            </div>

            <nav v-if="service.failures" class="nav nav-pills flex-column flex-sm-row mt-3" id="service_tabs">
                <a @click="tab='failures'" class="flex-sm-fill text-sm-center nav-link active">Failures</a>
                <a @click="tab='incidents'" class="flex-sm-fill text-sm-center nav-link">Incidents</a>
                <a @click="tab='checkins'" v-if="$store.getters.token" class="flex-sm-fill text-sm-center nav-link">Checkins</a>
                <a @click="tab='response'" v-if="$store.getters.token" class="flex-sm-fill text-sm-center nav-link">Response</a>
            </nav>


            <div v-if="service.failures" class="tab-content">
                <div class="tab-pane fade active show">
                    <ServiceFailures :service="service"/>
                </div>

                <div class="tab-pane fade" :class="{active: tab === 'incidents'}" id="incidents">

                </div>

                <div class="tab-pane fade" :class="{show: tab === 'checkins'}" id="checkins">

                    <div class="card">
                        <div class="card-body">
                            <Checkin :service="service"/>
                        </div>
                    </div>

                </div>

                <div class="tab-pane fade" :class="{show: tab === 'response'}" id="response">
                    <div class="col-12 mt-4">
                        <h3>Last Response</h3>
                        <textarea rows="8" class="form-control" readonly>invalid route</textarea>
                        <div class="form-group row mt-2">
                            <label for="last_status_code" class="col-sm-3 col-form-label">HTTP Status Code</label>
                            <div class="col-sm-2">
                                <input type="text" id="last_status_code" class="form-control" value="200" readonly>
                            </div>
                        </div>
                    </div>
                </div>

            </div>

        </div>
    </div>

</template>

<script>
  import Api from "../API"
  import MessageBlock from '../components/Index/MessageBlock';
  import ServiceFailures from '../components/Service/ServiceFailures';
  import Checkin from "../forms/Checkin";
  import ServiceHeatmap from "@/components/Service/ServiceHeatmap";
  import ServiceTopStats from "@/components/Service/ServiceTopStats";
  import store from '../store'
  import flatPickr from 'vue-flatpickr-component';
  import 'flatpickr/dist/flatpickr.css';
  const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };

  const axisOptions = {
    labels: {
      show: true
    },
    crosshairs: {
      show: false
    },
    lines: {
      show: true
    },
    tooltip: {
      enabled: false
    },
    axisTicks: {
      show: true
    },
    grid: {
      show: true
    },
    marker: {
      show: false
    }
  };

export default {
    name: 'Service',
    components: {
        ServiceTopStats,
        ServiceHeatmap,
        ServiceFailures,
        MessageBlock,
        Checkin,
        flatPickr
    },
    data() {
        return {
            id: 0,
            tab: "failures",
            authenticated: false,
            ready: true,
            data: null,
            messages: [],
            failures: [],
            start_time: this.nowSubtract(84600 * 30),
            end_time: new Date(),
            timedata: [],
            load_timedata: false,
            dailyRangeOpts: {
                chart: {
                    height: 500,
                    width: "100%",
                    type: "area",
                }
            },
          timeRangeOptions: {
            chart: {
              height: 200,
              type: 'rangeBar'
            },
            plotOptions: {
              bar: {
                horizontal: true,
                distributed: true,
                dataLabels: {
                  hideOverflowingLabels: false
                }
              }
            },
            dataLabels: {
              enabled: true,
              formatter: (val, opts) => {
                var label = opts.w.globals.labels[opts.dataPointIndex]
                var a = this.parseISO(val[0])
                var b = this.parseISO(val[1])
                return label
              },
              style: {
                colors: ['#f3f4f5', '#fff']
              }
            },
            xaxis: {
              type: 'datetime'
            },
            yaxis: {
              show: false
            },
            grid: {
              row: {
                colors: ['#f3f4f5', '#fff'],
                opacity: 1
              }
            }
          },
            chartOptions: {
                chart: {
                    events: {
                        beforeZoom: async (chartContext, { xaxis }) => {
                            const start = (xaxis.min / 1000).toFixed(0)
                            const end = (xaxis.max / 1000).toFixed(0)
                            await this.chartHits(start, end, "10m")
                            return {
                                xaxis: {
                                    min: this.fromUnix(start),
                                    max: this.fromUnix(end)
                                }
                            }
                        },
                    },
                    height: 500,
                    width: "100%",
                    type: "area",
                    animations: {
                        enabled: true,
                        initialAnimation: {
                            enabled: true
                        }
                    },
                    selection: {
                        enabled: true
                    },
                    zoom: {
                        enabled: true
                    },
                    toolbar: {
                        show: true
                    },
                    stroke: {
                        show: false,
                        curve: 'smooth',
                        lineCap: 'butt',
                    },
                },
                xaxis: {
                    type: "datetime",
                    labels: {
                        show: true
                    },
                    tooltip: {
                        enabled: true
                    }
                },
                yaxis: {
                    labels: {
                        show: true
                    },
                },
                tooltip: {
                    theme: false,
                    enabled: true,
                    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
                        let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
                        const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
                        let val = series[seriesIndex][dataPointIndex];
                        if (val >= 1000) {
                            val = (val * 0.1).toFixed(0) + " milliseconds"
                        } else {
                            val = (val * 0.01).toFixed(0) + " microseconds"
                        }
                        return `<div class="chartmarker"><span>Average Response Time: </span><span class="font-3">${val}</span><span>${dt}</span></div>`
                    },
                    fixed: {
                        enabled: true,
                        position: 'topRight',
                        offsetX: -30,
                        offsetY: 40,
                    },
                    x: {
                        show: false,
                        format: 'dd MMM',
                        formatter: undefined,
                    },
                    y: {
                        formatter: undefined,
                        title: {
                            formatter: (seriesName) => seriesName,
                        },
                    },
                },
                legend: {
                    show: false,
                },
                dataLabels: {
                    enabled: false
                },
                floating: true,
                axisTicks: {
                    show: true
                },
                axisBorder: {
                    show: false
                },
                fill: {
                    colors: ["#48d338"],
                    opacity: 1,
                    type: 'solid'
                },
                stroke: {
                    show: true,
                    curve: 'smooth',
                    lineCap: 'butt',
                    colors: ["#3aa82d"],
                }
            },
            series: [{
                data: []
            }],
            heatmap_data: [],
            config: {
                enableTime: true
            },
        }
    },
    computed: {
      service () {
        return this.$store.getters.serviceByAll(this.id)
      },
      core () {
        return this.$store.getters.core
      },
      uptime_data() {
          const data = this.timedata.series.filter(g => g.online)
          const offData = this.timedata.series.filter(g => !g.online)
          let arr = [];
          data.forEach((d) => {
            arr.push({
              name: "Online", data: {
                x: 'Online',
                y: [
                  new Date(d.start).getTime(),
                  new Date(d.end).getTime()
                ],
                fillColor: '#0db407'
              }
            })
          })
          offData.forEach((d) => {
            arr.push({
              name: "offline", data: {
                x: 'Offline',
                y: [
                  new Date(d.start).getTime(),
                  new Date(d.end).getTime()
                ],
                fillColor: '#b40707'
              }
            })
          })
          return arr
      },
      rangeSeries() {
        return [{data: this.time_chart_data}]
      },
    },
    watch: {
      service: function(n, o) {
        this.chartHits()
        this.fetchUptime()
      },
      load_timedata: function(n, o) {
        this.chartHits()
      }
    },
    async created() {
        this.id = this.$route.params.id;
    },
    methods: {
      async fetchUptime() {
         this.timedata = await Api.service_uptime(this.id)
         this.load_timedata = true
        },
        async get() {
            const s = this.$store.getters.serviceByAll(this.id)
            window.console.log("service: ", s)
            this.getService(this.service)
            this.messages = this.$store.getters.serviceMessages(this.service.id)
        },
        messageInRange(message) {
            const start = this.isBetween(new Date(), message.start_on)
            const end = this.isBetween(message.end_on, new Date())
            return start && end
        },
        async getService() {
            await this.chartHits()
            await this.serviceFailures()
        },
        async serviceFailures() {
            let tt = this.startEndTimes()
            this.failures = await Api.service_failures(this.service.id, tt.start, tt.end)
        },
        async chartHits(start=0, end=99999999999, group="30m") {
            let tt = {};
            if (start === 0) {
                tt = this.startEndTimes()
            } else {
                tt = {start, end}
            }

            this.data = await Api.service_hits(this.service.id, tt.start, tt.end, group, false)
            if (this.data.length === 0 && group !== "1h") {
                await this.chartHits("1h")
            }
            this.series = [{
                name: this.service.name,
                ...this.convertToChartData(this.data)
            }]
            this.ready = true
        },
        startEndTimes() {
            const start = this.toUnix(this.service.stats.first_hit)
            const end = this.toUnix(new Date())
            return {start, end}
        }
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
