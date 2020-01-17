<template>
    <div>
        <h1 class="col-12 text-center mb-4 mt-sm-3 header-title">{{$store.getters.core.name}}</h1>
        <h5 class="col-12 text-center mb-5 header-desc">{{$store.getters.core.description}}</h5>
    </div>
</template>

<script>
export default {
  name: 'Header',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
