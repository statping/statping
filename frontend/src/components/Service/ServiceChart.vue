<template v-show="showing">
    <apexchart v-if="ready" width="100%" height="235" type="area" :options="chartOptions" :series="series"/>
</template>

<script>
  import Api from "../../API"

  const axisOptions = {
    labels: {
      show: false
    },
    crosshairs: {
      show: false
    },
    lines: {
      show: false
    },
    tooltip: {
      enabled: false
    },
    axisTicks: {
      show: false
    },
    grid: {
      show: false
    },
    marker: {
      show: false
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
                  xaxis: {
                      type: "datetime",
                      ...axisOptions
                  },
                  yaxis: {
                      ...axisOptions
                  },
                  tooltip: {
                      enabled: false,
                      marker: {
                          show: false,
                      },
                      x: {
                          show: false,
                      }
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
                  this.chartHits("hour")
              }
          }
      },
      methods: {
          async chartHits(group) {
              const start = this.nowSubtract((3600 * 24) * 30)
              this.data = await Api.service_hits(this.service.id, this.toUnix(start), this.toUnix(new Date()), group)

              if (this.data.length === 0 && group !== "hour") {
                  await this.chartHits("hour")
              }
              this.series = [{
                  name: this.service.name,
                  ...this.convertToChartData(this.data)
              }]
              this.ready = true
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
