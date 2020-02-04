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
    duration(t1, t2) {
      const val = (this.toUnix(t1) - this.toUnix(t2))
      if (val <= 59) {
        return this.$moment.duration(val, 'seconds').get('seconds') + " seconds ago"
      }
      return this.$moment.duration(val, 'seconds').humanize();
    },
      niceDate(val) {
        return this.parseTime(val).format('LLLL')
      },
      parseTime(val) {
          return this.$moment(val, this.$moment.ISO_8601, true)
      },
    toLocal(val, suf='at') {
      return this.parseTime(val).local().format(`dddd, MMM Do \\${suf} h:mma`)
    },
      toUnix(val) {
         return this.$moment(val).utc().unix().valueOf()
      },
    fromUnix(val) {
      return this.$moment.unix(val).utc()
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
