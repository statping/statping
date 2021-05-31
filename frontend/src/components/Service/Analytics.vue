<template>
    <div class="col-12 col-md-6 mt-2 mt-md-4">
        <div class="col-12 pt-2 sub-service-card">
        <div class="col-8 float-left p-0">
            <span class="font-4 d-block text-muted">{{func.title}}</span>
            <span class="font-2 d-block subtitle">{{func.subtitle}}</span>
        </div>
        <div class="col-4 float-right text-right mt-2 p-0">
            <span class="text-success font-4 font-weight-bold">{{func.value}}</span>
        </div>
        </div>
    </div>
</template>

<script>
    import MiniSparkLine from './MiniSparkLine';
    import ServiceSparkLine from '../Dashboard/ServiceSparkLine';

    export default {
        name: 'Analytics',
        components: { MiniSparkLine, ServiceSparkLine },
        props: {
            func: {
                type: Object,
                required: true
            },
        },
      data() {
        return {
            value: 0,
            title: "",
            subtitle: "",
            chart: [],
        }
      },
      async mounted() {
            if (this.func) {
                this.value = this.func.value;
                this.title = this.func.title;
                this.subtitle = this.func.subtitle;
                this.chart = this.convertToChartData(this.func.chart);
            }
      },
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
