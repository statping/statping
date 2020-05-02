<template>
    <div class="col-12">
        <div class="row">
            <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
                <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
                    <h6 class="text-muted">Main Settings</h6>

                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-home-tab')}" id="v-pills-home-tab" data-toggle="pill" href="#v-pills-home" role="tab" aria-controls="v-pills-home" aria-selected="true">
                        <font-awesome-icon icon="cog" class="mr-2"/> Settings
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-style-tab')}" id="v-pills-style-tab" data-toggle="pill" href="#v-pills-style" role="tab" aria-controls="v-pills-style" aria-selected="false">
                        <font-awesome-icon icon="image" class="mr-2"/> Theme Editor
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-cache-tab')}" id="v-pills-cache-tab" data-toggle="pill" href="#v-pills-cache" role="tab" aria-controls="v-pills-cache" aria-selected="false">
                        <font-awesome-icon icon="paperclip" class="mr-2"/> Cache
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-oauth-tab')}" id="v-pills-oauth-tab" data-toggle="pill" href="#v-pills-oauth" role="tab" aria-controls="v-pills-oauth" aria-selected="false">
                        <font-awesome-icon icon="key" class="mr-2"/> OAuth <span class="mt-1 float-right badge badge-light text-dark font-1">BETA</span>
                    </a>

                    <h6 class="mt-4 text-muted">Notifiers</h6>

                    <div id="notifiers_tabs">
                        <a v-for="(notifier, index) in notifiers" v-bind:key="`${notifier.method}_${index}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                            <font-awesome-icon :icon="iconName(notifier.icon)" class="mr-2"/> {{notifier.title}}
                            <span v-if="notifier.enabled" class="badge badge-pill float-right mt-1" :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}">ON</span>
                        </a>
                        <a @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-notifier-docs-tab`)}" v-bind:id="`v-pills-notifier-docs-tab`" data-toggle="pill" v-bind:href="`#v-pills-notifier-docs`" role="tab" v-bind:aria-controls="`v-pills-notifier-docs`" aria-selected="false">
                            <font-awesome-icon icon="question" class="mr-2"/> Variables
                        </a>
                    </div>

                    <h6 class="mt-4 mb-3 text-muted">Statping Links</h6>

                    <a href="https://github.com/statping/statping/wiki" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="question" class="mr-3"/> Documentation
                    </a>

                    <a href="https://github.com/statping/statping/wiki/API" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="laptop" class="mr-2"/> API Documentation
                    </a>

                    <a href="https://raw.githubusercontent.com/statping/statping/master/CHANGELOG.md" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="book" class="mr-3"/> Changelog
                    </a>

                    <a href="https://github.com/statping/statping" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="code-branch" class="mr-3"/> Statping Github Repo
                    </a>

                    <div class="row justify-content-center mt-2">
                        <github-button href="https://github.com/statping/statping" data-icon="octicon-star" data-show-count="true" aria-label="Star Statping on GitHub">Star</github-button>
                    </div>

                </div>

            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <div class="card text-black-50 bg-white">
                            <div class="card-header">Statping Settings</div>
                            <div class="card-body">
                                <CoreSettings/>
                            </div>
                        </div>

                        <div class="card text-black-50 bg-white mt-3">
                            <div class="card-header">API Settings</div>
                            <div class="card-body">
                                <div class="form-group row">
                                    <label class="col-sm-3 col-form-label">API Secret</label>
                                    <div class="col-sm-9">
                                        <div class="input-group">
                                        <input v-model="core.api_secret" @focus="$event.target.select()" type="text" class="form-control select-input" id="api_secret" readonly>
                                            <div class="input-group-append copy-btn">
                                                <button @click="copy(core.api_secret)" class="btn btn-outline-secondary" type="button">Copy</button>
                                            </div>
                                        </div>
                                        <small class="form-text text-muted">API Secret is used for read, create, update and delete routes</small>
                                        <small class="form-text text-muted">You can <a href="#" id="regenkeys" @click="renewApiKeys">Regenerate API Keys</a> if you need to.</small>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="card text-black-50 bg-white mt-3">
                            <div class="card-header">QR Code for Mobile App</div>
                            <div class="card-body">

                                <img class="rounded" width="300" height="300" :src="qrcode">

                            </div>
                        </div>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">
                        <div class="card text-black-50 bg-white mb-5">
                            <div class="card-header">Theme Editor</div>
                            <div class="card-body">
                                <ThemeEditor/>
                            </div>
                        </div>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-cache-tab'), show: liClass('v-pills-cache-tab')}" id="v-pills-cache" role="tabpanel" aria-labelledby="v-pills-cache-tab">
                        <div class="card text-black-50 bg-white mb-5">
                            <div class="card-header">Cache</div>
                            <div class="card-body">
                                <Cache/>
                            </div>
                        </div>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-oauth-tab'), show: liClass('v-pills-oauth-tab')}" id="v-pills-oauth" role="tabpanel" aria-labelledby="v-pills-oauth-tab">
                        <OAuth/>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-notifier-docs-tab`), show: liClass(`v-pills-notifier-docs-tab`)}" v-bind:id="`v-pills-notifier-docs-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-notifier-docs-tab`">
                        <h2>Notifier Variables</h2>
                        You can insert dynamic fields within the notifier payloads for some notifiers.

                        <p class="mt-2">
                        Checkout the <a href="https://github.com/statping/statping/blob/master/types/services/struct.go">Service struct</a> and the <a href="https://github.com/statping/statping/blob/master/types/failures/struct.go">Failures struct</a> and create variables in golang template format.
                        </p>

                        <p class="mt-2">
                            For example, if you have <b>{{"\{\{.Service.Name\}\}"}}</b> it will return the service name.
                        </p>

                        <h2 class="mt-3">Service Variables</h2>

                        <table class="table">
                            <thead>
                            <tr>
                                <th scope="col">Variable</th>
                                <th scope="col">True Value</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td><kbd>{{"\{\{.Service.Id\}\}"}}</kbd></td>
                                <td>1</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.Name\}\}"}}</kbd></td>
                                <td>Example Service</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.Domain\}\}"}}</kbd></td>
                                <td>https://statping.com</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.Port\}\}"}}</kbd></td>
                                <td>8080</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.DowntimeAgo\}\}"}}</kbd></td>
                                <td>35 minutes ago</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.LastStatusCode\}\}"}}</kbd></td>
                                <td>404</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Service.FailuresLast24Hours\}\}"}}</kbd></td>
                                <td>38</td>
                            </tr>
                            </tbody>
                            <small>Additional variables within the Service struct</small>
                        </table>

                        <h2 class="mt-3">Failure Variables</h2>

                        <table class="table">
                            <thead>
                            <tr>
                                <th scope="col">Variable</th>
                                <th scope="col">True Value</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.Issue\}\}"}}</kbd></td>
                                <td>Received 404 status code</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.ErrorCode\}\}"}}</kbd></td>
                                <td>404</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.Service\}\}"}}</kbd></td>
                                <td>1</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.PingTime\}\}"}}</kbd></td>
                                <td>12482 (microseconds)</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.DowntimeAgo\}\}"}}</kbd></td>
                                <td>35 minutes ago</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Failure.CreatedAt\}\}"}}</kbd></td>
                                <td>2020-05-02 09:14:43.66381 +0000 UTC</td>
                            </tr>
                            </tbody>
                            <small>Additional variables within the Failures struct</small>
                        </table>

                        <h2 class="mt-3">Core Variables</h2>

                        <table class="table">
                            <thead>
                            <tr>
                                <th scope="col">Variable</th>
                                <th scope="col">True Value</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td><kbd>{{"\{\{.Core.Domain\}\}"}}</kbd></td>
                                <td>http://localhost:8080</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Core.Name\}\}"}}</kbd></td>
                                <td>Statping Demo</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Core.Description\}\}"}}</kbd></td>
                                <td>Statping will monitor your stuff!</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Core.Version\}\}"}}</kbd></td>
                                <td>v0.90.34</td>
                            </tr>
                            <tr>
                                <td><kbd>{{"\{\{.Core.Started\}\}"}}</kbd></td>
                                <td>2020-05-02 09:14:43.66381 +0000 UTC</td>
                            </tr>
                            </tbody>
                            <small>Additional variables within the Core struct</small>
                        </table>


                    </div>

                    <div v-for="(notifier, index) in notifiers" v-bind:key="`${notifier.method}_${index}`" class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), show: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-${notifier.method.toLowerCase()}-tab`">
                        <Notifier :notifier="notifier"/>
                    </div>

                </div>
            </div>

        </div>
    </div>
</template>

<script>
  import Api from '../API';
  import CoreSettings from '../forms/CoreSettings';
  import FormIntegration from '../forms/Integration';
  import Notifier from "../forms/Notifier";
  import ThemeEditor from "../components/Dashboard/ThemeEditor";
  import Cache from "@/components/Dashboard/Cache";
  import OAuth from "../forms/OAuth";
  import GithubButton from 'vue-github-button'

  export default {
      name: 'Settings',
      components: {
        GithubButton,
        OAuth,
          Cache,
          ThemeEditor,
          FormIntegration,
          Notifier,
          CoreSettings
      },
      data() {
          return {
              tab: "v-pills-home-tab",
              qrcode: "",
              qrurl: "",
          }
      },
      computed: {
          core() {
              return this.$store.getters.core
          },
          notifiers() {
              return this.$store.getters.notifiers
          }
      },
      async mounted() {
          this.cache = await Api.cache()
      },
      async created() {
          const c = this.core
          this.qrurl = `statping://setup?domain=${c.domain}&api=${c.api_secret}`
          this.qrcode = "https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + encodeURI(this.qrurl)
      },
      methods: {
          changeTab(e) {
              this.tab = e.target.id
          },
          liClass(id) {
              return this.tab === id
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
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
