import Vue from "vue";
const { startOfToday, startOfMonth, lastDayOfMonth, subSeconds, getUnixTime, fromUnixTime, differenceInSeconds, formatDistance, addMonths, isWithinInterval } = require('date-fns')
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
    startToday() {
      return startOfToday()
    },
      secondsHumanize (val) {
        return `${val} ${this.$t('second', val)}`
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
    isZero(val) {
      return getUnixTime(parseISO(val)) <= 0
    },
    smallText(s) {
      const incidents = s.incidents
      if (s.online) {
        return `Online, checked ${this.ago(s.last_success)} ago`
      } else {
        const last = s.last_failure
        if (last) {
          return `Offline, last error: ${last} ${this.ago(last.created_at)}`
        }
        if (this.isZero(s.last_success)) {
          return `Service has never been online`
        }
        return `Service has been offline for ${this.ago(s.last_success)}`
      }
    },
    toUnix(val) {
      return getUnixTime(val)
    },
    fromUnix(val) {
      return fromUnixTime(val)
    },
    isBetween(t, start, end) {
      return isWithinInterval(t, {start: parseISO(start), end: parseISO(end)})
    },
    hour() {
      return 3600
    },
    day() {
      return 3600 * 24
    },
    maxDate() {
      return new Date(8640000000000000)
    },
    copy(txt) {
      this.$copyText(txt).then(function (e) {
        alert('Copied: \n'+txt)
        console.log(e)
      });
    },
    serviceLink(service) {
      if (service.permalink) {
        service = this.$store.getters.serviceByPermalink(service.permalink)
      }
      if (service===undefined || this.isEmptyObject(service)) {
        return `/service/0`
      }
      let link = service.permalink ? service.permalink : service.id
      return `/service/${link}`
    },
    isEmptyObject(obj) {
      return Object.keys(obj).length === 0 && obj.constructor === Object
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
        case "fa dot-circle":
          return ["fa", "dot-circle"]
        case "fas envelope-square":
          return ["fas", "envelope-square"]
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
      if (val >= 1000) {
          return Math.round(val / 1000) + " ms"
      }
        return val + " μs"
    },
    firstDayOfMonth(date) {
      return startOfMonth(date)
    },
    lastDayOfMonth(month) {
      return lastDayOfMonth(month)
    },
    addMonths(date, amount) {
      return addMonths(date, amount)
    }
  }
});
