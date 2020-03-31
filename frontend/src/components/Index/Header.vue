<template>
    <div>
        <h1 id="title" class="col-12 text-center pt-4 mt-4 mb-3 header-title font-6">{{$store.getters.core.name}}</h1>
        <h5 id="description" class="col-12 text-center mb-5 header-desc font-3">{{$store.getters.core.description}}</h5>
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
