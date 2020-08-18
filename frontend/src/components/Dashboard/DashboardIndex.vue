<template>
    <div class="col-12 mt-4 mt-md-3">

        <div class="row stats_area mb-5">
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{$store.getters.services.length}}</span>
                <span class="font-2">{{ $t('dashboard.total_services') }}</span>
            </div>
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{failuresLast24Hours()}}</span>
                <span class="font-2">{{ $t('dashboard.failures_24_hours') }}</span>
            </div>
            <div class="col-4">
                <span class="font-6 font-weight-bold d-block">{{$store.getters.onlineServices(true).length}}</span>
                <span class="font-2">{{ $t('dashboard.online_services') }}</span>
            </div>
        </div>

        <div class="col-12" v-if="services.length === 0">
            <div class="alert alert-dark d-block">
                You currently don't have any services!
                <router-link v-if="$store.state.admin" to="/dashboard/create_service" class="btn btn-sm btn-success float-right">
                    <font-awesome-icon icon="plus"/>  Create
                </router-link>
            </div>
        </div>

      <div v-for="message in messagesInRange" class="bg-light shadow-sm p-3 pr-4 pl-4 col-12 mb-4">
        <font-awesome-icon icon="calendar" class="mr-3" size="1x"/> {{message.description}}
        <span class="d-block small text-muted mt-3">
        Starts at <strong>{{niceDate(message.start_on)}}</strong> till <strong>{{niceDate(message.end_on)}}</strong>
        ({{dur(parseISO(message.start_on), parseISO(message.end_on))}})
      </span>
      </div>


        <div v-for="(service, index) in services" class="service_block" v-bind:key="index">
            <ServiceInfo :service=service />
        </div>
    </div>
</template>

<script>
  import isAfter from "date-fns/isAfter";
  import parseISO from "date-fns/parseISO";
  import isBefore from "date-fns/isBefore";

  const ServiceInfo = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/ServiceInfo')

  export default {
      name: 'DashboardIndex',
      components: {
          ServiceInfo
      },
    data() {
        return {
          visible: false
        }
    },
      computed: {
        messagesInRange() {
          return this.$store.getters.globalMessages.filter(m => this.isAfter(this.now(), m.start_on) && this.isBefore(this.now(), m.end_on))
        },
          services() {
              return this.$store.getters.services
          }
      },
      methods: {

          failuresLast24Hours() {
              let total = 0;
              this.services.map((s) => {
                  total += s.failures_24_hours
              })
              return total
          },

      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
