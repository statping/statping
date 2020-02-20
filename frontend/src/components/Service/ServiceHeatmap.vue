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
              const start = this.nowSubtract((3600 * 24) * 7)
              const data = await Api.service_heatmap(this.service.id, this.toUnix(start), this.toUnix(new Date()), "hour")

              let dataArr = []
              data.forEach(function(d) {
                  let date = new Date(d.date);
                  dataArr.push({name: date.toLocaleString('en-us', { month: 'long' }), data: d.data});
              });

              this.series = dataArr
              this.ready = true
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
