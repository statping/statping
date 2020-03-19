<template>
    <footer>
        <div v-if="!$store.getters.core.footer" class="footer text-center mb-4 p-2">
            <a href="https://github.com/statping/statping" target="_blank">
                Statping {{$store.getters.core.version}} made with <font-awesome-icon style="color: #d40d0d" icon="heart"/>
            </a> |
            <router-link :to="$store.getters.core.logged_in ? '/dashboard' : '/login'">Dashboard</router-link>
        </div>
        <div v-else class="footer text-center mb-4 p-2" v-html="$store.getters.core.footer"></div>
    </footer>
</template>

<script>
  import Dashboard from "../../pages/Dashboard";

  export default {
  name: 'Footer',
  components: {
    Dashboard
  },
  props: {
    version: String,
    logged_in: Boolean
  },
      watch: {
          logged_in() {

          }
      }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
