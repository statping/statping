<template>
    <div class="card mb-4" :class="{'offline-card': !service.online}">
        <div class="card-title px-4 pt-3">
            <h4 v-observe-visibility="setVisible">
                <router-link :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge float-right text-uppercase" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>
            </h4>
        </div>

        <div class="card-body p-3 p-md-1 pt-md-1 pb-md-1">

            <transition name="fade">
            <div v-if="loaded" class="col-12 pb-2">

                <div v-if="false" class="row mb-4 align-content-center">

                    <div v-if="!service.online" class="col-3 text-left">
                        <span css="text-danger font-5 font-weight-bold"></span>
                        <span class="font-2 d-block">Current Downtime</span>
                    </div>

                    <div v-if="service.online" class="col-3 text-left">
                        <span class="text-success font-5 font-weight-bold">
                            {{service.online_24_hours.toString()}} %
                        </span>
                        <span class="font-2 d-block">Total Uptime</span>
                    </div>

                    <div v-if="service.online" class="col-3 text-left">
                        <span class="text-success font-5 font-weight-bold">
                            0
                        </span>
                        <span class="font-2 d-block">Downtime Today</span>
                    </div>

                    <div v-if="service.online" class="col-3 text-left">
                        <span class="text-success font-5 font-weight-bold">
                            {{(uptime.uptime / 10000).toFixed(0).toString()}}
                        </span>
                        <span class="font-2 d-block">Uptime Duration</span>
                    </div>

                    <div class="col-3 text-left">
                        <span class="text-danger font-5 font-weight-bold">
                            {{service.failures_24_hours}}
                        </span>
                        <span class="font-2 d-block">Failures last 24 hours</span>
                    </div>

                </div>

                <div class="row">
                    <div class="col-md-6 col-sm-12 mt-2 mt-md-0 mb-3">
                        <ServiceSparkLine :title="set2_name" subtitle="Latency Last 24 Hours" :series="set2"/>
                    </div>
                    <div class="col-md-6 col-sm-12 mt-4 mt-md-0 mb-3">
                        <ServiceSparkLine :title="set1_name" subtitle="Latency Last 7 Days" :series="set1"/>
                    </div>
                </div>

            </div>
            </transition>
        </div>

        <div class="card-footer">
            <div class="row">

                <div class="col-12 col-md-3 mb-2 mb-md-0">
                    <router-link :to="{path: `/dashboard/service/${service.id}/incidents`, params: {id: service.id} }" class="btn btn-block btn-white text-capitalize incident">
                        {{$tc('incident', 2)}}
                    </router-link>
                </div>
                <div class="col-12 col-md-3 mb-2 mb-md-0">
                    <router-link :to="{path: `/dashboard/service/${service.id}/checkins`, params: {id: service.id} }" class="btn btn-block btn-white text-capitalize checkins">
                        {{$tc('checkin', 2)}}
                    </router-link>
                </div>
                <div class="col-12 col-md-3 mb-2 mb-md-0">
                    <router-link :to="{path: `/dashboard/service/${service.id}/failures`, params: {id: service.id} }" class="btn btn-block btn-white text-capitalize failures">
                        {{$tc('failure', 2)}} <span class="badge badge-danger float-right mt-1">{{service.stats.failures}}</span>
                    </router-link>
                </div>
                <div class="col-12 col-md-3 mb-2 mb-md-0 mt-0 mt-md-1">
                    <span class="text-black-50 float-md-right">
                        {{$t('uptime', [service.online_7_days])}}
                    </span>
                </div>

            </div>
        </div>

        <span v-for="(failure, index) in failures" v-bind:key="index" class="alert alert-light">
            Failed {{failure.created_at}}<br>
            {{failure.issue}}
        </span>

    </div>
</template>

<script>
  import Checkin from '../../forms/Checkin';
  import FormIncident from '../../forms/Incident';
  import FormMessage from '../../forms/Message';
  import ServiceFailures from '../Service/ServiceFailures';
  import ServiceSparkLine from "./ServiceSparkLine";
  import Api from "../../API";

  export default {
      name: 'ServiceInfo',
      components: {
          Checkin,
          ServiceFailures,
          FormIncident,
          FormMessage,
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
              uptime: null,
              openTab: "",
              set1: [],
              set2: [],
              loaded: false,
              set1_name: "",
              set2_name: "",
              failures: null,
            visible: false
          }
      },
    watch: {

    },
      methods: {
        async setVisible(isVisible, entry) {
          if (isVisible && !this.visible) {
            await this.loadInfo()
            await this.getUptime()
            this.visible = true
          }
        },
        async getUptime() {
          const start = this.nowSubtract(3 * 86400)
          this.uptime = await Api.service_uptime(this.service.id, this.toUnix(start), this.toUnix(this.now()))
        },
        async loadInfo() {
          this.set1 = await this.getHits(24 * 7, "6h")
          this.set1_name = this.calc(this.set1)
          this.set2 = await this.getHits(24, "1h")
          this.set2_name = this.calc(this.set2)
          this.loaded = true
        },
          Tab(name) {
              if (this.openTab === name) {
                  this.openTab = ''
                  return
              }
              this.openTab=name;
          },
        sinceYesterday(data) {
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
                  return Math.round(total) + " ms"
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
