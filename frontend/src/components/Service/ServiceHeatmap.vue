<template>
    <apexchart v-if="ready" width="100%" height="300" type="heatmap" :options="chartOptions" :series="series"></apexchart>
</template>

<script>
  import Api from "../../API"

  export default {
      name: 'ServiceHeatmap',
      props: {
          service: {
              type: Object,
              required: true
          }
      },
      async created() {
          await this.chartHeatmap()
      },
      data() {
          return {
              ready: false,
              data: [],
              chartOptions: {
                  heatmap: {
                      colorScale: {
                          ranges: [{
                              from: 0,
                              to: 1,
                              color: 'rgba(235,63,48,0.69)',
                              name: 'low',
                          },
                              {
                                  from: 2,
                                  to: 10,
                                  color: 'rgba(245,43,43,0.58)',
                                  name: 'medium',
                              },
                              {
                                  from: 11,
                                  to: 999,
                                  color: '#cb221c',
                                  name: 'high',
                              }
                          ]
                      }
                  },
                  chart: {
                      height: "100%",
                      width: "100%",
                      type: 'heatmap',
                      toolbar: {
                          show: false
                      }
                  },
                  dataLabels: {
                      enabled: false,
                  },
                  enableShades: true,
                  shadeIntensity: 0.5,
                  colors: ["#d53a3b"],
                  series: [{data: [{}]}],
                  yaxis: {
                      labels: {
                          formatter: (value) => {
                              return value
                          },
                      },
                  },
                  tooltip: {
                      enabled: true,
                      x: {
                          show: false,
                      },
                      y: {
                          formatter: function(val, opts) { return val+" Failures" },
                          title: {
                              formatter: (seriesName) => seriesName,
                          },
                      },
                  }
              },
              series: [{
                  data: []
              }]
          }
      },
      methods: {
          async chartHeatmap() {
              let start = new Date(new Date().getUTCFullYear(), new Date().getUTCMonth()-3, 1);

              let monthData = [];

              for (i=0; i<=3; i++) {
                  let end = this.lastDayOfMonth(start.getUTCMonth()+start)
                  const inputdata = this.heatmapData(start,end)
                  monthData.push(inputdata)
              }

              this.series = monthData
              this.ready = true
          },
          async heatmapData(start, end) {
              console.log(start, end)

              const data = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(end), "24h", true)

              let dataArr = []
              data.forEach(function(d) {
                  dataArr.push({x: d.timeframe, y: d.amount});
              });

              let date = new Date(dataArr[0].x);
              const output = [{name: date.toLocaleString('en-us', { month: 'long'}), data: dataArr}]

          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
