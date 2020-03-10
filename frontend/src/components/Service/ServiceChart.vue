<template v-show="showing">
    <apexchart v-if="ready" width="100%" height="235" type="area" :options="chartOptions" :series="series"/>
</template>

<script>
  import Api from "../../API"
  const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };

  const axisOptions = {
    labels: {
      show: false
    },
    crosshairs: {
      show: true
    },
    lines: {
      show: false
    },
    tooltip: {
      enabled: true
    },
    axisTicks: {
      show: false
    },
    grid: {
      show: false
    },
    marker: {
      show: true
    }
  };

  export default {
      name: 'ServiceChart',
      props: {
          service: {
              type: Object,
              required: true
          },
          visible: {
              type: Boolean,
              required: true
          }
      },
      data() {
          return {
              ready: false,
              showing: false,
              data: [],
              chartOptions: {
                  noData: {
                      text: 'Loading...'
                  },
                  chart: {
                      height: 210,
                      width: "100%",
                      type: "area",
                      animations: {
                          enabled: true,
                          initialAnimation: {
                              enabled: true
                          }
                      },
                      selection: {
                          enabled: false
                      },
                      zoom: {
                          enabled: false
                      },
                      toolbar: {
                          show: false
                      },
                  },
                  grid: {
                      show: false,
                      padding: {
                          top: 0,
                          right: 0,
                          bottom: 0,
                          left: -10,
                      }
                  },
                  dropShadow: {
                      enabled: false,
                  },
                  xaxis: {
                      type: "datetime",
                      labels: {
                          show: false
                      },
                  },
                  yaxis: {
                      labels: {
                          show: false
                      },
                  },
                  tooltip: {
                      theme: false,
                      enabled: true,
                      markers: {
                          size: 0
                      },
                      custom: function({series, seriesIndex, dataPointIndex, w}) {
                          let service = w.globals.seriesNames[0];
                          let ts = w.globals;
                          window.console.log(ts);
                          let val = series[seriesIndex][dataPointIndex];
                          if (val > 1000) {
                              val = (val * 0.1).toFixed(0) + " milliseconds"
                          } else {
                              val = (val * 0.01).toFixed(0) + " microseconds"
                          }
                          return `<div class="chartmarker"><span>${service} Average Response</span> <span class="font-3">${val}</span></div>`
                      },
                      fixed: {
                          enabled: true,
                          position: 'topRight',
                          offsetX: -30,
                          offsetY: 0,
                      },
                      x: {
                          show: false,
                      },
                  },
                  legend: {
                      show: false,
                  },
                  dataLabels: {
                      enabled: false
                  },
                  floating: true,
                  axisTicks: {
                      show: false
                  },
                  axisBorder: {
                      show: false
                  },
                  fill: {
                      colors: [this.service.online ? "#48d338" : "#dd3545"],
                      opacity: 1,
                      type: 'solid'
                  },
                  stroke: {
                      show: false,
                      curve: 'smooth',
                      lineCap: 'butt',
                      colors: [this.service.online ? "#3aa82d" : "#dd3545"],
                  }
              },
              series: [{
                  data: []
              }]
          }
      },
      watch: {
          visible: function(newVal, oldVal) {
              if (newVal && !this.showing) {
                  this.showing = true
                  this.chartHits("2h")
              }
          }
      },
      methods: {
          async chartHits(group) {
              window.console.log(this.service.created_at)
              this.data = await Api.service_hits(this.service.id, this.toUnix(this.service.created_at), this.toUnix(new Date()), group, false)

              if (this.data.length === 0 && group !== "1h") {
                  await this.chartHits("1h")
              }
              this.series = [{
                  name: this.service.name,
                  ...this.convertToChartData(this.data, 0.01)
              }]
              this.ready = true
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
