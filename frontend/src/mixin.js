import Vue from "vue";

export default Vue.mixin({
  methods: {
    now() {
      return Math.round(new Date().getTime() / 1000)
    },
    ago(seconds) {
      return this.now() - seconds
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
