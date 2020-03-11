<template>
    <div>
        <h1 class="col-12 text-center pt-4 mt-4 header-title">{{$store.getters.core.name}}</h1>
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
