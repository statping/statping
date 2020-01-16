<template>
    <div class="col-12 card mb-3" style="min-height: 260px">
        <div class="card-body">
            <h5 class="card-title"><a href="service/7">{{service.name}}</a>
                <span class="badge float-right badge-success">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
            </h5>
            <div class="row">
                <div class="col-md-3 col-sm-6">
                    <div id="spark_service_7_1"></div>
                </div>
                <div class="col-md-3 col-sm-6">
                    <div id="spark_service_7_2"></div>
                </div>
                <div class="col-md-3 col-sm-6">
                    <div id="spark_service_7_3"></div>
                </div>
                <div class="col-md-3 col-sm-6">
                    <div id="spark_service_7_4"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'ServiceInfo',
  props: {
    service: Object
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
