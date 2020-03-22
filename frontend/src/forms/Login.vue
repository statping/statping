<template>
    <form @submit.prevent="login">
        <div class="form-group row">
            <label for="username" class="col-sm-2 col-form-label">Username</label>
            <div class="col-sm-10">
                <input @keyup="checkForm" type="text" v-model="username" name="username" class="form-control" id="username" placeholder="Username" autocorrect="off" autocapitalize="none">
            </div>
        </div>
        <div class="form-group row">
            <label for="password" class="col-sm-2 col-form-label">Password</label>
            <div class="col-sm-10">
                <input @keyup="checkForm" type="password" v-model="password" name="password" class="form-control" id="password" placeholder="Password">
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <div v-if="error" class="alert alert-danger" role="alert">
                    Incorrect username or password
                </div>
                <button @click.prevent="login" type="submit" class="btn btn-block mb-3 btn-primary" :disabled="disabled || loading">
                    {{loading ? "Loading" : "Sign in"}}
                </button>
            </div>
        </div>

    </form>
</template>

<script>
  import Api from "../API";

  export default {
      name: 'FormLogin',
      data() {
          return {
              username: "",
              password: "",
              auth: {},
              loading: false,
              error: false,
              disabled: true,
              ghLoginURL: ""
          }
      },
      mounted() {
          this.GHlogin()
      },
      methods: {
          checkForm() {
              if (!this.username || !this.password) {
                  this.disabled = true
              } else {
                  this.disabled = false
              }
          },
          async login() {
              this.loading = true
              this.error = false
              const auth = await Api.login(this.username, this.password)
              if (auth.error) {
                  this.error = true
              } else if (auth.token) {
                  this.auth = Api.saveToken(this.username, auth.token, auth.admin)
                  this.$store.dispatch('loadAdmin')
                  this.$store.commit('setAdmin', auth.admin)
                  this.$router.push('/dashboard')
              }
              this.loading = false
          },
          async GHlogin() {
              const core = this.$store.getters.core;
              this.ghLoginURL = `https://github.com/login/oauth/authorize?client_id=${core.gh_client_id}&redirect_uri=${core.domain}/oauth/callback&scope=user,repo`
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
