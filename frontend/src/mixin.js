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
    day() { return 3600 * 24 }
  }
});
