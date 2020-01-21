<template>
    <div>{{data}}</div>
</template>

<script>
  import Api from "../../components/API"

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
      this.data = await Api.service_hits(this.props.service.id, 0, 99999999999, "minute")
      this.series = [this.data]
      this.ready = true
    }
  },
  data () {
    return {
      ready: true,
      data: null,
      chartOptions: {
        chart: {
          id: 'vuechart-example',
        },
        xaxis: {
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998],
        },
      },
      series: [{
        name: 'Vue Chart',
        data: [30, 40, 45, 50, 49, 60, 70, 81]
      }]
    }
  },
  mounted() {
    const max = 90;
    const min = 20;
    const newData = this.series[0].data.map(() => {
      return Math.floor(Math.random() * (max - min + 1)) + min
    })
    // In the same way, update the series option
    this.series = [{
      data: newData
    }]
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
