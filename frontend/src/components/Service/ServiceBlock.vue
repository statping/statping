<template>
    <div class="mb-md-4 mb-5">
        <div class="card index-chart" :class="{'expanded-service': expanded}">
            <div class="card-body">
                <div class="col-12">
                    <h4 class="mt-3">
                        <router-link :to="serviceLink(service)" class="d-inline-block text-truncate" style="max-width: 65vw;" :in_service="service">{{service.name}}</router-link>
                        <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online}">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    </h4>

                    <ServiceTopStats :service="service"/>

                        <div v-if="expanded" class="row">
                            <Analytics title="Last Failure" level="100" value="35%" subtitle="417 Days ago"/>
                            <Analytics title="Total Failures" level="100" value="35%" subtitle="417 Days ago"/>
                            <Analytics title="Highest Latency" level="100" value="450ms" subtitle="417 Days ago"/>
                            <Analytics title="Lowest Latency" level="100" value="120ms" subtitle="417 Days ago"/>
                            <Analytics title="Total Uptime" level="100" value="35%" subtitle="850ms"/>
                            <Analytics title="Total Downtime" level="100" value="35%" subtitle="32ms"/>
                         </div>

                </div>
            </div>

            <div v-if="!expanded" v-observe-visibility="visibleChart" class="chart-container">
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

                            <span class="d-none float-right d-md-inline">
                                {{smallText(service)}}
                            </span>
                        </div>
                </div>

                <div class="col-md-4 col-6 float-right">
                    <router-link :to="serviceLink(service)" class="d-none btn btn-sm float-right dyn-dark text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                        View Service</router-link>
                    <button @click="expanded = !expanded" class="btn btn-sm float-right dyn-dark text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">View Service</button>
                </div>

                <div v-if="expanded" class="row">
                    <Analytics title="Last Failure" value="417 Days ago"/>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
import Analytics from './Analytics';
import ServiceChart from "./ServiceChart";
import ServiceTopStats from "@/components/Service/ServiceTopStats";

export default {
    name: 'ServiceBlock',
    components: { Analytics, ServiceTopStats, ServiceChart},
    props: {
        service: {
            type: Object,
            required: true
        },
    },
    data() {
        return {
            expanded: false,
            visible: false,
            dropDownMenu: false,
            timeframes: [
                {value: "72h", text: "3 Days", picked: true },
                {value: "24h", text: "Since Yesterday" },
                {value: "3", text: "3 Hours" },
                {value: "1m", text: "1 Month" },
                {value: "3", text: "Last 3 Months" },
            ]
        }
    },
    methods: {
        smallText(s) {
            if (s.online) {
                return `Online, last checked ${this.ago(this.parseTime(s.last_success))}`
            } else {
                const last = s.last_failure
                if (last) {
                    return `Offline, last error: ${last} ${this.ago(this.parseTime(last.created_at))}`
                }
                return `Offline`
            }
        },
        visibleChart(isVisible, entry) {
                if (isVisible && !this.visible) {
                    this.visible = true
                }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
