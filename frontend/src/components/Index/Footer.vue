<template>
    <footer>
        <div v-if="!core.footer" class="footer text-center mb-4 p-2">
          <div class="d-block text-dim">
          <div class="mb-3">
            <router-link :to="admin ? '/dashboard' : '/login'">{{$t('top_nav.dashboard')}}</router-link>
          </div>
            <span class="text-muted font-1 mt-3">
            <a href="https://github.com/statping/statping" target="_blank">Statping v{{core.version}}
              made with <font-awesome-icon icon="heart" class="font-1 text-danger"/></a>
            </span>
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
    },
    admin() {
      return this.$store.getters.admin
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
