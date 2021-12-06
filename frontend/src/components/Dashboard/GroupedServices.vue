<template>
  <div class="row">
    <h5 v-if="group.name && group_services" class="h5 col-12 mb-3 mt-2 text-dim">
      <font-awesome-icon @click="toggle" :icon="expanded ? 'minus' : 'plus'" class="pointer mr-3"/> {{group.name}}
      <span class="badge badge-success text-uppercase float-right ml-2">{{services_online.length}} {{$t('online')}}</span>
      <span v-if="services_online.services_offline > 0" class="badge badge-danger text-uppercase float-right">
        {{services_offline.length}} {{$t('offline')}}
      </span>
    </h5>
    <div class="col-12 col-md-4" v-if="expanded" v-for="service in group_services">
      <ServiceInfo :service="service" />
    </div>
  </div>
</template>

<script>
const ServiceInfo = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/ServiceInfo')


export default {
name: "GroupedServices",
  components: {
    ServiceInfo
  },
  data() {
    return {
      expanded: true
    }
  },
  props: {
    group: {
      required: true,
      type: Object,
    }
  },
  computed: {
    services_online() {
      return this.$store.getters.servicesInGroup(this.group.id).filter((s) => s.online)
    },
    services_offline() {
      return this.$store.getters.servicesInGroup(this.group.id).filter((s) => !s.online)
    },
    group_services() {
      return this.$store.getters.servicesInGroup(this.group.id)
    },
  },
  methods: {
    toggle() {
      this.expanded = !this.expanded
    },
    dashboard_cookies() {
      const data = [{group: 5, show: false}]
      if (!this.$cookies.isKey("statping_layout")) {
        this.$cookies.set("statping_layout", JSON.stringify(data))
      }
    }
  }
}
</script>
