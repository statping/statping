<template v-if="series.length">
    <apexchart width="100%" height="180" type="area" :options="chartOpts" :series="series"></apexchart>
</template>

<script>
export default {
  name: 'ServiceSparkLine',
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
          type: 'area',
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
        colors: ['#DCE6EC'],
        tooltip: {
          enabled: false
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
