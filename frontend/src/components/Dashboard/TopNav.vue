<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <router-link to="/" class="navbar-brand">Statping</router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarText">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <router-link to="/dashboard" class="nav-link">Dashboard</router-link>
                </li>
                <li class="nav-item">
                    <router-link to="/dashboard/services" class="nav-link">Services</router-link>
                </li>
                <li class="nav-item">
                    <router-link to="/dashboard/users" class="nav-link">Users</router-link>
                </li>
                <li class="nav-item">
                    <router-link to="/dashboard/messages" class="nav-link">Messages</router-link>
                </li>

                <li class="nav-item">
                    <router-link to="/dashboard/settings" class="nav-link">Settings</router-link>
                </li>
                <li class="nav-item">
                    <router-link to="/dashboard/logs" class="nav-link">Logs</router-link>
                </li>

                <li class="nav-item">
                    <router-link to="/dashboard/help" class="nav-link">Help</router-link>
                </li>
            </ul>
            <span class="navbar-text">
      <a href="#" class="nav-link" v-on:click="logout">Logout</a>
    </span>
        </div>
    </nav>

</template>

<script>
  import Api from "../API"

  export default {
  name: 'TopNav',
  props: {

  },
  methods: {
    async logout () {
      await Api.logout()
      this.$store.commit('setHasAllData', false)
      this.$store.commit('setToken', null)
      await this.$router.push('/')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
