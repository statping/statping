<template>
    <div class="mb-4">
        <div class="card">
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
                <div class="col-10 text-truncate">
                    <span class="d-none d-md-inline">
                        {{smallText(service)}}
                    </span>
                </div>
                <div class="col-sm-12 col-md-2">
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
        }
    },
    methods: {
        smallText(s) {
            if (s.online) {
                return `Online, last checked ${this.ago(this.parseTime(s.last_success))}`
            } else {
                return `Offline, last error: ${s.last_failure.issue} ${this.ago(this.parseTime(s.last_failure.created_at))}`
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
