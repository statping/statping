<template v-if="series.length">
    <apexchart width="100%" height="70" type="bar" :options="chartOpts" :series="series"></apexchart>
</template>

<script>
  const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };


  export default {
  name: 'MiniSparkLine',
  props: {
    series: {
      type: Array,
      default: []
    },
    title: {
      type: String,
    },
    subtitle: {
      type: String,
    }
  },
    watch: {
        title () {

        },
        subtitle () {

        }
    },
  data() {
    return {
      chartOpts: {
        chart: {
          type: 'bar',
          height: 180,
          sparkline: {
            enabled: true
          },
        },
        stroke: {
          curve: 'straight'
        },
        fill: {
          opacity: 0.3,
        },
        yaxis: {
          min: 0
        },
        colors: ['#b3bdc3'],
        tooltip: {
          theme: false,
          enabled: false,
          x: {
            show: false,
          },
          y: {
            formatter: (value) => { return value + " %" },
          },
        },
        title: {
          text: this.title,
          offsetX: 0,
          style: {
            fontSize: '28px',
            cssClass: 'apexcharts-yaxis-title'
          }
        },
        subtitle: {
          text: this.subtitle,
          offsetX: 0,
          style: {
            fontSize: '14px',
            cssClass: 'apexcharts-yaxis-title'
          }
        }
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
