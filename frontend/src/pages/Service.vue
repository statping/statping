<template>
    <div class="container col-md-7 col-sm-12 mt-md-5">
      <div v-if="!ready" class="row mt-5">
        <div class="col-12 text-center">
          <font-awesome-icon icon="circle-notch" size="3x" spin/>
        </div>
        <div class="col-12 text-center mt-3 mb-3">
          <span class="text-muted">Loading Service</span>
        </div>
      </div>

        <div v-if="ready && service" class="col-12 mb-4">
            <span class="mt-3 mb-3 text-white d-md-none btn d-block d-md-none text-uppercase" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                {{service.online ? $t('online') : $t('offline')}}
            </span>

            <span class="mt-2 font-3">
                <router-link to="/" class="text-black-50 text-decoration-none">{{core.name}}</router-link> - <span class="text-muted">{{service.name}}</span>
                <span class="badge float-right d-none d-md-block text-uppercase" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>
            </span>

            <ServiceTopStats v-if="loaded" :service="service"/>

            <MessageBlock v-if="loaded" v-for="message in messagesInRange" v-bind:key="message.id" :message="message"/>

            <div class="card text-black-50 bg-white mt-3">
                <div class="card-header text-capitalize">Timeframe</div>
                <div class="card-body pb-4">
                    <div class="row">
                        <div class="col">
                            <flatPickr :disabled="!loaded" @on-change="reload" v-model="start_time" :config="{ wrap: true, allowInput: true, enableTime: true, dateFormat: 'Z', altInput: true, altFormat: 'Y-m-d h:i K', maxDate: this.endOf('today') }" type="text" class="form-control text-left" required />
                            <small class="d-block">From {{this.format(new Date(start_time))}}</small>
                        </div>
                        <div class="col">
                            <flatPickr :disabled="!loaded" @on-change="reload" v-model="end_time" :config="{ wrap: true, allowInput: true, enableTime: true, dateFormat: 'Z', altInput: true, altFormat: 'Y-m-d h:i K', maxDate: this.endOf('today') }" type="text" class="form-control text-left" required />
                            <small class="d-block">To {{this.format(new Date(end_time))}}</small>
                        </div>
                        <div class="col">
                            <select :disabled="!loaded" @change="chartHits(service)" v-model="group" class="form-control">
                                <option value="1m">1 Minute</option>
                                <option value="5m">5 Minutes</option>
                                <option value="15m">15 Minute</option>
                                <option value="30m">30 Minutes</option>
                                <option value="1h">1 Hour</option>
                                <option value="3h">3 Hours</option>
                                <option value="6h">6 Hours</option>
                                <option value="12h">12 Hours</option>
                                <option value="24h">1 Day</option>
                                <option value="168h">7 Days</option>
                                <option value="360h">15 Days</option>
                            </select>
                            <small class="d-block d-md-none d-block">Increment Timeframe</small>
                        </div>
                    </div>

                </div>
            </div>

            <div class="card text-black-50 bg-white mt-3 mb-3">
                <div class="card-header text-capitalize">Service Latency</div>
                <div v-if="loaded" class="card-body">
                    <div class="row">
                      <AdvancedChart :group="group" :updated="updated_chart" :start="start_time.toString()" :end="end_time.toString()" :service="service"/>
                    </div>
                  <div>
                    <FailuresBarChart :service="service" :start="start_time.toString()" :end="end_time.toString()" :group="group"/>
                  </div>

                </div>
              <div v-else class="row mt-3 mb-3">
                <div class="col-12 text-center">
                  <font-awesome-icon icon="circle-notch" size="3x" spin/>
                </div>
              </div>

            </div>

            <div class="card text-black-50 bg-white mb-3">
                <div class="card-header text-capitalize">Service Failures</div>
                <div class="card-body">
                    <div class="service-chart-heatmap mt-5 mb-4">
                        <ServiceHeatmap :service="service"/>
                    </div>
                </div>
            </div>

        </div>
    </div>

</template>

<script>
  import Api from "../API"
  const MessageBlock = () => import(/* webpackChunkName: "index" */ '@/components/Index/MessageBlock')
  const ServiceFailures = () => import(/* webpackChunkName: "service" */ '@/components/Service/ServiceFailures')
  const Checkin = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Checkin')
  const ServiceHeatmap = () => import(/* webpackChunkName: "service" */ '@/components/Service/ServiceHeatmap')
  const ServiceTopStats = () => import(/* webpackChunkName: "service" */ '@/components/Service/ServiceTopStats')
  const AdvancedChart = () => import(/* webpackChunkName: "service" */ '@/components/Service/AdvancedChart')
  const FailuresBarChart = () => import(/* webpackChunkName: "service" */ '@/components/Service/FailuresBarChart')

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
      FailuresBarChart,
      AdvancedChart,
        ServiceTopStats,
        ServiceHeatmap,
        ServiceFailures,
        MessageBlock,
        Checkin,
        flatPickr
    },
    data() {
        return {
            id: null,
            tab: "failures",
            authenticated: false,
            ready: false,
            group: "15m",
            data: null,
            uptime_data: null,
            loaded: false,
            messages: [],
            failures: [],
            start_time: this.beginningOf('day', this.nowSubtract(259200 * 3)),
            end_time: this.endOf('today'),
            timedata: null,
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
              id: 'uptime',
              height: 220,
              width: '100%',
              type: 'rangeBar',
              toolbar: {
                show: false
              },
              zoom: {
                enabled: false
              }
            },
            selection: {
              enabled: true
            },
            zoom: {
              enabled: true
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
              enabled: false
            },
            tooltip: {
              enabled: false,
            },
            xaxis: {
              type: 'datetime'
            },
            yaxis: {
              show: true,
            },
            grid: {
              row: {
                colors: ['#f3f4f5', '#fff'],
                opacity: 1
              }
            }
          },
            chartOptions: {
              noData: {
                text: "Loading...",
                align: 'center',
                verticalAlign: 'middle',
                offsetX: 0,
                offsetY: -20,
                style: {
                  color: "#bababa",
                  fontSize: '27px'
                }
              },
                chart: {
                  id: 'mainchart',
                    events: {
                      dataPointSelection: (event, chartContext, config) => {
                        window.console.log('slect')
                        window.console.log(event)
                      },
                      updated: (chartContext, config) => {
                        window.console.log('updated chart')
                      },
                        beforeZoom: (chartContext, { xaxis }) => {
                            const start = (xaxis.min / 1000).toFixed(0)
                            const end = (xaxis.max / 1000).toFixed(0)
                          this.start_time = this.fromUnix(start)
                          this.end_time = this.fromUnix(end)
                            return {
                                xaxis: {
                                    min: this.fromUnix(start),
                                    max: this.fromUnix(end)
                                }
                            }
                        },
                      scrolled: (chartContext, { xaxis }) => {
                        window.console.log(xaxis)
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
                        curve: 'stepline',
                        lineCap: 'butt',
                    },
                },
              xaxis: {
                type: "datetime",
                labels: {
                  format: 'MM yyyy'
                },
                tooltip: {
                  enabled: false
                }
              },
              yaxis: {
                labels: {
                  show: true
                },
              },
              markers: {
                size: 0,
                strokeWidth: 0,
                hover: {
                  size: undefined,
                  sizeOffset: 0
                }
              },
                tooltip: {
                    theme: false,
                    enabled: true,
                    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
                        let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
                        const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
                        let val = series[seriesIndex][dataPointIndex];
                        if (val >= 10000) {
                            val = Math.round(val / 1000) + " ms"
                        } else {
                            val = val + " μs"
                        }
                        return `<div class="chartmarker"><span>Response Time: </span><span class="font-3">${val}</span><span>${dt}</span></div>`
                    },
                    fixed: {
                        enabled: true,
                        position: 'topRight',
                        offsetX: -30,
                        offsetY: 40,
                    },
                    x: {
                        show: true,

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
    watch: {
      '$route': 'fetchData'
    },
    computed: {
      core () {
        return this.$store.getters.core
      },
      service() {
        return this.$store.getters.serviceByAll(this.$route.params.id)
      },
      params () {
        return {start: this.toUnix(new Date(this.start_time)), end: this.toUnix(new Date(this.end_time))}
      },
      uptimeSeries () {
        return this.timedata.series
      },
      mainChart () {
        return [{
          name: this.service.name,
          ...this.convertToChartData(this.data)
        }]
      },
      messagesInRange() {
        return this.$store.getters.serviceMessages(this.service.id).filter(m => this.inRange(m))
      },
    },
    created() {
      this.fetchData()
    },
    mounted() {
      this.fetchData()
    },
    methods: {
      fetchData () {
        if (!this.$route.params.id) {
          this.ready = false
          return
        }
        this.reload()
        this.ready = true
        this.loaded = true
      },
      async reload() {
        if (!this.ready || !this.service) {
          return
        }
        await this.chartHits()
        await this.chartFailures()
        await this.fetchUptime()
      },
      async updated_chart(start, end) {
        this.loaded = false
        this.start_time = start
        this.end_time = end
        this.loaded = true
      },
      async fetchUptime() {
        const uptime = await Api.service_uptime(this.service.id, this.params.start, this.params.end)
        this.uptime_data = this.parse_uptime(uptime)
      },
      parse_uptime(timedata) {
        if (!timedata.series) {
          return []
        }
        const data = timedata.series.filter((g) => g.online) || []
        const offData = timedata.series.filter((g) => !g.online) || []
        let arr = [];
        if (data) {
          data.forEach((d) => {
            arr.push({
              x: 'Online',
              y: [
                new Date(d.start).getTime(),
                new Date(d.end).getTime()
              ],
              fillColor: '#0db407'
            })
          })
        }
        if (offData) {
          offData.forEach((d) => {
            arr.push({
              x: 'Offline',
              y: [
                new Date(d.start).getTime(),
                new Date(d.end).getTime()
              ],
              fillColor: '#b40707'
            })
          })
        }
        return [{data: arr}]
      },
      inRange(message) {
        return this.isBetween(this.now(), message.start_on, message.start_on === message.end_on ? this.maxDate().toISOString() : message.end_on)
      },
      async chartHits(start=0, end=99999999999) {
        this.data = await Api.service_hits(this.service.id, this.params.start, this.params.end, this.group, false)
      },
      async chartFailures(start=0, end=99999999999) {
        this.failures_data = await Api.service_failures_data(this.service.id, this.params.start, this.params.end, this.group, true)
      }
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
