import Vue from "vue";
const { zonedTimeToUtc, utcToZonedTime, lastDayOfMonth, subSeconds, parse, getUnixTime, fromUnixTime, differenceInSeconds, formatDistance } = require('date-fns')
import formatDistanceToNow from 'date-fns/formatDistanceToNow'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import addSeconds from 'date-fns/addSeconds'

export default Vue.mixin({
  methods: {
    now() {
      return new Date()
    },
    isNumeric: function (n) {
      return !isNaN(parseFloat(n)) && isFinite(n);
    },
    current() {
      return parseISO(new Date())
    },
      secondsHumanize (val) {
        const t2 = addSeconds(new Date(0), val)
          if (val >= 60) {
              let minword = "minute"
              if (val >= 120) {
                  minword = "minutes"
              }
              return format(t2, "m '"+minword+"' s 'seconds'")
          }
        return format(t2, "s 'seconds'")
      },
    utc(val) {
      return new Date.UTC(val)
    },
    ago(t1) {
      return formatDistanceToNow(parseISO(t1))
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
    format(val, type="EEEE, MMM do h:mma") {
      return format(val, type)
    },
    niceDate(val) {
      return format(parseISO(val), "EEEE, MMM do h:mma")
    },
      parseISO(v) {
        return parseISO(v)
      },
    toUnix(val) {
      return getUnixTime(val)
    },
    fromUnix(val) {
      return fromUnixTime(val)
    },
    isBetween(t1, t2) {
      return differenceInSeconds(parseISO(t1), parseISO(t2)) >= 0
    },
    hour() {
      return 3600
    },
    day() {
      return 3600 * 24
    },
    copy(txt) {
      this.$copyText(txt).then(function (e) {
        alert('Copied: \n'+txt)
        console.log(e)
      });
    },
    serviceLink(service) {
      if (service.permalink) {
        service = this.$store.getters.serviceByPermalink(service)
      }
      if (service===undefined) {
        return `/service/0`
      }
      let link = service.permalink ? service.permalink : service.id
      return `/service/${link}`
    },
    isInt(n) {
      return n % 1 === 0;
    },
    isAdmin() {
      return this.$store.state.admin
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
    toBarData(data = []) {
      let newSet = [];
      data.forEach((f) => {
        newSet.push([this.toUnix(this.parseISO(f.timeframe)), f.amount])
      })
      return newSet
    },
    convertToChartData(data = [], multiplier=1, asInt=false) {
      if (!data) {
        return {data: []}
      }
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
    },
      humanTime(val) {
        if (val >= 10000) {
            return Math.floor(val / 10000) + "ms"
        }
          return Math.floor(val / 1000) + "μs"
      },
    lastDayOfMonth(month) {
      return new Date(Date.UTC(new Date().getUTCFullYear(), month + 1, 0))
    },
    firstDayOfMonth(month) {
      return new Date(Date.UTC(new Date().getUTCFullYear(), month, 1)).getUTCDate()
    }
  }
});
