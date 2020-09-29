<template>
  <div v-if="loading" class="row mt-3 mb-3">
    <div class="col-12 text-center text-muted">
      <font-awesome-icon icon="circle-notch" size="3x" spin/>
    </div>
  </div>
</template>

<script>
export default {
name: "Loading",
  props: {
    loading: {
      type: Boolean,
      required: true,
    }
  }
}
</script>

<style scoped>

</style>
