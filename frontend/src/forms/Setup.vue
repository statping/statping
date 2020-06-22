<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">
        <div class="col-12 col-md-8 offset-md-2 mb-4">
            <img alt="Statping Setup" class="col-12 mt-5 mt-md-0" style="max-width:680px" src="banner.png">
        </div>

        <div class="col-12">
            <form @submit.prevent="saveSetup">
                <div class="row">
                    <div class="col-6">
                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.language') }}</label>
                            <select @change="changeLanguages" v-model="setup.language" id="language" class="form-control">
                                <option value="en">English</option>
                                <option value="es">Spanish</option>
                                <option value="fr">French</option>
                                <option value="ru">Russian</option>
                                <option value="de">German</option>
                            </select>
                        </div>
                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.connection') }}</label>
                            <select @change="canSubmit" v-model="setup.db_connection" id="db_connection" class="form-control">
                                <option value="sqlite">SQLite</option>
                                <option value="postgres">Postgres</option>
                                <option value="mysql">MySQL</option>
                            </select>
                        </div>
                        <div class="row">
                            <div class="col-6">
                                <div v-if="setup.db_connection !== 'sqlite'" class="form-group">
                                    <label class="text-capitalize">{{ $t('setup.host') }}</label>
                                    <input @keyup="canSubmit" v-model="setup.db_host" id="db_host" type="text" class="form-control" placeholder="localhost">
                                </div>
                            </div>
                            <div class="col-6">
                                <div v-if="setup.db_connection !== 'sqlite'" class="form-group">
                                    <label class="text-capitalize">{{ $t('port') }}</label>
                                    <input @keyup="canSubmit" v-model="setup.db_port" id="db_port" type="number" class="form-control" placeholder="5432">
                                </div>
                            </div>
                        </div>
                        <div v-if="setup.db_connection !== 'sqlite'" class="form-group">
                            <label class="text-capitalize">{{ $t('username') }}</label>
                            <input @keyup="canSubmit" v-model="setup.db_user" id="db_user" type="text" class="form-control" placeholder="root">
                        </div>
                        <div v-if="setup.db_connection !== 'sqlite'" class="form-group">
                            <label for="db_password" class="text-capitalize">{{ $t('password') }}</label>
                            <input @keyup="canSubmit" v-model="setup.db_password" id="db_password" type="password" class="form-control" placeholder="password123">
                        </div>
                        <div v-if="setup.db_connection !== 'sqlite'" class="form-group">
                            <label for="db_database" class="text-capitalize">{{ $t('setup.database') }}</label>
                            <input @keyup="canSubmit" v-model="setup.db_database" id="db_database" type="text" class="form-control" placeholder="Database name">
                        </div>

                        <div class="form-group mt-3">
                            <div class="row">
                                <div class="col-9">
                                    <span class="text-left text-capitalize">{{ $t('setup.send_reports') }}</span>
                                </div>
                                <div class="col-3 text-right">
                                    <span @click="setup.send_reports = !!setup.send_reports" class="switch">
                                      <input v-model="setup.send_reports" type="checkbox" name="send_reports" class="switch" id="send_reports" :checked="setup.send_reports">
                                      <label for="send_reports"></label>
                                    </span>
                                </div>
                            </div>

                        </div>

                    </div>

                    <div class="col-6">

                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.project_name') }}</label>
                            <input @keyup="canSubmit" v-model="setup.project" id="project" type="text" class="form-control" placeholder="Work Servers" required>
                        </div>

                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.project_description') }}</label>
                            <input @keyup="canSubmit" v-model="setup.description" id="description" type="text" class="form-control" placeholder="Monitors all of my work services">
                        </div>

                        <div class="form-group">
                            <label class="text-capitalize" for="domain">{{ $t('setup.domain') }}</label>
                            <input @keyup="canSubmit" v-model="setup.domain" type="text" class="form-control" id="domain" required>
                        </div>

                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.username') }}</label>
                            <input @keyup="canSubmit" v-model="setup.username" id="username" type="text" class="form-control" placeholder="admin" required>
                        </div>

                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.password') }}</label>
                            <input @keyup="canSubmit" v-model="setup.password" id="password" type="password" class="form-control" placeholder="password" required>
                        </div>

                        <div class="form-group">
                            <label class="text-capitalize">{{ $t('setup.password_confirm') }}</label>
                            <input @keyup="canSubmit" v-model="setup.confirm_password" id="password_confirm" type="password" class="form-control" placeholder="password" required>
                            <span v-if="passnomatch" class="small text-danger">Both passwords should match</span>
                        </div>

                        <div class="form-group">
                            <div class="row">
                                <div class="col-8">
                                    <label class="text-capitalize">{{ $t('email') }}</label>
                                    <input @keyup="canSubmit" v-model="setup.email" id="email" type="text" class="form-control" placeholder="myemail@domain.com">
                                </div>
                                <div class="col-4 text-right">
                                    <label class="d-none d-sm-block text-capitalize text-capitalize">{{ $t('setup.newsletter') }}</label>
                                    <span @click="setup.newsletter = !!setup.newsletter" class="switch">
                                      <input v-model="setup.newsletter" type="checkbox" name="send_newsletter" class="switch" id="send_newsletter" :checked="setup.newsletter">
                                      <label for="send_newsletter"></label>
                                    </span>
                                </div>
                            </div>
                            <small>{{ $t('setup.newsletter_note') }}</small>
                        </div>
                    </div>

                    <div v-if="error" class="col-12 alert alert-danger">
                        {{error}}
                    </div>

                    <button @click.prevent="saveSetup" v-bind:disabled="disabled || loading" type="submit" class="btn btn-primary btn-block" :class="{'btn-primary': !loading, 'btn-default': loading}">
                        <font-awesome-icon v-if="loading" icon="circle-notch" class="mr-2" spin/>{{loading ? "Loading..." : "Save Settings"}}
                    </button>
                </div>
            </form>

        </div>
    </div>
</template>

<script>
  import Api from "../API";
  import Index from "../pages/Index";

  export default {
  name: 'Setup',
  data () {
    return {
      error: null,
      loading: false,
      disabled: true,
      passnomatch: false,
      setup: {
        language: "en",
        db_connection: "sqlite",
        db_host: "",
        db_port: "",
        db_user: "",
        db_password: "",
        db_database: "",
        project: "",
        description: "",
        domain: "",
        username: "",
        password: "",
        confirm_password: "",
        sample_data: true,
        newsletter: true,
        send_reports: true,
        email: "",
      }
    }
  },
  async created() {
    const core = await Api.core()
    if (core.setup) {
        if (!this.$store.getters.hasPublicData) {
            await this.$store.dispatch('loadRequired')
        }
        this.$router.push('/')
    }
  },
  mounted() {
    this.changeLanguages()
    this.setup.domain = window.location.protocol + "//" + window.location.hostname + (window.location.port ? ":"+window.location.port : "")
  },
  methods: {
    changeLanguages() {
      this.$i18n.locale = this.setup.language
    },
      canSubmit() {
          this.error = null
          const s = this.setup
        if (s.confirm_password.length > 0 && s.confirm_password !== s.password) {
          this.passnomatch = true
        } else {
          this.passnomatch = false
        }
          if (s.db_connection !== 'sqlite') {
              if (!s.db_host || !s.db_port || !s.db_user || !s.db_password || !s.db_database) {
                  this.disabled = true
                  return
              }
          }
          if (!s.project || !s.domain || !s.username || !s.password || !s.confirm_password) {
              this.disabled = true
              return
          }
          if (s.password !== s.confirm_password) {
              this.disabled = true
              return
          }
          this.disabled = false
      },
    async saveSetup() {
      this.loading = true
      let resp
      try {
        resp = await Api.setup_save(this.setup)
      } catch(e) {
        resp = {status: 'error', error: e.response.data.error}
      }
      if (resp.status === 'error') {
        this.error = resp.error
        this.loading = false
        return
      }

      await this.$store.dispatch('loadCore')
      await this.$store.dispatch('loadRequired')

      this.loading = false
      this.$router.push('/')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
