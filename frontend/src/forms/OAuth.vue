<template>
    <form @submit.prevent="saveOAuth">
        <div class="card text-black-50 bg-white mb-3">
            <div class="card-header">Internal Login</div>
            <div class="card-body">
                <div class="form-group row">
                    <label for="switch-gh-oauth" class="col-sm-4 col-form-label">OAuth Login Settings</label>
                    <div class="col-md-8 col-xs-12 mt-1">
                        <span @click="oauth.internal_enabled = !!core.oauth.internal_enabled" class="switch float-left">
                            <input v-model="oauth.internal_enabled" type="checkbox" class="switch" id="switch-local-oauth" :checked="oauth.internal_enabled">
                            <label for="switch-local-oauth">Use email/password Authentication</label>
                        </span>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="whitelist_domains" class="col-sm-4 col-form-label">Whitelist Domains</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.oauth.oauth_domains" type="text" class="form-control" placeholder="domain.com" id="whitelist_domains">
                    </div>
                </div>
            </div>
        </div>
        <div class="card text-black-50 bg-white mb-3">
            <div class="card-header">Github Settings</div>
            <div class="card-body">
                <span>You will need to create a new <a href="https://github.com/settings/developers">OAuth App</a> within Github.</span>

                <div class="form-group row mt-3">
                    <label for="github_client" class="col-sm-4 col-form-label">Github Client ID</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.oauth.gh_client_id" type="text" class="form-control" id="github_client" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="github_secret" class="col-sm-4 col-form-label">Github Client Secret</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.oauth.gh_client_secret" type="text" class="form-control" id="github_secret" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="switch-gh-oauth" class="col-sm-4 col-form-label">Enable Github Login</label>
                    <div class="col-md-8 col-xs-12 mt-1">
                    <span @click="oauth.github_enabled = !!oauth.github_enabled" class="switch float-left">
                        <input v-model="oauth.github_enabled" type="checkbox" class="switch" id="switch-gh-oauth" :checked="oauth.github_enabled">
                        <label for="switch-gh-oauth"> </label>
                    </span>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="gh_callback" class="col-sm-4 col-form-label">Callback URL</label>
                    <div class="col-sm-8">
                        <div class="input-group">
                            <input v-bind:value="`${core.domain}/oauth/github`" type="text" class="form-control" id="gh_callback" readonly>
                            <div class="input-group-append copy-btn">
                                <button @click.prevent="copy(`${core.domain}/oauth/github`)" class="btn btn-outline-secondary" type="button">Copy</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="card text-black-50 bg-white mb-3">
            <div class="card-header">Google Settings</div>
            <div class="card-body">
                <span>Go to <a href="https://console.cloud.google.com/apis/credentials">OAuth Consent Screen</a> on Google Console to create a new OAuth application.</span>

                <div class="form-group row mt-3">
                    <label for="github_client" class="col-sm-4 col-form-label">Google Client ID</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.google_client_id" type="text" class="form-control" id="google_client" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="github_secret" class="col-sm-4 col-form-label">Google Client Secret</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.google_client_secret" type="text" class="form-control" id="google_secret" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="switch-google-oauth" class="col-sm-4 col-form-label">Enable Google Login</label>
                    <div class="col-md-8 col-xs-12 mt-1">
                    <span @click="google_enabled = !!google_enabled" class="switch float-left">
                        <input v-model="google_enabled" type="checkbox" class="switch" id="switch-google-oauth" :checked="google_enabled">
                        <label for="switch-google-oauth"> </label>
                    </span>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="google_callback" class="col-sm-4 col-form-label">Callback URL</label>
                    <div class="col-sm-8">
                        <div class="input-group">
                            <input v-bind:value="`${core.domain}/oauth/google`" type="text" class="form-control" id="google_callback" readonly>
                            <div class="input-group-append copy-btn">
                                <button @click.prevent="copy(`${core.domain}/oauth/google`)" class="btn btn-outline-secondary" type="button">Copy</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="card text-black-50 bg-white mb-3">
            <div class="card-header">Slack Settings</div>
            <div class="card-body">
                <span>Go to <a href="https://console.cloud.google.com/apis/credentials">OAuth Consent Screen</a> on Google Console to create a new OAuth application.</span>

                <div class="form-group row mt-3">
                    <label for="slack_client" class="col-sm-4 col-form-label">Slack Client ID</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.slack_client_id" type="text" class="form-control" id="slack_client" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="slack_secret" class="col-sm-4 col-form-label">Slack Client Secret</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.slack_client_secret" type="text" class="form-control" id="slack_secret" required>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="slack_secret" class="col-sm-4 col-form-label">Slack Team ID</label>
                    <div class="col-sm-8">
                        <input v-model="oauth.slack_team" type="text" class="form-control" id="slack_team">
                        <small>Optional</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="switch-slack-oauth" class="col-sm-4 col-form-label">Enable Slack Login</label>
                    <div class="col-md-8 col-xs-12 mt-1">
                    <span @click="slack_enabled = !!slack_enabled" class="switch float-left">
                        <input v-model="slack_enabled" type="checkbox" class="switch" id="switch-slack-oauth" :checked="google_enabled">
                        <label for="switch-slack-oauth"> </label>
                    </span>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="slack_callback" class="col-sm-4 col-form-label">Callback URL</label>
                    <div class="col-sm-8">
                        <div class="input-group">
                            <input v-bind:value="`${core.domain}/oauth/slack`" type="text" class="form-control" id="slack_callback" readonly>
                            <div class="input-group-append copy-btn">
                                <button @click.prevent="copy(`${core.domain}/oauth/slack`)" class="btn btn-outline-secondary" type="button">Copy</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <button class="btn btn-primary btn-block" @click.prevent="saveOAuth" type="submit">
            Save OAuth Settings
        </button>

    </form>
</template>

<script>
  import Api from "../API";

  export default {
      name: 'OAuth',
      computed: {
          oauth() {
             return this.$store.getters.core.oauth
          }
      },
      data() {
          return {
            internal_enabled: this.has('local'),
            google_enabled: this.has('google'),
            github_enabled: this.has('github'),
            slack_enabled: this.has('slack')
          }
      },
    mounted() {
      window.console.log(this.core.oauth)
    },
    beforeCreate() {
      // this.github_enabled = this.$store.getters.core.oauth.oauth_providers.split(",").includes('github')
      // const c = await Api.core()
      // this.auth = c.auth
    },
    methods: {
        has(val) {
          if (!this.core.oauth.oauth_providers) {
            return false
          }
          return this.core.oauth.oauth_providers.split(",").includes(val)
        },
        providers() {
          let providers = [];
          if (this.github_enabled) {
            providers.push("github")
          }
          if (this.internal_enabled) {
            providers.push("local")
          }
          if (this.google_enabled) {
            providers.push("google")
          }
          if (this.slack_enabled) {
            providers.push("slack")
          }
          return providers.join(",")
        },
          async saveOAuth() {
            let c = this.$store.getters.core
            c.oauth = this.oauth
            c.oauth.oauth_providers = this.providers()
            await Api.core_save(c)
            const core = await Api.core()
            this.$store.commit('setCore', core)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
