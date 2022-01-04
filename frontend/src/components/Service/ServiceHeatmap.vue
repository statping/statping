<template>
    <apexchart v-if="ready" width="100%" height="180" type="heatmap" :options="plotOptions" :series="series"></apexchart>
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
              plotOptions: {
                  chart: {
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
                      colors: [ "#cb3d36" ],
                      enableShades: true,
                      shadeIntensity: 0.5,
                      colorScale: {
                          ranges: [ {
                              from: 0,
                              to: 0,
                              color: '#bababa',
                              name: 'none',
                          },
                              {
                                  from: 2,
                                  to: 10,
                                  color: '#cb3d36',
                                  name: 'medium',
                              },
                              {
                                  from: 11,
                                  to: 999,
                                  color: '#cb221c',
                                  name: 'high',
                              }
                          ]
                      },
                  xaxis: {
                      tickAmount: '1',
                      tickPlacement: 'between',
                      min: 1,
                      max: 31,
                      type: "numeric",
                      labels: {
                          show: true
                      },
                      tooltip: {
                          enabled: false
                      }
                  },
                  yaxis: {
                      labels: {
                          show: true
                      },
                  }
                  },
              series: [ {
                  data: []
              } ]
          }
      },
      methods: {
          async chartHeatmap() {
            const monthData = []
            let start = this.firstDayOfMonth(this.now())

            for (let i=0; i<3; i++) {
                monthData.push(await this.heatmapData(this.addMonths(start, -i), this.lastDayOfMonth(this.addMonths(start, -i))))
            }

            this.series = monthData
            this.ready = true
          },
          async heatmapData(start, end) {
              const data = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(end), "24h", true)
              let dataArr = []
                if (!data) {
                  return {name: start.toLocaleString('en-us', { month: 'long'}), data: []}
                }
              data.forEach((d) => {
                dataArr.push({x: this.parseISO(d.timeframe), y: d.amount});
              });

              let date = new Date(dataArr[0].x);
              return {name: start.toLocaleString('en-us', { month: 'long'}), data: dataArr}
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
