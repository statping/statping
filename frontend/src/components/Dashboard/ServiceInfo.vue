<template>
    <div class="dashboard_card card mb-4" :class="{'offline-card': !service.online}">
        <div class="card-header pb-1">
            <h6 v-observe-visibility="setVisible">
                <router-link :to="serviceLink(service)" class="no-decoration">{{service.name}}</router-link>
                <span class="badge float-right text-uppercase" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>
            </h6>
        </div>

        <div class="card-body">
            <div v-if="loaded" class="row pl-2">
              <div class="col-md-6 col-sm-12 pl-2 mt-2 mt-md-0 mb-3">
                  <ServiceSparkLine :title="set2_name" subtitle="Latency Last 24 Hours" :series="set2"/>
              </div>
              <div class="col-md-6 col-sm-12 pl-0 mt-4 mt-md-0 mb-3">
                  <ServiceSparkLine :title="set1_name" subtitle="Latency Last 7 Days" :series="set1"/>
              </div>
              <ServiceEvents :service="service"/>
            </div>
              <div v-else class="row mb-5">
                <div class="col-12 col-md-6 text-center">
                  <font-awesome-icon icon="circle-notch" class="text-dim" size="2x" spin/>
                </div>
                <div class="col-12 col-md-6 text-center text-dim">
                  <font-awesome-icon icon="circle-notch" class="text-dim" size="2x" spin/>
                </div>
              </div>
        </div>
        <div class="card-footer">

          <div class="row">
          <div class="col-5 pr-0">
              <span class="small text-dim">{{ hoverbtn }}</span>
          </div>

            <div class="col-7 pr-2 pl-0">
              <div class="btn-group float-right">
                  <button @click="$router.push({path: `/dashboard/service/${service.id}/incidents`, params: {id: service.id}})" @mouseleave="unsetHover" @mouseover="setHover($t('incidents'))" class="btn btn-sm btn-white incident">
                    <font-awesome-icon icon="bullhorn"/>
                  </button>
                  <button @click="$router.push({path: `/dashboard/service/${service.id}/checkins`, params: {id: service.id}})" @mouseleave="unsetHover" @mouseover="setHover($t('checkins'))" class="btn btn-sm btn-white checkins">
                    <font-awesome-icon icon="calendar-check"/>
                  </button>
                  <button @click="$router.push({path: `/dashboard/service/${service.id}/failures`, params: {id: service.id}})" @mouseleave="unsetHover" @mouseover="setHover($t('failures'))" class="btn btn-sm btn-white failures">
                    <font-awesome-icon icon="exclamation-triangle"/> <span v-if="service.stats.failures !== 0" class="badge badge-danger ml-1">{{service.stats.failures}}</span>
                  </button>
              </div>
            </div>

          </div>

        </div>

        <span v-for="(failure, index) in failures" v-bind:key="index" class="alert alert-light">
            {{ $t('failed') }} {{failure.created_at}}<br>
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
              hovered: false,
              hoverbtn: "",
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
    mounted() {
      this.unsetHover()
    },
    methods: {
      setHover(name) {
        this.hoverbtn = name
      },
        unsetHover() {
          this.hoverbtn = this.$t('uptime', [this.service.online_7_days])
        },
        async setVisible(isVisible, entry) {
          if (isVisible && !this.visible) {
            await this.loadInfo()
            await this.getUptime()
            this.visible = true
          }
        },
        async getUptime() {
          const end = this.endOf("day", this.now())
          const start = this.beginningOf("day", this.nowSubtract(3 * 86400))
          this.uptime = await Api.service_uptime(this.service.id, this.toUnix(start), this.toUnix(end))
        },
        async loadInfo() {
          this.set1 = await this.getHits(86400 * 7, "12h")
          this.set1_name = this.calc(this.set1)
          this.set2 = await this.getHits(86400, "60m")
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
          async getHits(seconds, group) {
              let start = this.nowSubtract(seconds)
              let end = this.endOf("today")
              const startEnd = this.startEndParams(start, end, group)
              const fetched = await Api.service_hits(this.service.id, startEnd.start, startEnd.end, group, true)
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
