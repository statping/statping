<template>
    <div class="col-12">
        <div class="row">
            <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
                <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">

                    <div v-if="version_below" class="col-12 small text-center mt-0 pt-0 pb-0 mb-3">
                      Update {{github.tag_name}} Available
                      <div class="row">
                        <div class="col-6">
                          <a href="https://github.com/statping-ng/statping-ng/releases/latest" class="btn btn-sm text-success mt-2">Download</a>
                        </div>
                        <div class="col-6">
                          <a href="https://github.com/statping-ng/statping-ng/blob/master/CHANGELOG.md" class="btn btn-sm text-dim mt-2">Changelog</a>
                        </div>
                      </div>
                    </div>

                    <h6 class="text-muted">{{ $t('main_settings') }}</h6>

                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-home-tab')}" id="v-pills-home-tab" data-toggle="pill" href="#v-pills-home" role="tab" aria-controls="v-pills-home" aria-selected="true">
                        <font-awesome-icon icon="cog" class="mr-2"/> {{ $t('settings') }}
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-style-tab')}" id="v-pills-style-tab" data-toggle="pill" href="#v-pills-style" role="tab" aria-controls="v-pills-style" aria-selected="false">
                        <font-awesome-icon icon="image" class="mr-2"/> {{ $t('theme') }}
                    </a>
                  <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-oauth-tab')}" id="v-pills-oauth-tab" data-toggle="pill" href="#v-pills-oauth" role="tab" aria-controls="v-pills-oauth" aria-selected="false">
                    <font-awesome-icon icon="key" class="mr-2"/> {{ $t('authentication') }}
                  </a>
                  <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-import-tab')}" id="v-pills-import-tab" data-toggle="pill" href="#v-pills-import" role="tab" aria-controls="v-pills-import" aria-selected="false">
                    <font-awesome-icon icon="cloud-download-alt" class="mr-2"/> {{ $t('import') }}
                  </a>
                  <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-configs-tab')}" id="v-pills-configs-tab" data-toggle="pill" href="#v-pills-configs" role="tab" aria-controls="v-pills-configs" aria-selected="false">
                    <font-awesome-icon icon="cogs" class="mr-2"/> {{ $t('configs') }}
                  </a>

                    <h6 class="mt-4 text-muted">{{$t('notifiers')}}</h6>

                    <div id="notifiers_tabs">
                        <a v-for="(notifier, index) in notifiers" v-bind:key="`${notifier.method}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                            <font-awesome-icon :icon="iconName(notifier.icon)" class="mr-2"/> {{notifier.title}}
                            <span v-if="notifier.enabled" class="badge badge-pill float-right mt-1" :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}">ON</span>
                        </a>
                        <a @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-notifier-docs-tab`)}" v-bind:id="`v-pills-notifier-docs-tab`" data-toggle="pill" v-bind:href="`#v-pills-notifier-docs`" role="tab" v-bind:aria-controls="`v-pills-notifier-docs`" aria-selected="false">
                            <font-awesome-icon icon="question" class="mr-2"/> {{$t('variables')}}
                        </a>
                    </div>

                    <h6 class="mt-4 mb-3 text-muted">Statping-ng {{$t('links')}}</h6>

                    <a href="https://statping-ng.github.io" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="globe" class="mr-3"/> Statping-ng
                    </a>

                    <a href="https://github.com/statping-ng/statping-ng/wiki" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="question" class="mr-3"/> {{$t('docs')}}
                    </a>

                    <a href="https://github.com/statping-ng/statping-ng/wiki/API" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="laptop" class="mr-2"/> API {{$t('docs')}}
                    </a>

                    <a href="https://raw.githubusercontent.com/statping-ng/statping-ng/stable/CHANGELOG.md" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="book" class="mr-3"/> {{$t('changelog')}}
                    </a>

                    <a href="https://github.com/statping-ng/statping-ng" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="code-branch" class="mr-3"/> {{$t('repo')}}
                    </a>

                  <span class="small text-dim text-center mt-5">Statping-ng v{{core.version}}<br>
                    <a class="small text-muted no-decoration" v-if="core.commit" v-bind:href="`https://github.com/statping-ng/statping-ng/commit/${core.commit}`">{{core.commit.slice(0,8)}}</a>
                  </span>


                </div>

            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <CoreSettings/>

                        <div class="card mt-3">
                            <div class="card-header">API {{$t('settings')}}</div>
                            <div class="card-body">
                                <div class="form-group row">
                                    <label class="col-sm-3 col-form-label">API {{$t('secret')}}</label>
                                    <div class="col-sm-9">
                                        <div class="input-group">
                                        <input v-model="core.api_secret" @focus="$event.target.select()" type="text" class="form-control select-input" id="api_secret" readonly>
                                            <div class="input-group-append copy-btn">
                                                <button @click="copy(core.api_secret)" class="btn btn-outline-secondary" type="button">{{$t('copy')}}</button>
                                            </div>
                                        </div>
                                        <small class="form-text text-muted">{{$t('regen_desc')}}</small>

                                    </div>
                                </div>
                            </div>
                            <div class="card-footer">
                                <button id="regenkeys" @click="renewApiKeys" class="btn btn-sm btn-danger float-right">{{$t('regen_api')}}</button>
                            </div>
                        </div>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">
                        <ThemeEditor/>
                    </div>

                  <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-oauth-tab'), show: liClass('v-pills-oauth-tab')}" id="v-pills-oauth" role="tabpanel" aria-labelledby="v-pills-oauth-tab">
                    <OAuth/>
                  </div>

                  <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-configs-tab'), show: liClass('v-pills-configs-tab')}" id="v-pills-configs" role="tabpanel" aria-labelledby="v-pills-configs-tab">
                    <Configs/>
                  </div>

                  <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-import-tab'), show: liClass('v-pills-import-tab')}" id="v-pills-import" role="tabpanel" aria-labelledby="v-pills-import-tab">
                    <Importer/>
                  </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-notifier-docs-tab`), show: liClass(`v-pills-notifier-docs-tab`)}" v-bind:id="`v-pills-notifier-docs-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-notifier-docs-tab`">
                        <Variables/>
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
  const semver = require('semver')

  const CoreSettings = () => import(/* webpackChunkName: "dashboard" */ '@/forms/CoreSettings')
  const FormIntegration = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Integration')
  const Notifier = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Notifier')
  const OAuth = () => import(/* webpackChunkName: "dashboard" */ '@/forms/OAuth')
  const ThemeEditor = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/ThemeEditor')
  const Importer = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Importer')
  const Variables = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Variables')
  const Configs = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Configs')

  export default {
      name: 'Settings',
      components: {
        Configs,
        Importer,
        Variables,
        OAuth,
          ThemeEditor,
          FormIntegration,
          Notifier,
          CoreSettings
      },
      data() {
          return {
              tab: "v-pills-home-tab",
            github: null,
          }
      },
      computed: {
          core() {
              return this.$store.getters.core
          },
          notifiers() {
            return this.$store.getters.notifiers
          },
        version_below() {
            if (!this.github || !this.core.version) {
              return false
            }
            return semver.gt(semver.coerce(this.github.tag_name), semver.coerce(this.core.version))
        }
      },
    mounted() {

      },
    created() {
      this.update()
      },
      methods: {
        async update() {
          await this.getGithub()
        },
        async getGithub() {
          try {
            this.github = await Api.github_release()
          } catch(e) {
            console.error(e)
          }
        },
          changeTab(e) {
              this.tab = e.target.id
          },
          liClass(id) {
              return this.tab === id
          },
        async renew() {
          await Api.renewApiKeys()
          const core = await Api.core()
          this.$store.commit('setCore', core)
          this.core = core
          await this.logout()
        },
        async renewApiKeys() {
          const modal = {
            visible: true,
            title: "Reset API Key",
            body: `Are you sure you want to reset the API keys? You will be logged out.`,
            btnColor: "btn-danger",
            btnText: "Reset",
            func: () => this.renew(),
          }
          this.$store.commit("setModal", modal)
        },
        async logout () {
          await Api.logout()
          this.$store.commit('setHasAllData', false)
          this.$store.commit('setToken', null)
          this.$store.commit('setAdmin', false)
          this.$store.commit('setUser', false)
          // this.$cookies.remove("statping_auth")
          await this.$router.push('/logout')
        }
      }
  }
</script>
