<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <div class="col-12 mb-4">

            <span class="mt-3 mb-3 text-white d-md-none btn d-block d-md-none text-uppercase" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                {{service.online ? $t('online') : $t('offline')}}
            </span>

            <h4 class="mt-2">
                <router-link to="/" class="text-black-50 text-decoration-none">{{core.name}}</router-link> - <span class="text-muted">{{service.name}}</span>
                <span class="badge float-right d-none d-md-block text-uppercase" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>
            </h4>

            <ServiceTopStats :service="service"/>

            <MessageBlock v-for="message in messagesInRange" v-bind:key="message.id" :message="message"/>

            <div class="row mt-5 mb-4">
                <div class="col-12 col-md-5 font-2 mb-3 mb-md-0">
                    <flatPickr :disabled="loading" @on-change="onnn" v-model="start_time" :config="{ enableTime: true, altInput: true, altFormat: 'Y-m-d h:i K', maxDate: new Date() }" type="text" class="btn btn-white text-left" required />
                    <small class="d-block">From {{this.format(new Date(start_time))}}</small>
                </div>
                <div class="col-12 col-md-5 font-2 mb-3 mb-md-0">
                    <flatPickr :disabled="loading" @on-change="onnn" v-model="end_time" :config="{ enableTime: true, altInput: true, altFormat: 'Y-m-d h:i K', maxDate: new Date()}" type="text" class="btn btn-white text-left" required />
                    <small class="d-block">To {{this.format(new Date(end_time))}}</small>
                </div>
                <div class="col-12 col-md-2">
                    <select :disabled="loading" @change="chartHits" v-model="group" class="form-control">
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

            <AdvancedChart :group="group" :updated="updated_chart" :start="start_time.toString()" :end="end_time.toString()" :service="service"/>

            <div v-if="!loading" class="col-12">
                <apexchart width="100%" height="120" type="rangeBar" :options="timeRangeOptions" :series="uptime_data"></apexchart>
            </div>

            <div class="service-chart-heatmap mt-5 mb-4">
                <ServiceHeatmap :service="service"/>
            </div>

        </div>
    </div>

</template>

<script>
  import Api from "../API"
  const MessageBlock = () => import('@/components/Index/MessageBlock')
  const ServiceFailures = () => import('@/components/Service/ServiceFailures')
  const Checkin = () => import('@/forms/Checkin')
  const ServiceHeatmap = () => import('@/components/Service/ServiceHeatmap')
  const ServiceTopStats = () => import('@/components/Service/ServiceTopStats')
  const AdvancedChart = () => import('@/components/Service/AdvancedChart')

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
            tab: "failures",
            authenticated: false,
            ready: true,
            group: "1h",
            data: null,
            uptime_data: null,
            loading: true,
            messages: [],
            failures: [],
            start_time: this.nowSubtract(84600 * 30),
            end_time: this.nowSubtract(0),
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
              height: 120,
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
                        window.console.log('updated')
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
    computed: {
      service () {
        return this.$store.getters.serviceByAll(this.id)
      },
      core () {
        return this.$store.getters.core
      },
      params () {
        return {start: this.toUnix(new Date(this.start_time)), end: this.toUnix(new Date(this.end_time))}
      },
      id () {
          return this.$route.params.id;
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
    watch: {
      service: function(n, o) {
        this.onnn()
      },
      load_timedata: function(n, o) {
        this.onnn()
      }
    },
  async mounted() {
    if (!this.$store.getters.service) {
      const s = await Api.service(this.id)
      this.$store.commit('setService', s)
    }
  },
    methods: {
      async updated_chart(start, end) {
        this.start_time = start
        this.end_time = end
        this.loading = false
      },
      async onnn() {
        this.loading = true
        await this.chartHits()
        await this.fetchUptime()
        this.loading = false
      },
      async fetchUptime() {
         const uptime = await Api.service_uptime(this.id, this.params.start, this.params.end)
        window.console.log(uptime)
        this.uptime_data = this.parse_uptime(uptime)
      },
      parse_uptime(timedata) {
        const data = timedata.series.filter((g) => g.online) || []
        const offData = timedata.series.filter((g) => !g.online) || []
        let arr = [];
        window.console.log(data)
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
        async getService() {
            await this.chartHits()
            await this.serviceFailures()
        },
        async serviceFailures() {
            this.failures = await Api.service_failures(this.service.id, this.params.start, this.params.end)
        },
        async chartHits(start=0, end=99999999999) {
            this.data = await Api.service_hits(this.service.id, this.params.start, this.params.end, this.group, false)
            if (this.data.length === 0 && this.group !== "1h") {
                this.group = "1h"
                await this.chartHits("1h")
            }
            this.ready = true
        }
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
