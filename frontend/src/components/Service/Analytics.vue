<template>
    <div class="col-6 mt-4">
        <div class="col-12 sub-service-card">
        <div class="col-8 float-left p-0 mt-1 mb-3">
            <span class="font-5 d-block">{{title}}</span>
            <span class="text-muted font-3 d-block font-weight-bold">{{subtitle}}</span>
        </div>
        <div class="col-4 float-right text-right mt-2 p-0">
            <span class="text-success font-5 font-weight-bold">{{value}}</span>
        </div>

        <MiniSparkLine :series="[{name: 'okokokok', data:[{x: '2019-01-01', y: 120},{x: '2019-01-02', y: 160},{x: '2019-01-03', y: 240},{x: '2019-01-04', y: 45}]}]"/>
        </div>
    </div>
</template>

<script>
    import Api from "../../API";
    import MiniSparkLine from './MiniSparkLine';
    import ServiceSparkLine from './ServiceSparkLine';

    export default {
        name: 'Analytics',
        components: { MiniSparkLine, ServiceSparkLine },
        props: {
            title: {
                type: String,
                required: true
            },
            subtitle: {
                type: String,
                required: true
            },
            value: {
                type: Number,
                required: true
            },
            level: {
                type: Number,
                required: false
            }
        },
      data() {
        return {

        }
      },
      async mounted() {
        await this.latencyYesterday();
      },
      async latencyYesterday() {
        const todayTime = await Api.service_hits(this.service.id, this.toUnix(this.nowSubtract(86400)), this.toUnix(new Date()), this.group, false)
        const fetched = await Api.service_hits(this.service.id, this.start, this.end, this.group, false)

        let todayAmount = this.addAmounts(todayTime)
        let yesterday = this.addAmounts(fetched)

        window.console.log(todayAmount)
        window.console.log(yesterday)

      },
      addAmounts(data) {
        let total = 0
        data.forEach((f) => {
          total += parseInt(f.amount)
        });
        return total
      }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
