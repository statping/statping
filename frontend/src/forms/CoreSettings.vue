<template>
    <form @submit.prevent="saveSettings">
        <div class="form-group">
            <label>Project Name</label>
            <input v-model="core.name" type="text" class="form-control" placeholder="Great Uptime" id="project">
        </div>

        <div class="form-group">
            <label>Project Description</label>
            <input v-model="core.description" type="text" class="form-control" placeholder="Great Uptime" id="description">
        </div>

        <div class="form-group row">
            <div class="col-8 col-sm-9">
                <label>Domain</label>
                <input v-model="core.domain" type="url" class="form-control" id="domain">
            </div>
            <div class="col-4 col-sm-3 mt-sm-1 mt-0">
                <label class="d-inline d-sm-none">Enable CDN</label>
                <label class="d-none d-sm-block">Enable CDN</label>
                <span @click="core.using_cdn = !!core.using_cdn" class="switch" id="using_cdn">
                    <input v-model="core.using_cdn" type="checkbox" name="using_cdn" class="switch" id="switch-normal" :checked="core.using_cdn">
                    <label for="switch-normal"></label>
                  </span>
            </div>
        </div>

        <div class="form-group">
            <label>Custom Footer</label>
            <textarea v-model="core.footer" rows="4" class="form-control" id="footer">{{core.footer}}</textarea>
            <small class="form-text text-muted">HTML is allowed inside the footer</small>
        </div>

        <div class="form-group">
            <label>Language</label>
            <select v-model="core.language" class="form-control">
                <option value="en">English</option>
                <option value="es">Spanish</option>
                <option value="fr">French</option>
                <option value="ru">Russian</option>
                <option value="de">German</option>
            </select>
            <small class="form-text text-muted">HTML is allowed inside the footer</small>
        </div>

        <div class="form-group row mt-3">
            <label class="col-sm-10 col-form-label">Enable Error Reporting</label>
            <div class="col-sm-2 float-right">
                <span @click="core.allow_reports = !!core.allow_reports" class="switch" id="allow_report">
                <input v-model="core.allow_reports" type="checkbox" name="allow_report" class="switch" id="switch_allow_report" :checked="core.allow_reports">
                <label for="switch_allow_report"></label>
              </span>
            </div>
            <div class="col-12">
                <small>Help the Statping project out by sending anonymous error logs back to our server.</small>
            </div>
        </div>

        <button @click.prevent="saveSettings" id="save_core" type="submit" class="btn btn-primary btn-block mt-3" v-bind:disabled="loading">
            <font-awesome-icon v-if="loading" icon="circle-notch" class="mr-2" spin/>Save Settings
        </button>

    </form>
</template>

<script>
  import Api from '../API'

  export default {
      name: 'CoreSettings',
    data () {
      return {
        loading: false
      }
    },
      computed: {
          core() {
              return this.$store.getters.core
          }
      },
      methods: {
          async saveSettings() {
            this.loading = true
              const c = this.core
              await Api.core_save(c)
              const core = await Api.core()
              this.$store.commit('setCore', core)
            setInterval(() => { this.loading = false }, 1500)
          },
          selectAll() {
              this.$refs.input.select();
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
