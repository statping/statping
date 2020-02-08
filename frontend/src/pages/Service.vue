<template>
    <div v-if="ready" class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <div class="col-12 mb-4">

            <span class="mt-3 mb-3 text-white d-md-none btn d-block d-md-none" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                {{service.online ? "ONLINE" : "OFFLINE"}}
            </span>

            <h4 class="mt-2"><router-link to="/">{{$store.getters.core.name}}</router-link> - {{service.name}}
                <span class="badge float-right d-none d-md-block" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h4>

            <div class="row stats_area mt-5 mb-5">
                <div class="col-4">
                    <span class="lg_number">{{service.online_24_hours}}%</span>
                    Online last 24 Hours
                </div>
                <div class="col-4">
                    <span class="lg_number">31ms</span>
                    Average Response
                </div>
                <div class="col-4">
                    <span class="lg_number">85.70%</span>
                    Total Uptime
                </div>
            </div>

            <div v-for="(message, index) in messages" v-if="messageInRange(message)">
                <MessageBlock :message="message"/>
            </div>

            <div v-if="series" class="service-chart-container">
                <apexchart width="100%" height="420" type="area" :options="chartOptions" :series="series"></apexchart>
            </div>

            <div v-if="series" class="service-chart-heatmap">
                <apexchart width="100%" height="215" type="heatmap" :options="chartOptions" :series="series"></apexchart>
            </div>

            <form id="service_date_form" class="col-12 mt-2 mb-3">
                <input type="text" class="d-none" name="start" id="service_start" data-input>
                <span data-toggle title="toggle" id="start_date" class="text-muted small float-left pointer mt-2">Thu, 09 Jan 2020 to Thu, 16 Jan 2020</span>
                <button type="submit" class="btn btn-light btn-sm mt-2">Set Timeframe</button>
                <input type="text" class="d-none" name="end" id="service_end" data-input>

                <div id="start_container"></div>
                <div id="end_container"></div>
            </form>

            <nav class="nav nav-pills flex-column flex-sm-row mt-3" id="service_tabs">
                <a @click="tab='failures'" class="flex-sm-fill text-sm-center nav-link active">Failures</a>
                <a @click="tab='incidents'" class="flex-sm-fill text-sm-center nav-link">Incidents</a>
                <a @click="tab='checkins'" v-if="$store.getters.token" class="flex-sm-fill text-sm-center nav-link">Checkins</a>
                <a @click="tab='response'" v-if="$store.getters.token" class="flex-sm-fill text-sm-center nav-link">Response</a>
            </nav>


            <div class="tab-content">
                <div class="tab-pane fade active show">
                    <div class="list-group mt-3 mb-4">

                        <div v-for="(failure, index) in failures" :key="index" class="mb-2 list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">{{failure.issue}}</h5>
                                <small>{{failure.created_at | moment("dddd, MMMM Do YYYY")}}</small>
                            </div>
                            <p class="mb-1">{{failure.issue}}</p>
                        </div>
                    </div>
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
  import Api from "../components/API"
  import MessageBlock from '../components/Index/MessageBlock';
  import Checkin from "../forms/Checkin";

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
      MessageBlock,
    Checkin
  },
  data () {
    return {
      id: null,
      tab: "failures",
      service: {},
      authenticated: false,
      ready: false,
      data: null,
        messages: [],
      failures: [],
      chartOptions: {
        chart: {
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
            enabled: false
          },
          zoom: {
            enabled: false
          },
          toolbar: {
            show: false
          },
        },
        grid: {
          show: true,
          padding: {
            top: 0,
            right: 0,
            bottom: 0,
            left: 0,
          }
        },
        xaxis: {
          type: "datetime",
          ...axisOptions
        },
        yaxis: {
          ...axisOptions
        },
        tooltip: {
          enabled: false,
          marker: {
            show: false,
          },
          x: {
            show: false,
          }
        },
        legend: {
          show: false,
        },
        dataLabels: {
          enabled: false
        },
        floating: true,
        axisTicks: {
          show: false
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
        heatmap_data: []
    }
  },
    async mounted() {
      const id = this.$attrs.id
        this.id = id
      let service;
      if (this.isInt(id)) {
        service = this.$store.getters.serviceById(id)
      } else {
        service = this.$store.getters.serviceByPermalink(id)
      }
        this.service = service
        this.getService(service)
        this.messages = this.$store.getters.serviceMessages(service.id)
    },
  methods: {
      messageInRange(message) {
          const start = this.isBetween(new Date(), message.start_on)
          const end = this.isBetween(message.end_on, new Date())
          return start && end
      },
    async getService(s) {
        await this.chartHits()
        await this.heatmapData()
        await this.serviceFailures()
    },
    async serviceFailures() {
      this.failures = await Api.service_failures(this.service.id, 0, 99999999999)
    },
    async chartHits() {
      this.data = await Api.service_hits(this.service.id, 0, 99999999999, "hour")
      this.series = [{
        name: this.service.name,
        ...this.data
      }]
      this.ready = true
    },
      async heatmapData() {
          this.data = await Api.service_heatmap(this.service.id, 0, 99999999999, "hour")
          this.series = [{
              name: this.service.name,
              ...this.data
          }]
          this.ready = true
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
