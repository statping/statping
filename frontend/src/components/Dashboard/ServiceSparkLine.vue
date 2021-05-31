<template v-if="series.length">
    <apexchart width="100%" height="100" type="bar" :options="chartOpts" :series="series"></apexchart>
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
        showPoint: false,
        fullWidth:true,
        chartPadding: {top: 0,right: 0,bottom: 0,left: 0},
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
          custom: ({series, seriesIndex, dataPointIndex, w}) => {
            let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
            const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
            let val = series[seriesIndex][dataPointIndex];
            return `<div class="chartmarker"><span class="font-3">Average Response Time:  ${this.humanTime(val)}</span><span>${dt}</span></div>`
          },
          fixed: {
            enabled: true,
            position: 'bottomLeft',
            offsetX: 0,
            offsetY: -30,
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
            fontSize: '18px',
            cssClass: 'apexcharts-yaxis-title'
          }
        },
        subtitle: {
          text: this.subtitle,
          offsetX: 0,
          offsetY: 20,
          style: {
            fontSize: '9px',
            cssClass: 'apexcharts-yaxis-title'
          }
        }
      }
    }
  }
}
</script>
