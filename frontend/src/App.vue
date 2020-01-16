<template>
  <div id="app">
    <router-view/>
      <Footer version="DEV" />
  </div>
</template>

<script>
  import Footer from "./components/Footer";

  export default {
  name: 'app',
  components: {
    Footer
  },
}
</script>

<style>
</style>
