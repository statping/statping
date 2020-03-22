<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">
        <div class="col-12 col-md-8 offset-md-2 mb-4">
            <img class="col-12 mt-5 mt-md-0" style="max-width:680px" src="/banner.png">
        </div>

        <div class="col-12">

    <form @submit.prevent="saveSetup">
        <div class="row">
            <div class="col-6">
                <div class="form-group">
                    <label>Database Connection</label>
                    <select @change="canSubmit" v-model="setup.db_connection" class="form-control">
                        <option value="sqlite">Sqlite</option>
                        <option value="postgres">Postgres</option>
                        <option value="mysql">MySQL</option>
                    </select>
                </div>
                <div v-if="setup.db_connection !== 'sqlite'" class="form-group" id="db_host">
                    <label>Host</label>
                    <input @keyup="canSubmit" v-model="setup.db_host" type="text" class="form-control" placeholder="localhost">
                </div>
                <div v-if="setup.db_connection !== 'sqlite'" class="form-group" id="db_port">
                    <label>Database Port</label>
                    <input @keyup="canSubmit" v-model="setup.db_port" type="text" class="form-control" placeholder="localhost">
                </div>
                <div v-if="setup.db_connection !== 'sqlite'" class="form-group" id="db_user">
                    <label>Username</label>
                    <input @keyup="canSubmit" v-model="setup.db_user" type="text" class="form-control" placeholder="root">
                </div>
                <div v-if="setup.db_connection !== 'sqlite'" class="form-group" id="db_password">
                    <label for="db_password">Password</label>
                    <input @keyup="canSubmit" v-model="setup.db_password" type="password" class="form-control" placeholder="password123">
                </div>
                <div v-if="setup.db_connection !== 'sqlite'" class="form-group" id="db_database">
                    <label for="db_database">Database</label>
                    <input @keyup="canSubmit" v-model="setup.db_database" type="text" class="form-control" placeholder="Database name">
                </div>

            </div>

            <div class="col-6">

                <div class="form-group">
                    <label>Project Name</label>
                    <input @keyup="canSubmit" v-model="setup.project" type="text" class="form-control" placeholder="Great Uptime" required>
                </div>

                <div class="form-group">
                    <label>Project Description</label>
                    <input @keyup="canSubmit" v-model="setup.description" type="text" class="form-control" placeholder="Great Uptime">
                </div>

                <div class="form-group">
                    <label for="domain_input">Domain URL</label>
                    <input @keyup="canSubmit" v-model="setup.domain" type="text" class="form-control" id="domain_input" required>
                </div>

                <div class="form-group">
                    <label>Admin Username</label>
                    <input @keyup="canSubmit" v-model="setup.username" type="text" class="form-control" placeholder="admin" required>
                </div>

                <div class="form-group">
                    <label>Admin Password</label>
                    <input @keyup="canSubmit" v-model="setup.password" type="password" class="form-control" placeholder="password" required>
                </div>

                <div class="form-group">
                    <label>Confirm Admin Password</label>
                    <input @keyup="canSubmit" v-model="setup.confirm_password" type="password" class="form-control" placeholder="password" required>
                </div>

            </div>

            <div v-if="error" class="col-12 alert alert-danger">
                {{error}}
            </div>

            <button @click.prevent="saveSetup" v-bind:disabled="disabled || loading" type="submit" class="btn btn-primary btn-block" :class="{'btn-primary': !loading, 'btn-default': loading}">
               {{loading ? "Loading..." : "Save Settings"}}
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
      setup: {
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
        sample_data: true
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
    this.setup.domain = window.location.protocol + "//" + window.location.hostname + (window.location.port ? ":"+window.location.port : "")
  },
  methods: {
      canSubmit() {
          this.error = null
          const s = this.setup
          if (s.db_connection !== 'sqlite') {
              if (!s.db_host || !s.db_port || !s.db_user || !s.db_password || !s.db_database) {
                  this.disabled = true
                  return
              }
          }
          if (!s.project || !s.description || !s.domain || !s.username || !s.password || !s.confirm_password) {
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
      const s = this.setup
      if (s.password !== s.confirm_password) {
        alert('Passwords do not match!')
        this.loading = false
        return
      }
      const resp = await Api.setup_save(s)
      if (resp.status === 'error') {
        this.error = resp.error
        this.loading = false
        return
      }

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
