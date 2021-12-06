<template>
    <div>
      <div v-observe-visibility="{callback: visibleChart, once: true}" v-if="!loaded" class="row">
        <div class="col-12 text-center mt-3">
          <font-awesome-icon icon="circle-notch" class="text-dim" size="2x" spin/>
        </div>
      </div>
      <transition name="fade">
        <div v-if="loaded">
        <div class="d-flex mt-3">
            <div class="flex-fill service_day" v-for="(d, index) in failureData" @mouseover="mouseover(d)" @mouseout="mouseout" :class="{'day-error': d.amount > 0, 'day-success': d.amount === 0}">
                <span v-if="d.amount !== 0" class="d-none d-md-block text-center small"></span>
            </div>
        </div>
        <div class="row mt-2">
          <div class="col-12 no-select">
            <p class="divided">
              <span class="font-2 text-muted">90 {{$t('days_ago')}}</span>
              <span class="divider"></span>
              <span class="text-center font-2" :class="{'text-muted': service.online, 'text-danger': !service.online}">{{service_txt}}</span>
              <span class="divider"></span>
              <span class="font-2 text-muted">{{$t('today')}}</span>
            </p>
          </div>
        </div>
      <div class="daily-failures small text-right text-dim">{{hover_text}}</div>
      </div>
  </transition>
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
          hover_text: "",
          loaded: false,
          visible: false,
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
      return this.smallText(this.service)
    }
  },
  mounted () {

    },
    methods: {
      visibleChart(isVisible, entry) {
        if (isVisible && !this.visible) {
          this.visible = true
          this.lastDaysFailures().then(() =>  this.loaded = true)
        }
      },
      mouseout() {
        this.hover_text = ""
      },
    mouseover(e) {
      let txt = `${e.amount} Failures`
      if (e.amount === 0) {
        txt = `No Issues`
      }
      this.hover_text = `${e.date.toLocaleDateString()} - ${txt}`
    },
      async lastDaysFailures() {
        const start = this.beginningOf('day', this.nowSubtract(86400 * 90))
        const end = this.endOf('tomorrow')
        const data = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(end), "24h", true)
        data.forEach((d) => {
          let date = this.parseISO(d.timeframe)
          this.failureData.push({month: date.getMonth(), day: date.getDate(), date: date, amount: d.amount})
        })
      }
    }
}
</script>
