import Vue from "vue";

export default Vue.mixin({
  methods: {
    now() {
      return Math.round(new Date().getTime() / 1000)
    },
      current() {
        return new Date()
      },
    ago(seconds) {
      return this.now() - seconds
    },
      niceDate(val) {
        return this.parseTime(val).format('LLLL')
      },
      parseTime(val) {
          return this.$moment(val, "YYYY-MM-DDTHH:mm:ssZ", true).local()
      },
      toUnix(val) {
         return this.$moment(val).utc().format('MM-DD-YYYY')
      },
      isBetween(t1, t2) {
          const now = this.$moment(t1).utc().valueOf()
          const sub = this.$moment(t2).utc().valueOf()
          return (now - sub) > 0
      },
    hour(){ return 3600 },
    day() { return 3600 * 24 },
    serviceLink(service) {
      if (!service) {
        return ""
      }
      if (!service.id) {
        service = this.$store.getters.serviceById(service)
      }
      let link = service.permalink ? service.permalink : service.id
      return `/service/${link}`
    },
    isInt(n) {
      return n % 1 === 0;
    }
  }
});
