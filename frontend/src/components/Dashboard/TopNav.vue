<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <router-link to="/" class="navbar-brand">Statping</router-link>
        <button @click="navopen = !navopen" class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
            <font-awesome-icon v-if="!navopen" icon="bars"/>
            <font-awesome-icon v-if="navopen" icon="times"/>
        </button>

        <div class="navbar-collapse" :class="{collapse: !navopen}" id="navbarText">
            <ul class="navbar-nav mr-auto">
                <li @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard" class="nav-link">{{ $t('top_nav.dashboard') }}</router-link>
                </li>
                <li @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard/services" class="nav-link">{{ $t('top_nav.services') }}</router-link>
                </li>
                <li v-if="$store.state.admin" @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard/users" class="nav-link">{{ $t('top_nav.users') }}</router-link>
                </li>
                <li @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard/messages" class="nav-link">{{ $t('top_nav.announcements') }}</router-link>
                </li>
                <li v-if="$store.state.admin" @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard/settings" class="nav-link">{{ $t('top_nav.settings') }}</router-link>
                </li>
                <li v-if="$store.state.admin" @click="navopen = !navopen" class="nav-item navbar-item">
                    <router-link to="/dashboard/logs" class="nav-link">{{ $t('top_nav.logs') }}</router-link>
                </li>
            </ul>
            <span class="navbar-text">
      <a href="#" class="nav-link" @click.prevent="logout">{{ $t('top_nav.logout') }}</a>
    </span>
        </div>
    </nav>

</template>

<script>
  import Api from "../../API"
  import Vue from "vue";

  export default {
  name: 'TopNav',
      data () {
          return {
              navopen: false
          }
      },
      methods: {
        async logout () {
          await Api.logout()
          this.$store.commit('setHasAllData', false)
          this.$store.commit('setToken', null)
          this.$store.commit('setAdmin', false)
          // this.$cookies.remove("statping_auth")
          await this.$router.push('/logout')
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
