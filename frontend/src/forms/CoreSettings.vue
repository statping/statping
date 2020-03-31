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

        <button @click.prevent="saveSettings" id="save_core" type="submit" class="btn btn-primary btn-block">Save Settings</button>

        <div class="form-group row mt-5">
            <label class="col-sm-3 col-form-label">API Key</label>
            <div class="col-sm-9">
                <input v-model="core.api_key" @focus="$event.target.select()" type="text" class="form-control select-input" id="api_key" readonly>
                <small class="form-text text-muted">API Key can be used for read only routes</small>
            </div>
        </div>

        <div class="form-group row">
            <label class="col-sm-3 col-form-label">API Secret</label>
            <div class="col-sm-9">
                <input v-model="core.api_secret" @focus="$event.target.select()" type="text" class="form-control select-input" id="api_secret" readonly>
                <small class="form-text text-muted">API Secret is used for read, create, update and delete routes</small>
                <small class="form-text text-muted">You can <a href="#" @click="renewApiKeys">Regenerate API Keys</a> if you need to.</small>
            </div>
        </div>

        <div class="row d-none">
        <div class="col-12">
        <h4 class="mt-5">Github Authentication</h4>

        <div class="form-group row d-none">
            <label class="col-sm-4 col-form-label">Github Client ID</label>
            <div class="col-sm-8">
                <input v-model="core.github_clientId" type="text" class="form-control" placeholder="" required>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Github Client Secret</label>
            <div class="col-sm-8">
                <input v-model="core.github_clientScret" type="text" class="form-control" placeholder="" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="switch-group-public" class="col-sm-4 col-form-label">Enabled</label>
            <div class="col-md-8 col-xs-12 mt-1">
            <span @click="enabled = !!enabled" class="switch float-left">
                <input v-model="enabled" type="checkbox" class="switch" id="switch-group-public" :checked="enabled">
                <label for="switch-group-public">Enabled Github Auth</label>
            </span>
            </div>
        </div>

            <button @click.prevent="saveSettings" type="submit" class="btn btn-primary btn-block">Save Settings</button>

        </div>
        </div>

    </form>
</template>

<script>
  import Api from '../API'

  export default {
      name: 'CoreSettings',
    props: {
      core: {
        type: Object,
        required: true,
      }
    },
      methods: {
          async saveSettings() {
              const c = this.core
              const coreForm = {
                  name: c.name, description: c.description, domain: c.domain,
                  timezone: c.timezone, using_cdn: c.using_cdn, footer: c.footer, update_notify: c.update_notify,
                  gh_client_id: c.github_clientId, gh_client_secret: c.github_clientSecret
              };
              await Api.core_save(coreForm)
              const core = await Api.core()
              this.$store.commit('setCore', core)
              this.core = core
          },
          async renewApiKeys() {
              let r = confirm("Are you sure you want to reset the API keys?");
              if (r === true) {
                  await Api.renewApiKeys()
                  const core = await Api.core()
                  this.$store.commit('setCore', core)
                  this.core = core
              }
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
