<template>
    <apexchart v-if="ready" width="100%" height="215" type="area" :options="chartOptions" :series="series"></apexchart>
</template>

<script>
  import Api from "../../components/API"

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
    }
  },
  async created() {
    await this.chartHits()
  },
  methods: {
    async chartHits() {
      const start = this.ago(3600 * 24)
      this.data = await Api.service_hits(this.service.id, start, this.now(), "hour")
      this.series = [{
        name: this.service.name,
        ...this.data
      }]
      this.ready = true
    }
  },
  data () {
    return {
      ready: false,
      data: null,
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
            colors: ["#48d338"],
            opacity: 1,
            type: 'solid'
          },
          stroke: {
            show: true,
            curve: 'smooth',
            lineCap: 'butt',
            colors: ["#3aa82d"],
          }
        },
      series: [{
        data: []
      }]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
