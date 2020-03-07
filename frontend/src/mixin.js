import Vue from "vue";
const { zonedTimeToUtc, utcToZonedTime, lastDayOfMonth, subSeconds, parse, parseISO, getUnixTime, fromUnixTime, format, differenceInSeconds, formatDistanceToNow, formatDistance } = require('date-fns')

export default Vue.mixin({
  methods: {
    now() {
      return new Date()
    },
    current() {
      return parse(new Date())
    },
    utc(val) {
      return fromUnixTime(this.toUnix(val) + val.getTimezoneOffset() * 60 * 1000)
    },
    ago(t1) {
      return formatDistanceToNow(t1)
    },
      daysInMonth(t1) {
          return lastDayOfMonth(t1)
      },
    nowSubtract(seconds) {
      return subSeconds(new Date(), seconds)
    },
    dur(t1, t2) {
      return formatDistance(t1, t2)
    },
    niceDate(val) {
      return this.parseTime(val).format('LLLL')
    },
    parseTime(val) {
      return parseISO(val)
    },
    toLocal(val, suf = 'at') {
      const t = this.parseTime(val)
      return format(t, `EEEE, MMM do h:mma`)
    },
    toUnix(val) {
      return getUnixTime(val)
    },
    fromUnix(val) {
      return fromUnixTime(val)
    },
    isBetween(t1, t2) {
      return differenceInSeconds(parseISO(t1), parseISO(t2)) > 0
    },
    hour() {
      return 3600
    },
    day() {
      return 3600 * 24
    },
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
    },
    loggedIn() {
      const core = this.$store.getters.core
      return core.logged_in === true
    },
    iconName(name) {
      switch (name) {
        case "fas fa-terminal":
          return "terminal"
        case "fab fa-discord":
          return ["fab", "discord"]
        case "far fa-envelope":
          return "envelope"
        case "far fa-bell":
          return "bell"
        case "fas fa-mobile-alt":
          return "mobile"
        case "fab fa-slack":
          return ["fab", "slack-hash"]
        case "fab fa-telegram-plane":
          return ["fab", "telegram-plane"]
        case "far fa-comment-alt":
          return "comment"
        case "fas fa-code-branch":
          return "code-branch"
        case "csv":
          return "file"
        case "docker":
          return ["fab", "docker"]
        case "traefik":
          return "server"
        default:
          return "bars"
      }
    },
    convertToChartData(data = [], multiplier=1, asInt=false) {
      let newSet = [];
      data.forEach((f) => {
        let amount = f.amount * multiplier;
        if (asInt) {
          amount = amount.toFixed(0)
        }
        newSet.push({
          x: f.timeframe,
          y: amount
        })
      })
      return {data: newSet}
    }
  }
});
