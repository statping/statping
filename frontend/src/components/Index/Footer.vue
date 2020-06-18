<template>
    <footer>
        <div v-if="!core.footer" class="footer text-center mb-4 p-2">
            <a href="https://github.com/statping/statping" target="_blank">
                Statping {{core.version}} made with <font-awesome-icon icon="heart" class="text-danger"/>
            </a> |
            <router-link :to="$store.state.admin ? '/dashboard' : '/login'">{{$t('top_nav.dashboard')}}</router-link>
        </div>
        <div v-else class="footer text-center mb-4 p-2" v-html="core.footer"></div>
    </footer>
</template>

<script>
  import Dashboard from "../../pages/Dashboard";

  export default {
  name: 'Footer',
  components: {
    Dashboard
  },
  computed: {
      core() {
        return this.$store.getters.core
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
