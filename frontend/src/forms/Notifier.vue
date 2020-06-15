<template>
    <div>
        <form @submit.prevent="saveNotifier">
    <div class="card contain-card text-black-50 bg-white mb-3">
        <div class="card-header text-capitalize">
            {{notifier.title}}
            <span @click="enableToggle" class="switch switch-rd-gr float-right">
                <input v-model="notifier.enabled" type="checkbox" class="switch-sm" :id="`enable_${notifier.method}`" v-bind:checked="notifier.enabled">
                <label class="mb-0" :for="`enable_${notifier.method}`"></label>
            </span>
        </div>
        <div class="card-body">

        <p class="small text-muted" v-html="notifier.description"/>

        <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>
            <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[form.field.toLowerCase()]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <small class="form-text text-muted" v-html="form.small_text"></small>
        </div>

        <div class="row mt-4">

            <div class="col-sm-12">
                <span class="slider-info">Limit {{notifier.limits}} per hour</span>
                <input v-model="notifier.limits" type="range" name="limits" class="slider" min="1" max="300">
                <small class="form-text text-muted">Notifier '{{notifier.title}}' will send a maximum of {{notifier.limits}} notifications per hour.</small>
            </div>

        </div>
        </div>
    </div>

        <div v-if="notifier.data_type" class="card text-black-50 bg-white mb-3">
            <div class="card-header text-capitalize">
                {{notifier.title}} Outgoing Request
                <span class="badge badge-dark float-right text-uppercase mt-1">{{notifier.data_type}}</span>
            </div>
            <div class="card-body">
                <span class="text-muted d-block mb-3" v-if="notifier.request_info" v-html="notifier.request_info"></span>

        <div class="row" v-observe-visibility="visible">
            <div class="col-12">
                <h5 class="text-capitalize">Success Data</h5>
                <codemirror v-model="success_data"
                            ref="cmsuccess"
                            :options="cmOptions"
                            @ready="onCmSuccessReady"/>
            </div>
        </div>

        <div class="row mt-4">
            <div class="col-12">
                <h5 class="text-capitalize">Failure Data</h5>
                <codemirror v-model="failure_data"
                            ref="cmfailure"
                            :options="cmOptions"
                            @ready="onCmFailureReady"/>
            </div>
        </div>

            </div>
        </div>

    </form>

        <div v-if="error && !success" class="card text-black-50 bg-white mb-3">
            <div class="card-body">

            <div v-if="error && !success" class="alert alert-danger col-12" role="alert">
                {{error}}<p v-if="response">Response:<br>{{response}}</p>
            </div>
            <div v-if="success" class="alert alert-success col-12" role="alert">
                {{notifier.title}} appears to be working!
                <p v-if="response">Response:<br>{{response}}</p>
            </div>

            </div>
        </div>

        <div class="card text-black-50 bg-white mb-3">
            <div class="card-body">

                <div class="row">
                    <div class="col-4 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="saveNotifier" type="submit" class="btn btn-block text-capitalize btn-primary save-notifier">
                            <i class="fa fa-check-circle"></i> {{loading ? "Loading..." : saved ? "Saved" : "Save Settings"}}
                        </button>
                    </div>
                    <div class="col-4 col-md-4">
                        <button @click.prevent="testNotifier" class="btn btn-outline-dark btn-block text-capitalize test-notifier">
                            <i class="fa fa-vial"></i>{{loadingTest ? "Loading..." : "Test Success"}}</button>
                    </div>
                    <div class="col-4 col-md-4">
                        <button @click.prevent="testNotifier" class="btn btn-outline-dark btn-block text-capitalize test-notifier">
                            <i class="fa fa-vial"></i>{{loadingTest ? "Loading..." : "Test Failure"}}</button>
                    </div>
                </div>

            </div>
        </div>

        <span class="d-block small text-center mb-3">
            <span class="text-capitalize">{{notifier.title}}</span> Notifier created by <a :href="notifier.author_url" target="_blank">{{notifier.author}}</a>
        </span>

    </div>
</template>

<script>
  import Api from "../API";

  const beautify = require('js-beautify').js

  // require component
  import { codemirror } from 'vue-codemirror'
  import 'codemirror/mode/javascript/javascript.js'
  import 'codemirror/lib/codemirror.css'
  import 'codemirror/theme/neat.css'
  import '../codemirror_json'

export default {
    name: 'Notifier',
  components: {
    codemirror
  },
    props: {
        notifier: {
            type: Object,
            required: true
        }
    },
  watch: {

  },
    data() {
        return {
            loading: false,
            loadingTest: false,
            error: null,
            response: null,
            success: false,
            saved: false,
          success_data: null,
          failure_data: null,
            form: {},
              cmOptions: {
                height: 700,
                tabSize: 2,
                lineNumbers: true,
                line: true,
                class: "json-field",
                theme: 'neat',
                mode: "mymode",
                lineWrapping: true,
                json: true,
                autoRefresh: true,
                mime: this.notifier.data_type === "json" ? "application/json" : "text/plain"
              },
          beautifySettings: { indent_size: 2, space_in_empty_paren: true },
        }
    },
      computed: {

      },
    methods: {
      visible(isVisible, entry) {
        if (isVisible) {
          this.$refs.cmfailure.codemirror.refresh()
          this.$refs.cmsuccess.codemirror.refresh()
        }
      },
      onCmSuccessReady(cm) {
        this.success_data = beautify(this.notifier.success_data, this.beautifySettings)
        console.log('the editor is ready!', cm)
        setTimeout(function() {
          cm.refresh();
        },1);
      },
      onCmFailureReady(cm) {
        this.failure_data = beautify(this.notifier.failure_data, this.beautifySettings)
        setTimeout(function() {
          cm.refresh();
        },1);
      },
      onCmFocus(cm) {
        console.log('the editor is focused!', cm)
      },
      onCmCodeChange(newCode) {
        console.log('this is new code', newCode)
        this.success_data = newCode
      },
        async enableToggle() {
            this.notifier.enabled = !!this.notifier.enabled
            const form = {
                enabled: !this.notifier.enabled,
                method: this.notifier.method,
            }
            await Api.notifier_save(form)
        },
        async saveNotifier() {
            this.loading = true
            this.form.enabled = this.notifier.enabled
            this.form.limits = parseInt(this.notifier.limits)
            this.form.method = this.notifier.method
            this.notifier.form.forEach((f) => {
              let field = f.field.toLowerCase()
              let val = this.notifier[field]
              if (this.isNumeric(val)) {
                val = parseInt(val)
              }
                this.form[field] = val
            });
          this.form.success_data = this.success_data
          this.form.failure_data = this.failure_data
          window.console.log(this.form)
            await Api.notifier_save(this.form)
            const notifiers = await Api.notifiers()
            await this.$store.commit('setNotifiers', notifiers)
            this.saved = true
            this.loading = false
        },
        async testNotifier() {
            this.success = false
            this.loadingTest = true
            this.form.method = this.notifier.method
            this.notifier.form.forEach((f) => {
                let field = f.field.toLowerCase()
                let val = this.notifier[field]
                if (this.isNumeric(val)) {
                    val = parseInt(val)
                }
                this.form[field] = val
            });
            const tested = await Api.notifier_test(this.form)
            if (tested.success) {
                this.success = true
            } else {
                this.error = tested.error
            }
            this.response = tested.response
            this.loadingTest = false
        },
    }
}
</script>
<style scoped>
    .CodeMirror {
        border: 1px solid #eee;
        height: 550px;
        font-size: 9pt;
    }
</style>
