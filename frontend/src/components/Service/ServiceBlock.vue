<template>
    <div class="mb-4" id="service_id_1">
        <div class="card">
            <div class="card-body">
                <div class="col-12">
                    <h4 class="mt-3">
                        <router-link :to="`/service/${service.id}`">{{service.name}}</router-link>
                        <span class="badge bg-success float-right">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    </h4>

                    <div class="row stats_area mt-5">
                        <div class="col-4">
                            <span class="lg_number">131ms</span>
                            Average Response
                        </div>
                        <div class="col-4">
                            <span class="lg_number">100%</span>
                            Uptime last 24 Hours
                        </div>
                        <div class="col-4">
                            <span class="lg_number">100%</span>
                            Uptime last 7 Days
                        </div>
                    </div>

                </div>
            </div>

            <div class="chart-container">
                <div id="service_1"></div>
            </div>

            <div class="row lower_canvas full-col-12 text-white">
                <div class="col-10 text-truncate">
                    <span class="d-none d-md-inline">Online, last Failure was Wednesday 1:16:49PM, Dec 18 2019</span>
                </div>
                <div class="col-sm-12 col-md-2">
                    <router-link :to="`/service/${service.id}`" class="btn btn-success btn-sm float-right dyn-dark btn-block">View Service</router-link>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
  name: 'ServiceBlock',
  props: {
    service: Object
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
