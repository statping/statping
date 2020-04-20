<template>
    <div>
    <form @submit.prevent="login" autocomplete="on">
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

        <a v-if="oauth.gh_client_id" :href="GHlogin()" class="btn btn-block">
            Github Login
        </a>

        <a v-if="oauth.slack_client_id" :href="Slacklogin()" class="btn btn-block">
            Slack Login
        </a>

        <a v-if="oauth.google_client_id" :href="Googlelogin()" class="btn btn-block">
            Google Login
        </a>

    </div>
</template>

<script>
  import Api from "../API";

  export default {
      name: 'FormLogin',
      computed: {
        core() {
          return this.$store.getters.core
        },
        oauth() {
          return this.$store.getters.oauth
        }
      },
      data() {
          return {
              username: "",
              password: "",
              auth: {},
              loading: false,
              error: false,
              disabled: true,
            google_scope: "https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email",
            slack_scope: "identity.email,identity.basic"
          }
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
        GHlogin() {
          return `https://github.com/login/oauth/authorize?client_id=${this.oauth.gh_client_id}&redirect_uri=${this.core.domain}/api/oauth/github&scope=user,repo`
        },
        Slacklogin() {
          return `https://slack.com/oauth/authorize?client_id=${this.oauth.slack_client_id}&redirect_uri=${this.core.domain}/api/oauth/slack&scope=users.profile:read,users:read.email`
        },
        Googlelogin() {
          return `https://accounts.google.com/signin/oauth?client_id=${this.oauth.google_client_id}&redirect_uri=${this.core.domain}/api/oauth/google&response_type=code&scope=https://www.googleapis.com/auth/userinfo.profile+https://www.googleapis.com/auth/userinfo.email`
        }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
