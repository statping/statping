<template>
    <div v-if="ready" class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <div class="col-12 mb-4">

            <span class="mt-3 mb-3 text-white d-md-none btn d-block d-md-none" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                {{service.online ? "ONLINE" : "OFFLINE"}}
            </span>

            <h4 class="mt-2">
                <router-link to="/">{{$store.getters.core.name}}</router-link> - {{service.name}}
                <span class="badge float-right d-none d-md-block" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h4>

            <ServiceTopStats :service="service"/>

            <div v-for="(message, index) in messages" v-if="messageInRange(message)">
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

            <div class="service-chart-heatmap mt-3 mb-4">
                <ServiceHeatmap :service="service"/>
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
  import flatPickr from 'vue-flatpickr-component';
  import 'flatpickr/dist/flatpickr.css';

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
            id: null,
            tab: "failures",
            service: {},
            authenticated: false,
            ready: false,
            data: null,
            messages: [],
            failures: [],
            start_time: "",
            end_time: "",
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
            heatmap_data: [],
            config: {
                enableTime: true
            },
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
            await this.serviceFailures()
        },
        async serviceFailures() {
            this.failures = await Api.service_failures(this.service.id, this.now() - 3600, this.now(), 15)
        },
        async chartHits() {
            const start = this.nowSubtract((3600 * 24) * 3)
            this.start_time = start
            this.end_time = new Date()
            this.data = await Api.service_hits(this.service.id, this.toUnix(start), this.toUnix(new Date()), "30m", false)
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
