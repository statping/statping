<template>
    <div>
        <div class="d-flex mt-3 mb-2">
            <div class="flex-fill service_day" v-for="(d, index) in failureData" :class="{'mini_error': d.amount > 0, 'mini_success': d.amount === 0}">
                <span v-if="d.amount != 0" class="small">{{d.amount}}</span>
            </div>
        </div>
        <div class="row mt-2">
            <div class="col-4 text-left font-2 text-muted">30 Days Ago</div>
            <div class="col-4 text-center font-2" :class="{'text-muted': service.online, 'text-danger': !service.online}">
               {{service_txt}}
            </div>
            <div class="col-4 text-right font-2 text-muted">Today</div>
        </div>
    </div>
</template>

<script>
    import Api from '../../API';

export default {
  name: 'GroupServiceFailures',
  components: {

  },
    data() {
        return {
            failureData: [],
        }
    },
  props: {
      service: {
          type: Object,
          required: true
      }
  },
  computed: {
    service_txt() {
      const s = this.service
      if (!s.online) {
        if (!this.toUnix(this.parseISO(s.last_success))) {
          return `Always Offline`
        }
        return `Offline for ${this.ago(s.last_success)}`
      }
      return `${this.service.online_24_hours}% Uptime`
    }
  },
    mounted () {
      this.lastDaysFailures()
    },
    methods: {
      async lastDaysFailures() {
        const start = this.nowSubtract(86400 * 30)
        const data = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(this.startToday()), "24h")
        data.forEach((d) => {
          let date = this.parseISO(d.timeframe)
          this.failureData.push({month: 1, day: date.getDate(), amount: d.amount})
        })
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
