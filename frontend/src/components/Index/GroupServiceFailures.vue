<template>
    <div>
        <div class="d-flex mt-3 mb-2">
            <div class="flex-fill service_day" v-for="(d, index) in failureData" :class="{'mini_error': d.amount > 0, 'mini_success': d.amount === 0}"></div>
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
            failureData: null,
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
      if (!this.service.online) {
        if (!this.toUnix(this.service.last_success)) {
          return `Always Offline`
        }
        return `Offline for ${this.ago(this.service.last_success)}`
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
        this.failureData = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(this.now()), "24h")
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .service_day {
        height: 20px;
        margin-right: 2px;
        border-radius: 4px;
    }

    @keyframes pulse_animation {
        0% { transform: scale(1); }
        30% { transform: scale(1); }
        40% { transform: scale(1.02); }
        50% { transform: scale(1); }
        60% { transform: scale(1); }
        70% { transform: scale(1.05); }
        80% { transform: scale(1); }
        100% { transform: scale(1); }
    }

    .pulse {
        animation-name: pulse_animation;
        animation-duration: 1500ms;
        transform-origin:70% 70%;
        animation-iteration-count: infinite;
        animation-timing-function: linear;
    }


    @keyframes glow-grow {
        0% {
            opacity: 0;
            transform: scale(1);
        }
        80% {
            opacity: 1;
        }
        100% {
            transform: scale(2);
            opacity: 0;
        }
    }
    .pulse-glow {
        animation-name: glow-grown;
        animation-duration: 100ms;
        transform-origin: 70% 30%;
        animation-iteration-count: infinite;
        animation-timing-function: linear;
    }

    .pulse-glow:before,
    .pulse-glow:after {
        position: absolute;
        content: "";
        height: 0.4rem;
        width: 1.7rem;
        top: 1.3rem;
        right: 2.15rem;
        border-radius: 0;
        box-shadow: 0 0 6px #47d337;
        animation: glow-grow 2s ease-out infinite;
    }
</style>
