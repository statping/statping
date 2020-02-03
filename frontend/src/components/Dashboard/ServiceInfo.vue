<template v-if="service">
    <div class="col-12 card mb-3" style="min-height: 260px">
        <div class="card-body">
            <h5 class="card-title"><router-link :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge float-right" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                    {{service.online ? "ONLINE" : "OFFLINE"}}
                </span>
            </h5>
            <div class="row">
                <div class="col-6">
                    <ServiceSparkLine :title="calc(set1)" subtitle="Last Day Latency" :series="set1"/>
                </div>
                <div class="col-6">
                    <ServiceSparkLine :title="calc(set2)" subtitle="Last 7 Days Latency" :series="set2"/>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
  import ServiceSparkLine from "./ServiceSparkLine";
  import Api from "../API";

  export default {
  name: 'ServiceInfo',
  components: {
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
        set1: [],
        set2: []
    }
  },
  async created() {
    this.set1 = await this.getHits(24, "hour")
    this.set2 = await this.getHits(24 * 7, "day")
  },
  methods: {
    async getHits(hours, group) {
      const start = this.ago(3600 * hours)
      const data = await Api.service_hits(this.service.id, start, this.now(), group)
      if (!data) {
          return [{name: "None", data: []}]
      }
        return [{name: "Latency", data: data.data}]
    },
    calc (s) {
        let data = s[0].data
        let total = 0
        data.forEach((f) => {
            total += f.y
        });
        total = total / data.length
        return total.toFixed(0) + "ms Average"
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
