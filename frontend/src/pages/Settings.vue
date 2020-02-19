<template>
    <div class="col-12">
        <div class="row">
            <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
                <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
                    <h6 class="text-muted">Main Settings</h6>

                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-home-tab')}" id="v-pills-home-tab" data-toggle="pill" href="#v-pills-home" role="tab" aria-controls="v-pills-home" aria-selected="true"><i class="fa fa-cogs"></i> Settings</a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-style-tab')}" id="v-pills-style-tab" data-toggle="pill" href="#v-pills-style" role="tab" aria-controls="v-pills-style" aria-selected="false"><i class="fa fa-image"></i> Theme Editor</a>
                    <a @click.prevent="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-cache-tab')}" id="v-pills-cache-tab" data-toggle="pill" href="#v-pills-cache" role="tab" aria-controls="v-pills-cache" aria-selected="false"><i class="fa fa-paperclip"></i> Cache</a>

                    <h6 class="mt-4 text-muted">Notifiers</h6>

                    <a v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="`${notifier.method}_${index}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                        <i class="fas fa-terminal"></i> {{notifier.method}}
                    </a>

                    <h6 class="mt-4 text-muted">Integrations (beta)</h6>

                    <a v-for="(integration, index) in $store.getters.integrations" v-bind:key="`${integration.name}_${index}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-integration-${integration.name}`)}" v-bind:id="`v-pills-integration-${integration.name}`" data-toggle="pill" v-bind:href="`#v-pills-integration-${integration.name}`" role="tab" :aria-controls="`v-pills-integration-${integration.name}`" aria-selected="false">
                        <i class="fas fa-file-csv"></i> {{integration.full_name}}
                    </a>

                </div>
            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <CoreSettings/>

                        <h2 class="mt-5">Additional Settings</h2>
                        <div v-if="core.domain !== ''" class="row">
                            <div class="col-12">
                                <div class="row align-content-center">
                                    <img class="rounded text-center" width="300" height="300" :src="qrcode">
                                </div>
                                <a class="btn btn-sm btn-primary" href=statping://setup?domain&#61;https://demo.statping.com&amp;api&#61;6b05b48f4b3a1460f3864c31b26cab6a27dbaff9>Open in Statping App</a>
                                <a href="settings/export" class="btn btn-sm btn-secondary">Export Settings</a>
                            </div>
                        </div>
                        <div v-else>
                            Insert a domain to view QR code for the mobile app.
                        </div>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">

                        <ThemeEditor :core="core"/>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-cache-tab'), show: liClass('v-pills-cache-tab')}" id="v-pills-cache" role="tabpanel" aria-labelledby="v-pills-cache-tab">
                        <h3>Cache</h3>
                        <table class="table">
                            <thead>
                            <tr>
                                <th scope="col">URL</th>
                                <th scope="col">Size</th>
                                <th scope="col">Expiration</th>
                            </tr>
                            </thead>
                            <tbody>

                            <tr v-for="(cache, index) in cache">
                                <td>{{cache.url}}</td>
                                <td>{{cache.size}}</td>
                                <td>{{expireTime(cache.expiration)}}</td>
                            </tr>

                            </tbody>
                        </table>
                        <a @click.prevent="clearCache" href="#" class="btn btn-danger btn-block">Clear Cache</a>
                    </div>

                    <div v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="`${notifier.title}_${index}`" class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), show: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-${notifier.method.toLowerCase()}-tab`">
                        <Notifier :notifier="notifier"/>
                    </div>

                    <div v-for="(integration, index) in $store.getters.integrations" v-bind:key="`${integration.name}_${index}`" class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-integration-${integration.name}`), show: liClass(`v-pills-integration-${integration.name}`)}" v-bind:id="`v-pills-integration-${integration.name}`" role="tabpanel">
                        <FormIntegration :integration="integration"/>
                    </div>

                </div>
            </div>

        </div>
    </div>
</template>

<script>
  import Api from '../components/API';
  import CoreSettings from '../forms/CoreSettings';
  import FormIntegration from '../forms/Integration';
  import Notifier from "../forms/Notifier";
  import ThemeEditor from "../components/Dashboard/ThemeEditor";

  export default {
  name: 'Settings',
  components: {
    ThemeEditor,
      FormIntegration,
    Notifier,
    CoreSettings
  },
  data () {
    return {
      tab: "v-pills-home-tab",
      qrcode: "",
      core: this.$store.getters.core,
    cache: [],
    }
  },
      async mounted () {
          this.cache = await Api.cache()
      },
      async created() {
    const qrurl = `statping://setup?domain=${core.domain}&api=${core.api_secret}`
    this.qrcode = "https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + encodeURI(qrurl)
  },
  beforeMount() {

  },
  methods: {
    changeTab (e) {
      this.tab = e.target.id
    },
    liClass (id) {
      return this.tab === id
    },
    expireTime(ex) {
      return this.toLocal(ex)
    },
      async clearCache () {
          await Api.clearCache()
          this.cache = await Api.cache()
     }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
