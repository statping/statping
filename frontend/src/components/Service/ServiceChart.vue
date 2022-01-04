<template v-show="showing">
    <apexchart v-if="ready" class="service-chart" width="100%" height="100%" type="area" :options="chartOptions" :series="series"/>
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
          },
        chart_timeframe: {
          type: Object,
          required: true
        },
      },
      data() {
          return {
              ready: false,
              showing: false,
              data: null,
              ping_data: null,
              series: null,
          }
      },
    computed: {
      chartOptions() {
        return {
          noData: {
            text: 'Loading...'
          },
          chart: {
            height: "100%",
            width: "100%",
            type: "area",
            animations: {
              enabled: true,
              easing: 'easeinout',
              speed: 800,
              animateGradually: {
                enabled: false,
                delay: 400,
              },
              dynamicAnimation: {
                enabled: true,
                speed: 500
              },
              hover: {
                animationDuration: 0, // duration of animations when hovering an item
              },
              responsiveAnimationDuration: 0,
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
            tooltip: {
              enabled: false
            }
          },
          yaxis: {
            labels: {
              show: false
            },
          },
          markers: {
            size: 0,
            strokeWidth: 0,
            hover: {
              size: undefined,
              sizeOffset: 0
            }
          },
          tooltip: {
            theme: false,
            enabled: true,
            custom: ({series, seriesIndex, dataPointIndex, w}) => {
              let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
              const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
              let val = series[0][dataPointIndex];
              let pingVal = series[1][dataPointIndex];
              return `<div class="chartmarker">
<span>Average Response Time: ${this.humanTime(val)}/${this.chart_timeframe.interval}</span>
<span>Average Ping: ${this.humanTime(pingVal)}/${this.chart_timeframe.interval}</span>
<span>${dt}</span>
</div>`
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
            y: {
              formatter: (value) => {
                return value + " %"
              },
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
            colors: this.service.online ? ["#3dc82f", "#48d338"] : ["#c60f20", "#dd3545"],
            opacity: 1,
            type: 'solid',
          },
          stroke: {
            show: false,
            curve: 'smooth',
            lineCap: 'butt',
            colors: this.service.online ? ["#38bc2a", "#48d338"] : ["#c60f20", "#dd3545"],
          }
        }
      }
    },
      watch: {
        visible: function(newVal, oldVal) {
          if (newVal && !this.showing) {
            this.showing = true
            this.chartHits(this.chart_timeframe)
          }
        },
        chart_timeframe: function(newVal, oldVal) {
          if (newVal) {
            this.chartHits(newVal)
          }
        }
      },
      methods: {
          async chartHits(val) {
              this.ready = false
            const end = this.endOf("hour", this.now())
            const start = this.beginningOf("hour", this.fromUnix(val.start_time))
              this.data = await Api.service_hits(this.service.id, this.toUnix(start), this.toUnix(end), val.interval, false)
              this.ping_data = await Api.service_ping(this.service.id, this.toUnix(start), this.toUnix(end), val.interval, false)

            this.series = [
                {name: "Latency", ...this.convertToChartData(this.data)},
                {name: "Ping", ...this.convertToChartData(this.ping_data)},
                ]
            this.ready = true
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
