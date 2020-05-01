<template >
    <div>
            <div v-if="loaded && !directory" class="jumbotron jumbotron-fluid">
                <div class="text-center col-12">
                    <h1 class="display-5">Enable Local Assets</h1>
                    <span class="lead">Customize your status page design by enabling local assets. This will create a 'assets' directory containing all CSS.<p>
                        <button id="enable_assets" @click.prevent="createAssets" :disabled="pending" href="#" class="btn btn-primary mt-3">Enable Local Assets</button>
                    </p></span>
                </div>
        </div>
    <form v-if="loaded && directory" @submit.prevent="saveAssets" :disabled="pending">
        <h3>Variables</h3>
        <codemirror v-show="loaded" v-model="vars" ref="vars" :options="cmOptions" class="codemirrorInput"/>

        <h3 class="mt-3">Base Theme</h3>
        <codemirror v-show="loaded" v-model="base" ref="base" :options="cmOptions" class="codemirrorInput"/>

        <h3 class="mt-3">Mobile Overwrites</h3>
        <codemirror v-show="loaded" v-model="mobile" ref="mobile" :options="cmOptions" class="codemirrorInput"/>

        <div v-if="error" class="alert alert-danger mt-3" style="white-space: pre-line;">{{error}}</div>

        <button id="save_assets" @submit.prevent="saveAssets" type="submit" class="btn btn-primary btn-block mt-2" :disabled="pending">{{pending ? "Saving..." : "Save Style"}}</button>
        <button id="delete_assets" v-if="directory" @click.prevent="deleteAssets" href="#" class="btn btn-danger btn-block confirm-btn" :disabled="pending">Delete Local Assets</button>

        <h6 id="assets_dir" class="text-muted text-monospace text-sm-center font-1 mt-3">
            Asset Directory: {{directory}}
        </h6>
    </form>
    </div>
</template>

<script>
  import Api from "../../API";

  // require component
  import { codemirror } from 'vue-codemirror'
  import 'codemirror/mode/css/css.js'

  import 'codemirror/lib/codemirror.css'
  import 'codemirror-colorpicker/dist/codemirror-colorpicker.css'
  import 'codemirror-colorpicker'

  export default {
      name: 'ThemeEditor',
      components: {
          codemirror
      },
      computed: {
          core() {
              return this.$store.getters.core
          }
      },
      data () {
          return {
              base: null,
              vars: null,
              mobile: null,
              error: null,
              directory: null,
              tab: "vars",
              loaded: false,
              pending: false,
              cmOptions: {
                  height: 700,
                  tabSize: 4,
                  lineNumbers: true,
                  matchBrackets: true,
                  mode: "text/x-scss",
                  line: true,
                  colorpicker: true
              }
          }
      },
      async mounted () {
          await this.fetchTheme()
          this.changeTab('vars')
      },
      methods: {
          async fetchTheme() {
              this.loaded = true
              this.pending = true
              const theme = await Api.theme()
              this.directory = theme.directory
              if (this.directory) {
                  this.base = theme.base
                  this.vars = theme.variables
                  this.mobile = theme.mobile
              }
              this.pending = false
              this.loaded = true
          },
          async createAssets() {
              this.pending = true
              const resp = await Api.theme_generate(true)
              this.pending = false
              await this.fetchTheme()
          },
          async deleteAssets() {
              this.pending = true
              let c = confirm('Are you sure you want to delete all local assets?')
              if (c) {
                  const resp = await Api.theme_generate(false)
                  await this.fetchTheme()
              }
              this.pending = false
          },
          async saveAssets() {
              this.pending = true
              const data = {base: this.base, variables: this.vars, mobile: this.mobile}
              const resp = await Api.theme_save(data)
              if (resp.error) {
                  this.error = resp.error
                  this.pending = false
                  return
              } else {
                  this.error = null
              }
              this.pending = false
              await this.fetchTheme()
          },
          changeTab (v) {
              this.tab = v
              if (v === 'base') {
                  this.$refs.base.codemirror.refresh();
              } else if (v === 'vars') {
                  this.$refs.vars.codemirror.refresh();
              } else if (v === 'mobile') {
                  this.$refs.mobile.codemirror.refresh();
              }
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
    .CodeMirror {
        border: 1px solid #eee;
        height: 550px;
    }
</style>
