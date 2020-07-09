<template v-if="series.length">
    <apexchart width="100%" height="180" type="bar" :options="chartOpts" :series="series"></apexchart>
</template>

<script>
  const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };


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
          enabled: true,
          custom: function({series, seriesIndex, dataPointIndex, w}) {
            let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
            const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
            let val = series[seriesIndex][dataPointIndex];
            val = val + " ms"
            return `<div class="chartmarker"><span>Average Response Time: </span><span class="font-3">${val}</span><span>${dt}</span></div>`
          },
          fixed: {
            enabled: true,
            position: 'topRight',
            offsetX: 0,
            offsetY: 0,
          },
          x: {
            show: true,
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
