<template>
    <div>
        <form @submit.prevent="saveNotifier">
    <div class="card contain-card mb-3">
        <div class="card-header text-capitalize">
            {{notifier.title}}
            <span @click="enableToggle" class="switch switch-sm switch-rd-gr float-right">
                <input v-model="notifier.enabled" type="checkbox" :id="`enable_${notifier.method}`" v-bind:checked="notifier.enabled">
                <label class="mb-0" :for="`enable_${notifier.method}`"></label>
            </span>
        </div>
        <div class="card-body">
        <p class="small text-muted" v-html="notifier.description"/>

            <div v-if="notifier.method==='mobile'">
                <div class="form-group row mt-3">
                    <label for="statping_domain" class="col-sm-4 col-form-label">Statping Domain</label>
                    <div class="col-sm-8">
                        <div class="input-group">
                            <input v-bind:value="$store.getters.core.domain" type="text" class="form-control" id="statping_domain" readonly>
                            <div class="input-group-append copy-btn">
                                <button @click.prevent="copy($store.getters.core.domain)" class="btn btn-outline-secondary" type="button">Copy</button>
                            </div>
                        </div>
                    </div>
                </div>
            <div class="form-group row mt-3">
                <label for="apisecret" class="col-sm-4 col-form-label">API Secret</label>
                <div class="col-sm-8">
                    <div class="input-group">
                        <input v-bind:value="$store.getters.core.api_secret" type="text" class="form-control" id="apisecret" readonly>
                        <div class="input-group-append copy-btn">
                            <button @click.prevent="copy($store.getters.core.api_secret)" class="btn btn-outline-secondary" type="button">Copy</button>
                        </div>
                    </div>
                </div>
            </div>
                <div class="col-12 col-md-6 offset-0 offset-md-3">
                    <img :src="qrcode" class="img-thumbnail">
                    <span class="text-muted small center">Scan this QR Code on the Statping Mobile App for quick setup</span>
                </div>
            </div>

        <div v-if="notifier.method!=='mobile'" v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>
            <input v-if="formVisible(['text', 'number', 'password', 'email'], form)" v-model="notifier[form.field.toLowerCase()]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <select v-if="formVisible(['list'], form)" v-model="notifier[form.field.toLowerCase()]" class="form-control">
                <option v-for="(val, k) in form.list_options" :value="val">{{val}}</option>
            </select>

            <span v-if="formVisible(['switch'], form)" @click="notifier[form.field.toLowerCase()] = !!notifier[form.field.toLowerCase()]" class="switch switch-rd-gr float-right mt-2">
                <input v-model="notifier[form.field.toLowerCase()]" type="checkbox" class="switch-sm" :id="`switch_${notifier.name}_${form.field}`" v-bind:checked="notifier[form.field.toLowerCase()]">
                <label class="mb-0" :for="`switch_${notifier.name}_${form.field}`"></label>
            </span>

            <small class="form-text text-muted" v-html="form.small_text"></small>
        </div>

        <div class="row mt-4">

            <div class="col-sm-12">
                <span class="slider-info">Limit {{notifier.limits}} per hour</span>
                <input v-model.number="notifier.limits" type="range" name="limits" class="slider" min="1" max="300">
                <small class="form-text text-muted">Notifier '{{notifier.title}}' will send a maximum of {{notifier.limits}} notifications per hour.</small>
            </div>

        </div>
        </div>
    </div>

                    <div v-if="notifier.data_type" class="card mb-3">
            <div class="card-header text-capitalize">
                <font-awesome-icon @click="expanded = !expanded" :icon="expanded ? 'minus' : 'plus'" class="mr-2 pointer"/>
                {{notifier.title}} Outgoing Request
                <span class="badge badge-dark float-right text-uppercase mt-1">{{notifier.data_type}}</span>
            </div>
            <div class="card-body" :class="{'d-none': !expanded}">
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

        <div v-if="error || success" class="card mb-3">
            <div class="card-body">

            <div v-if="error && !success" class="alert alert-danger col-12" role="alert">
                {{error}}
            </div>
            <div v-if="success" class="alert alert-success col-12" role="alert">
                <span class="text-capitalize">{{notifier.title}}</span> appears to be working!
            </div>

                <h5>Response</h5>
                <codemirror :value="response"/>

            </div>
        </div>

        <div class="card mb-3">
            <div class="card-body">
                <div class="row">
                    <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="saveNotifier" :disabled="loading" type="submit" class="btn btn-block text-capitalize btn-primary save-notifier">
                            <font-awesome-icon v-if="loading" icon="circle-notch" class="mr-2" spin/> {{loading ? "Loading..." : saved ? "Saved" : "Save"}}
                        </button>
                    </div>
                    <div class="col-12 col-md-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="testNotifier('success')" :disabled="loadingTest" class="btn btn-secondary btn-block text-capitalize test-notifier">
                            <font-awesome-icon v-if="loadingTest" icon="circle-notch" class="mr-2" spin/>{{loadingTest ? "Loading..." : "Test Success"}}</button>
                    </div>
                    <div class="col-12 col-md-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="testNotifier('failure')" :disabled="loadingTest" class="btn btn-secondary btn-block text-capitalize test-notifier">
                            <font-awesome-icon v-if="loadingTest" icon="circle-notch" class="mr-2" spin/>{{loadingTest ? "Loading..." : "Test Failure"}}</button>
                    </div>
                </div>
            </div>
        </div>

        <div v-if="notifier.logs" class="card mb-3">
            <div class="card-header text-capitalize">
                <font-awesome-icon @click="expanded_logs = !expanded_logs" :icon="expanded_logs ? 'minus' : 'plus'" class="mr-2 pointer"/>
                {{notifier.title}} Logs
                <span class="badge badge-info float-right text-uppercase mt-1">{{notifier.logs.length}}</span>
            </div>
            <div class="card-body" :class="{'d-none': !expanded_logs}">
                <div v-for="(log, i) in notifier.logs" class="alert" :class="{'alert-danger': log.error, 'alert-dark': !log.success && !log.error, 'alert-success': log.success && !log.error}">
                        <span class="d-block">
                            Service {{log.service}}
                            {{log.success ? "Success Triggered" : "Failure Triggered"}}
                        </span>

                    <div v-if="log.message !== ''" class="bg-white p-3 small mt-2">
                        <code>{{log.message}}</code>
                    </div>

                    <div class="row mt-2">
                        <span class="col-6 small">{{niceDate(log.created_at)}}</span>
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
/* webpackChunkName: "codemirror" */
import {codemirror} from 'vue-codemirror'
/* webpackChunkName: "codemirror" */
import 'codemirror/mode/javascript/javascript.js'
/* webpackChunkName: "codemirror" */
import 'codemirror/lib/codemirror.css'
/* webpackChunkName: "codemirror" */
import 'codemirror/theme/neat.css'
/* webpackChunkName: "codemirror" */
import '../codemirror_json'

const beautify = require('js-beautify').js

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
  watch: {},
  data() {
    return {
      loading: false,
      loadingTest: false,
      error: null,
      response: null,
      request: null,
      success: false,
      saved: false,
      expanded: false,
      expanded_logs: false,
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
        json: this.notifier.data_type === "json",
        autoRefresh: true,
        mime: this.notifier.data_type === "json" ? "application/json" : "text/plain"
      },
      beautifySettings: {indent_size: 2, space_in_empty_paren: true},
    }
  },
  computed: {
    core() {
      return this.$store.getters.core
    },
    qrcode() {
      const u = `statping://setup?domain=${this.core.domain}&api=${this.core.api_secret}`
      return "https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + encodeURIComponent(u)
    }
  },
  methods: {
    formVisible(want, form) {
      return !!want.includes(form.type);
    },
    visible(isVisible, entry) {
      if (isVisible) {
        this.$refs.cmfailure.codemirror.refresh()
        this.$refs.cmsuccess.codemirror.refresh()
      }
    },
    onCmSuccessReady(cm) {
      this.success_data = this.notifier.success_data
      if (this.notifier.data_type === "json") {
        this.success_data = beautify(this.notifier.success_data, this.beautifySettings)
      }
      setTimeout(function () {
        cm.refresh();
      }, 1);
    },
    onCmFailureReady(cm) {
      this.failure_data = this.notifier.failure_data
      if (this.notifier.data_type === "json") {
        this.failure_data = beautify(this.notifier.failure_data, this.beautifySettings)
      }
      setTimeout(function () {
        cm.refresh();
      }, 1);
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
      if (this.notifier.form) {
        this.notifier.form.forEach((f) => {
          let field = f.field.toLowerCase()
          let val = this.notifier[field]
          if (this.isNumeric(val) && this.form.method!='telegram') {
            val = parseInt(val)
          }
          this.form[field] = val
        });
      }
      this.form.success_data = this.success_data
      this.form.failure_data = this.failure_data
      await Api.notifier_save(this.form)
      const notifiers = await Api.notifiers()
      await this.$store.commit('setNotifiers', notifiers)
      this.saved = true
      this.loading = false
    },
    async testNotifier(method = "success") {
      this.success = false
      this.loadingTest = true
      this.form.method = this.notifier.method
      if (this.notifier.form) {
        this.notifier.form.forEach((f) => {
          let field = f.field.toLowerCase()
          let val = this.notifier[field]
          if (this.isNumeric(val) && this.form.method!='telegram') {
            val = parseInt(val)
          }
          this.form[field] = val
        });
      }
      let req = {
        notifier: this.form,
        method: method,
      }
      const tested = await Api.notifier_test(req, this.notifier.method)
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
