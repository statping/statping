<template>
  <div class="col-12">
  <div class="text-center" style="width:210px" v-if="!loaded">
    <font-awesome-icon icon="circle-notch" class="h-25 text-dim" spin/>
  </div>
  <apexchart v-else width="100%" height="50" type="bar" :options="chartOpts" :series="data"></apexchart>
  </div>
</template>

<script>
import Api from "@/API";
const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };


export default {
  name: "FailuresBarChart",
  props: {
    service: {
      required: true,
      type: Object,
    },
    group: {
      required: true,
      type: String,
    },
    start: {
      required: true,
      type: String,
    },
    end: {
      required: true,
      type: String,
    },
  },
  data() {
    return {
      data: null,
      loaded: false,
      chartOpts: {
        chart: {
          type: 'bar',
          height: 150,
          sparkline: {
            enabled: true
          },
          animations: {
            enabled: false,
          },
        },
        xaxis: {
          type: 'datetime',
        },
        showPoint: false,
        fullWidth:true,
        chartPadding: {top: 0,right: 0,bottom: 0,left: 80},
        stroke: {
          curve: 'straight'
        },
        fill: {
          opacity: 0.4,
        },
        yaxis: {
          min: 0,
          max: 1,
        },
        plotOptions: {
          bar: {
            colors: {
              ranges: [{
                from: 0,
                to: 1,
                color: '#cfcfcf'
              }, {
                from: 2,
                to: 3,
                color: '#f58e49'
              }, {
                from: 3,
                to: 20,
                color: '#e01a1a'
              }, {
                from: 21,
                to: Infinity,
                color: '#9b0909'
              }]
            },
          },
        },
        tooltip: {
          theme: false,
          enabled: true,
          custom: ({series, seriesIndex, dataPointIndex, w}) => {
            let val = series[seriesIndex][dataPointIndex];
            let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
            const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
            let ago = `${(dataPointIndex-12) * -1} hours ago`
            if ((dataPointIndex-12) * -1 === 0) {
              ago = `Current hour`
            }
            return `<div class="chart_list_tooltip font-2 mb-4">${val-1} Failures<br>${dt}</div>`
          },
          fixed: {
            enabled: true,
            position: 'topLeft',
            offsetX: 0,
            offsetY: 0,
          },
          x: {
            formatter: (value) => { return value },
          },
          y: {
            show: false
          },
        },
        title: {
          enabled: false,
        },
        subtitle: {
          enabled: false,
        }
      }
    }
  },
  async mounted() {
    await this.loadFailures()
  },
  watch: {
    group(o, n) {
      this.loaded = false
      this.loadFailures()
      this.loaded = true
    },
    start(o, n) {
      this.loaded = false
      this.loadFailures()
      this.loaded = true
    },
    end(o, n) {
      this.loaded = false
      this.loadFailures()
      this.loaded = true
    },
  },
  methods: {
    convertChartData(data) {
      if (!data) {
        return []
      }
      let arr = []
      data.forEach((d, k) => {
        arr.push({
          x: d.timeframe,
          y: d.amount+1,
        })
      })
      return arr
    },
    async loadFailures() {
      this.loaded = false
      const startEnd = this.startEndParams(this.parseISO(this.start), this.parseISO(this.end), this.group)
      const data = await Api.service_failures_data(this.service.id, startEnd.start, startEnd.end, this.group, true)
      this.loaded = true
      this.data = [{data: this.convertChartData(data)}]
    }
  },
}
</script>
