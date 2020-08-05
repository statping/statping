<template>
    <div class="card mb-4" :class="{'offline-card': !service.online}">
        <div class="card-header pb-1">
            <h4 v-observe-visibility="setVisible">
                <router-link :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge float-right text-uppercase" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>
            </h4>
        </div>

        <div class="card-body">
            <transition name="fade">
            <div v-if="loaded" class="row pl-2 pr-2">
                    <div class="col-md-6 col-sm-12 mt-2 mt-md-0 mb-3">
                        <ServiceSparkLine :title="set2_name" subtitle="Latency Last 24 Hours" :series="set2"/>
                    </div>
                    <div class="col-md-6 col-sm-12 mt-4 mt-md-0 mb-3">
                        <ServiceSparkLine :title="set1_name" subtitle="Latency Last 7 Days" :series="set1"/>
                    </div>

              <div class="col-12 mt-2 mt-md-0 mb-3">
                  <ServiceEvents :service="service"/>
              </div>

            </div>
              <div v-else class="row mt-5 mb-5 pt-5 pb-5">
                <div class="col-6 text-center text-muted">
                  <font-awesome-icon icon="circle-notch" size="3x" spin/>
                </div>
                <div class="col-6 text-center text-muted">
                  <font-awesome-icon icon="circle-notch" size="3x" spin/>
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
                    <span class="float-md-right">
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
  const Checkin = () => import(/* webpackChunkName: "dashboard" */ '../../forms/Checkin');
  const FormMessage = () => import(/* webpackChunkName: "dashboard" */ '../../forms/Message');
  const ServiceFailures = () => import(/* webpackChunkName: "dashboard" */ '../Service/ServiceFailures');
  const ServiceSparkLine = () => import(/* webpackChunkName: "dashboard" */ "./ServiceSparkLine");
  import Api from "../../API";
  const ServiceEvents = () => import(/* webpackChunkName: "dashboard" */ "@/components/Dashboard/ServiceEvents");

  export default {
      name: 'ServiceInfo',
      components: {
        ServiceEvents,
          Checkin,
          ServiceFailures,
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
