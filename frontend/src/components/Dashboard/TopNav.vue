<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <router-link to="/" class="navbar-brand">Statping</router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarText">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <a v-on:click="changeView('DashboardIndex', '/dashboard')" class="nav-link" href="#">Dashboard</a>
                </li>
                <li class="nav-item">
                    <a v-on:click="changeView('DashboardServices', '/dashboard/services')" class="nav-link" href="#">Services</a>
                </li>
                <li class="nav-item">
                    <a v-on:click="changeView('DashboardUsers', '/dashboard/users')" class="nav-link" href="#">Users</a>
                </li>
                <li class="nav-item">
                    <a v-on:click="changeView('DashboardMessages', '/dashboard/messages')" class="nav-link" href="#">Messages</a>
                </li>

                <li class="nav-item">
                    <a v-on:click="changeView('Settings', '/dashboard/settings')" class="nav-link" href="#">Settings</a>
                </li>
                <li class="nav-item">
                    <a v-on:click="changeView('Logs', '/dashboard/logs')" class="nav-link" href="#">Logs</a>
                </li>

                <li class="nav-item">
                    <a v-on:click="changeView('DashboardIndex', '/dashboard/help')" class="nav-link" href="#">Help</a>
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
    changeView: Function
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
