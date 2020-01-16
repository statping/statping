<template>
    <footer>
        <div class="footer text-center mb-4 p-2">
            <a href="https://github.com/hunterlong/statping" target="_blank">Statping {{version}} made with <i class="text-danger fas fa-heart"></i></a> |
            <router-link to="/dashboard">Dashboard</router-link>
        </div>
    </footer>
</template>

<script>
export default {
  name: 'Footer',
  props: {
    version: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
