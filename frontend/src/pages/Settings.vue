<template>
    <div class="col-12">
        <div class="row">
            <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
                <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
                    <h6 class="text-muted">{{ $t('settings.main') }}</h6>

                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-home-tab')}" id="v-pills-home-tab" data-toggle="pill" href="#v-pills-home" role="tab" aria-controls="v-pills-home" aria-selected="true">
                        <font-awesome-icon icon="cog" class="mr-2"/> {{ $t('setting') }}
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-style-tab')}" id="v-pills-style-tab" data-toggle="pill" href="#v-pills-style" role="tab" aria-controls="v-pills-style" aria-selected="false">
                        <font-awesome-icon icon="image" class="mr-2"/> {{ $t('settings.theme') }}
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-cache-tab')}" id="v-pills-cache-tab" data-toggle="pill" href="#v-pills-cache" role="tab" aria-controls="v-pills-cache" aria-selected="false">
                        <font-awesome-icon icon="paperclip" class="mr-2"/> {{ $t('settings.cache') }}
                    </a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-oauth-tab')}" id="v-pills-oauth-tab" data-toggle="pill" href="#v-pills-oauth" role="tab" aria-controls="v-pills-oauth" aria-selected="false">
                        <font-awesome-icon icon="key" class="mr-2"/> {{ $t('settings.oauth') }}
                    </a>

                    <h6 class="mt-4 text-muted">Notifiers</h6>

                    <div id="notifiers_tabs">
                        <a v-for="(notifier, index) in notifiers" v-bind:key="`${notifier.method}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                            <font-awesome-icon :icon="iconName(notifier.icon)" class="mr-2"/> {{notifier.title}}
                            <span v-if="notifier.enabled" class="badge badge-pill float-right mt-1" :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}">ON</span>
                        </a>
                        <a @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-notifier-docs-tab`)}" v-bind:id="`v-pills-notifier-docs-tab`" data-toggle="pill" v-bind:href="`#v-pills-notifier-docs`" role="tab" v-bind:aria-controls="`v-pills-notifier-docs`" aria-selected="false">
                            <font-awesome-icon icon="question" class="mr-2"/> Variables
                        </a>
                    </div>

                    <h6 class="mt-4 mb-3 text-muted">Statping Links</h6>

                    <a href="https://github.com/statping/statping/wiki" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="question" class="mr-3"/> {{$t('settings.docs')}}
                    </a>

                    <a href="https://github.com/statping/statping/wiki/API" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="laptop" class="mr-2"/> API {{$t('settings.docs')}}
                    </a>

                    <a href="https://raw.githubusercontent.com/statping/statping/master/CHANGELOG.md" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="book" class="mr-3"/> {{$t('settings.changelog')}}
                    </a>

                    <a href="https://github.com/statping/statping" class="mb-2 font-2 text-decoration-none text-muted">
                        <font-awesome-icon icon="code-branch" class="mr-3"/> {{$t('settings.repo')}}
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

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">
                        <ThemeEditor/>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-cache-tab'), show: liClass('v-pills-cache-tab')}" id="v-pills-cache" role="tabpanel" aria-labelledby="v-pills-cache-tab">
                        <Cache/>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-oauth-tab'), show: liClass('v-pills-oauth-tab')}" id="v-pills-oauth" role="tabpanel" aria-labelledby="v-pills-oauth-tab">
                        <OAuth/>
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
  import GithubButton from 'vue-github-button'
  import Variables from "@/components/Dashboard/Variables";

  const CoreSettings = () => import('@/forms/CoreSettings')
  const FormIntegration = () => import('@/forms/Integration')
  const Notifier = () => import('@/forms/Notifier')
  const OAuth = () => import('@/forms/OAuth')
  const ThemeEditor = () => import('@/components/Dashboard/ThemeEditor')
  const Cache = () => import('@/components/Dashboard/Cache')

  export default {
      name: 'Settings',
      components: {
        Variables,
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
    mounted() {
        this.update()
      },
    created() {
          this.update()
      },
      methods: {
        async update() {
          const c = await Api.core()
          this.$store.commit('setCore', c)
          const n = await Api.notifiers()
          this.$store.commit('setNotifiers', n)
          this.cache = await Api.cache()
        },
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
