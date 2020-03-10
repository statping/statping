<template>
    <div class="mb-4">
        <div class="card index-chart">
            <div class="card-body">
                <div class="col-12">
                    <h4 class="mt-3">
                        <router-link :to="serviceLink(service)" :in_service="service">{{service.name}}</router-link>
                        <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online}">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    </h4>

                    <ServiceTopStats :service="service"/>

                </div>
            </div>

            <div v-observe-visibility="visibleChart" class="chart-container">
                <ServiceChart :service="service" :visible="visible"/>
            </div>

            <div class="row lower_canvas full-col-12 text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                <div class="col-md-8 col-6">
                        <div class="dropup" :class="{show: dropDownMenu}">
                              <button style="font-size: 10pt;" @focusout="dropDownMenu = false"  @click="dropDownMenu = !dropDownMenu" type="button" class="col-4 float-left btn btn-sm float-right btn-block text-white dropdown-toggle service_scale" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
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
                    <router-link :to="serviceLink(service)" class="btn btn-sm float-right dyn-dark btn-block text-white" :class="{'bg-success': service.online, 'bg-danger': !service.online}">
                        View Service</router-link>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
import ServiceChart from "./ServiceChart";
import ServiceTopStats from "@/components/Service/ServiceTopStats";

export default {
    name: 'ServiceBlock',
    components: {ServiceTopStats, ServiceChart},
    props: {
        service: {
            type: Object,
            required: true
        },
    },
    data() {
        return {
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
