<template v-if="service">
    <div class="col-12 card mb-3" style="min-height: 260px">
        <div class="card-body">
            <h5 class="card-title"><a href="service/7">{{service.name}}</a>
                <span class="badge float-right badge-success">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
            </h5>
            <div class="row">
                <div class="col-md-3 col-sm-6">
                    <ServiceSparkLine title="here" subtitle="Failures in 7 Days" :series="first"/>
                </div>
                <div class="col-md-3 col-sm-6">
                    <ServiceSparkLine title="here" subtitle="Failures Last Month" :series="second"/>
                </div>
                <div class="col-md-3 col-sm-6">
                    <ServiceSparkLine title="here" subtitle="Average Response" :series="third"/>
                </div>
                <div class="col-md-3 col-sm-6">
                    <ServiceSparkLine title="here" subtitle="Ping Time" :series="fourth"/>
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
        first: [],
        second: [],
        third: [],
        fourth: []
    }
  },
  async created() {
    this.first = await this.getFailures(7, "hour")
    this.second = await this.getFailures(30, "hour")
    this.third = await this.getHits(7, "hour")
    this.fourth = await this.getHits(30, "hour")
  },
  methods: {
    async getHits(days, group) {
      const start = this.ago(3600 * 24)
      const data = await Api.service_hits(this.service.id, start, this.now(), group)
      return [data]
    },
    async getFailures(days, group) {
      const start = this.ago(3600 * 24)
      const data = await Api.service_failures(this.service.id, start, this.now())
      return [data]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
