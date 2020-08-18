<template>
    <footer>
        <div v-if="!core.footer" class="footer text-center mb-4 p-2">
          <div class="d-block small text-dim">
          <div class="d-block">
            <router-link :to="$store.state.admin ? '/dashboard' : '/login'">{{$t('top_nav.dashboard')}}</router-link>
          </div>
            <a class="mt-3" href="https://github.com/statping/statping" target="_blank">Statping</a> v{{core.version}}
            made with <font-awesome-icon icon="heart" class="text-danger"/>
          </div>
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
      },
    commit() {
      return this.$store.getters.core.commit.slice(0,8)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
