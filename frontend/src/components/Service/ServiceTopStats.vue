<template>
    <div class="row stats_area mt-5 mb-4">
        <div class="col-4">
            <span class="lg_number">{{service.avg_response}}ms</span>
            Average Response
        </div>
        <div class="col-4">
            <span class="lg_number">{{service.online_24_hours}}%</span>
            Uptime last 24 Hours
        </div>
        <div class="col-4">
            <span class="lg_number">{{service.online_7_days}}%</span>
            Uptime last 7 Days
        </div>
    </div>
</template>

<script>
    export default {
        name: 'ServiceTopStats',
        props: {
            service: {
                type: Object,
                required: true
            },
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
