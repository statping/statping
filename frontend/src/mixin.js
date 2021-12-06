import Vue from "vue";
const { startOfDay, startOfHour, startOfWeek, endOfMonth, endOfHour, startOfToday, startOfTomorrow, startOfYesterday, endOfYesterday, endOfTomorrow, endOfToday, endOfDay, startOfMonth, lastDayOfMonth, subSeconds, getUnixTime, fromUnixTime, differenceInSeconds, formatDistance, addMonths, addSeconds, isWithinInterval } = require('date-fns')
import formatDistanceToNow from 'date-fns/formatDistanceToNow'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import isBefore from 'date-fns/isBefore'
import isAfter from 'date-fns/isAfter'
import { roundToNearestMinutes } from 'date-fns'

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
    secondsHumanize(val) {
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
      return subSeconds(this.now(), seconds)
    },
    isAfter(date, compare) {
      return isAfter(date, parseISO(compare))
    },
    isBefore(date, compare) {
      return isBefore(date, parseISO(compare))
    },
    dur(t1, t2) {
      return formatDistance(t1, t2)
    },
    format(val, type = "EEEE, MMM do h:mma") {
      return format(val, type)
    },
    niceDate(val) {
      return format(parseISO(val), "EEEE, MMM do h:mma")
    },
    parseISO(v) {
      return parseISO(v)
    },
    round(minutes) {
      return roundToNearestMinutes(minutes)
    },
    endOf(method, val) {
      switch (method) {
        case "hour":
          return endOfHour(val)
        case "day":
          return endOfDay(val)
        case "today":
          return endOfToday()
        case "tomorrow":
          return endOfTomorrow()
        case "yesterday":
          return endOfYesterday()
        case "month":
          return endOfMonth(val)
      }
      return val
    },
    startEndParams(start, end, group) {
      start = this.beginningOf("hour", start)
      end = this.endOf("hour", end)
      return {start: this.toUnix(start), end: this.toUnix(end), group: group}
    },
    beginningOf(method, val) {
      switch (method) {
        case "hour":
          return startOfHour(val)
        case "day":
          return startOfDay(val)
        case "today":
          return startOfToday()
        case "tomorrow":
          return startOfTomorrow()
        case "yesterday":
          return startOfYesterday()
        case "week":
          return startOfWeek(val)
        case "month":
          return startOfMonth(val)
      }
      return val
    },
    isZero(val) {
      return getUnixTime(parseISO(val)) <= 0
    },
    smallText(s) {
      if (s.online) {
        return `${this.$t('service_online_check')} ${this.ago(s.last_success)} ago`
      } else {
        const last = s.last_failure
        if (last) {
          return `Offline, last error: ${last} ${this.ago(last.created_at)}`
        }
        if (this.isZero(s.last_success)) {
          return this.$t('service_never_online')
        }
        return `${this.$t('service_offline_time')} ${this.ago(s.last_success)}`
      }
    },
    round_time(frame, val) {
      switch(frame) {
        case "15m":
          return roundToNearestMinutes(val, {nearestTo: 60 * 15})
        case "30m":
          return roundToNearestMinutes(val, {nearestTo: 60 * 30})
        case "1h":
          return roundToNearestMinutes(val, {nearestTo: 3600})
        case "3h":
          return roundToNearestMinutes(val, {nearestTo: 3600 * 3})
        case "6h":
          return roundToNearestMinutes(val, {nearestTo: 3600 * 6})
        case "12h":
          return roundToNearestMinutes(val, {nearestTo: 3600 * 12})
        case "24h":
          return roundToNearestMinutes(val, {nearestTo: 3600 * 24})
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
        alert('Copied: \n' + txt)
      });
    },
    serviceLink(service) {
      if (service.permalink) {
        service = this.$store.getters.serviceById(service.permalink)
      }
      if (service === undefined || this.isEmptyObject(service)) {
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
    convertToChartData(data = [], multiplier = 1, asInt = false) {
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
    humanTimeNum(val) {
      if (val >= 1000) {
        return Math.round(val / 1000)
      }
      return val
    },
    firstDayOfMonth(date) {
      return startOfMonth(date)
    },
    lastDayOfMonth(month) {
      return lastDayOfMonth(month)
    },
    addMonths(date, amount) {
      return addMonths(date, amount)
    },
    addSeconds(date, amount) {
      return addSeconds(date, amount)
    }
  }
});
