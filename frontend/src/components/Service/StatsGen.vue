<template>
    <div class="col-3 text-center">
        <span class="text-success font-5 font-weight-bold">{{value}}</span>
        <span class="font-2 d-block">{{title}}</span>
    </div>
</template>

<script>
    import Api from "../../API";

    export default {
        name: 'StatsGen',
        props: {
            service: {
                type: Object,
                required: true
            },
          title: {
            type: String,
            required: true
          },
          start: {
            type: Number,
            required: true
          },
          end: {
            type: Number,
            required: true
          },
          group: {
            type: String,
            required: true
          },
          expression: {
            type: String,
            required: true
          }
        },
      data() {
        return {
          value: "+17%"
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
