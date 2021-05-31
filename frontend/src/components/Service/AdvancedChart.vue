<template>
    <div class="service-chart-container">
        <apexchart width="100%" height="350" type="area" :options="main_chart_options" :series="main_chart"></apexchart>
    </div>
</template>

<script>
    import Api from "../../API";
    const timeoptions = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };

    export default {
        name: 'AdvancedChart',
        props: {
          service: {
            type: Object,
            required: true
          },
          start: {
            type: String,
            required: true
          },
          end: {
            type: String,
            required: true
          },
          group: {
            type: String,
            required: true
          },
          updated: {
            type: Function,
            required: true
          },
        },
      data() {
        return {
          loading: true,
          main_data: null,
          ping_data: null,
          expanded_data: null,
          main_chart_options: {
            noData: {
              text: "Loading...",
              align: 'center',
              verticalAlign: 'middle',
              offsetX: 0,
              offsetY: -20,
              style: {
                color: "#bababa",
                fontSize: '27px'
              }
            },
            chart: {
              id: 'mainchart',
              stacked: true,
              events: {
                dataPointSelection: (event, chartContext, config) => {
                  window.console.log('slect')
                  window.console.log(event)
                },
                updated: (chartContext, config) => {
                  window.console.log('updated')
                },
                beforeZoom: (chartContext, { xaxis }) => {
                  const start = (xaxis.min / 1000).toFixed(0)
                  const end = (xaxis.max / 1000).toFixed(0)
                  window.console.log(start, end)
                  this.updated(this.fromUnix(start), this.fromUnix(end))
                  return {
                    xaxis: {
                      min: this.fromUnix(start),
                      max: this.fromUnix(end)
                    }
                  }
                },
                scrolled: (chartContext, { xaxis }) => {
                  window.console.log(xaxis)
                },
              },
              height: 500,
              width: "100%",
              type: "area",
              animations: {
                enabled: false,
                initialAnimation: {
                  enabled: true
                }
              },
              selection: {
                enabled: true
              },
              zoom: {
                enabled: true
              },
              toolbar: {
                show: true
              },
              stroke: {
                show: false,
                curve: 'stepline',
                lineCap: 'butt',
              },
            },
            xaxis: {
              type: "datetime",
              labels: {
                show: true
              },
              tooltip: {
                enabled: false
              }
            },
            yaxis: {
              labels: {
                formatter: (value) => {
                  return this.humanTime(value)
                }
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
              custom: function ({ series, seriesIndex, dataPointIndex, w }) {
                let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
                const dt = new Date(ts).toLocaleDateString("en-us", timeoptions)
                let val = series[seriesIndex][dataPointIndex];
                if (val >= 10000) {
                  val = Math.round(val / 1000) + " ms"
                } else {
                  val = val + " μs"
                }
                return `<div class="chartmarker"><span>Response Time: </span><span class="font-3">${val}</span><span>${dt}</span></div>`
              },
              fixed: {
                enabled: true,
                position: 'topRight',
                offsetX: -30,
                offsetY: 40,
              },
              x: {
                show: true,
              },
              y: {
                formatter: undefined,
                title: {
                  formatter: (seriesName) => seriesName,
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
              show: true
            },
            axisBorder: {
              show: false
            },
            fill: {
              colors: ["#f1771f", "#48d338"],
              opacity: 1,
              type: 'solid'
            },
            stroke: {
              show: true,
              curve: 'stepline',
              lineCap: 'butt',
              colors: ["#f1771f", "#48d338"],
              width: 2,
            }
          },
          expanded_chart_options: {
            chart: {
              id: "chart1",
              height: 130,
              type: "bar",
              foreColor: "#ccc",
              brush: {
                target: "chart2",
                enabled: true
              },
              selection: {
                enabled: true,
                fill: {
                  color: "#fff",
                  opacity: 0.4
                },
                xaxis: {
                  min: new Date("27 Jul 2017 10:00:00").getTime(),
                  max: new Date("14 Aug 2999 10:00:00").getTime()
                }
              }
            },
            colors: ["#FF0080"],
            stroke: {
              width: 2
            },
            grid: {
              borderColor: "#444"
            },
            markers: {
              size: 0
            },
            xaxis: {
              type: "datetime",
              tooltip: {
                enabled: false
              }
            },
            yaxis: {
              tickAmount: 2
            }
          }
        }
      },
      async mounted() {
        await this.update_data();
      },
      computed: {
        main_chart () {
          return [{
            name: "latency",
            ...this.convertToChartData(this.main_data)
          },{
            name: "ping",
            ...this.convertToChartData(this.ping_data)
          }]
        },
        expanded_chart () {
          return this.toBarData(this.expanded_data)
        },
        params () {
          return {start: this.toUnix(new Date(this.start)), end: this.toUnix(new Date(this.end))}
        },
      },
      watch: {
        start: function(n, o) {
          this.update_data()
        },
        end: function(n, o) {
          this.update_data()
        },
        group: function(n, o) {
          this.update_data()
        },
      },
      methods: {
          async update_data() {
            this.loading = true
            await this.chartHits()
            // await this.expanded_hits()
            this.loading = false
          },
        async expanded_hits() {
          this.expanded_data = await this.load_hits(0, 99999999999, "24h")
        },
        async chartHits() {
          this.main_data = await this.load_hits()
          this.ping_data = await this.load_ping()
        },
        async load_hits(start=this.params.start, end=this.params.end, group=this.group) {
          return await Api.service_hits(this.service.id, start, end, group, false)
        },
        async load_ping(start=this.params.start, end=this.params.end, group=this.group) {
          return await Api.service_ping(this.service.id, start, end, group, false)
        }
      }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
