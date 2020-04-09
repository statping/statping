<template>
    <div class="col-12 bg-white p-4">
        <p v-if="logs.length === 0" class="text-monospace sm">
            Loading Logs...
        </p>
        <p v-for="(log, index) in logs" class="text-monospace sm">{{log}}</p>
    </div>
</template>

<script>
import Api from "../API";

export default {
    name: 'Logs',
    data() {
        return {
            logs: [],
            last: "",
            t: null
        }
    },
    async created() {
        await this.getLogs()
        if (!this.t) {
            this.t = setInterval(async () => {
                await this.lastLog()
            }, 650)
        }
    },
    beforeDestroy() {
        clearInterval(this.t)
    },
    methods: {
        cleanLog(l) {
            const splitLog = l.split(": ")
            const last = splitLog.slice(1);
            return last.join(": ")
        },
        async getLogs() {
            const logs = await Api.logs()
            const l = logs.reverse()
            this.last = this.cleanLog(l[l.length - 1])
            this.logs = l
        },
        async lastLog() {
            const log = await Api.logs_last()
            const cleanLast = this.cleanLog(log)
            if (this.last !== cleanLast) {
                this.last = cleanLast
                this.logs.reverse().push(log)
            }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .sm {
        font-size: 8pt;
    }
</style>
