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

                    <h6 class="mt-4 text-muted">Notifiers</h6>

                    <a v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="`${notifier.method}_${index}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                        <font-awesome-icon :icon="iconName(notifier.icon)" class="mr-2"/> {{notifier.method}}
                        <span v-if="notifier.enabled" class="badge badge-pill float-right mt-1" :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}">ON</span>
                    </a>

                </div>
            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <div class="card text-black-50 bg-white mb-5">
                            <div class="card-header">Statping Settings</div>
                            <div class="card-body">

                                <CoreSettings/>

                            </div>
                        </div>


                        <div class="card text-black-50 bg-white mb-3">
                            <div class="card-header">Statping Settings</div>
                            <div class="card-body">

                        <h2 class="mt-5">Additional Settings</h2>
                        <div v-if="core.domain !== ''" class="row">
                            <div class="col-12">
                                <div class="row align-content-center">
                                    <img class="rounded text-center" width="300" height="300" :src="qrcode">
                                </div>
                                <a class="btn btn-sm btn-primary" :href="qrurl">Open in Statping App</a>
                                <a href="settings/export" class="btn btn-sm btn-secondary">Export Settings</a>
                            </div>
                        </div>
                        <div v-else>
                            Insert a domain to view QR code for the mobile app.
                        </div>

                            </div>
                        </div>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">
                        <div class="card text-black-50 bg-white mb-5">
                            <div class="card-header">Theme Editor</div>
                            <div class="card-body">
                                <ThemeEditor :core="core"/>
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

                    <div v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="`${notifier.title}_${index}`" class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), show: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-${notifier.method.toLowerCase()}-tab`">
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

  export default {
      name: 'Settings',
      components: {
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
              core: this.$store.getters.core
          }
      },
      async mounted() {
          this.cache = await Api.cache()
      },
      async created() {
          const c = this.$store.getters.core
          this.qrurl = `statping://setup?domain=${c.domain}&api=${c.api_secret}`
          this.qrcode = "https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + encodeURI(this.qrurl)
      },
      beforeMount() {

      },
      methods: {
          changeTab(e) {
              this.tab = e.target.id
          },
          liClass(id) {
              return this.tab === id
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
