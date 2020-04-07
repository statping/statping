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

    </form>
</template>

<script>
  import Api from '../API'

  export default {
      name: 'CoreSettings',
        props: {
          in_core: {
            type: Object,
            required: true,
          }
        },
    data() {
      return {
        core: this.in_core
      }
    },
      methods: {
          async saveSettings() {
              const c = this.core
              await Api.core_save(c)
              const core = await Api.core()
              this.$store.commit('setCore', core)
              this.core = core
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
