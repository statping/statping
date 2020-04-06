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

                    <div id="notifiers_tabs">
                        <a v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="`${notifier.method}_${index}`" @click.prevent="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false">
                            <font-awesome-icon :icon="iconName(notifier.icon)" class="mr-2"/> {{notifier.method}}
                            <span v-if="notifier.enabled" class="badge badge-pill float-right mt-1" :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}">ON</span>
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

                </div>

            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <div class="card text-black-50 bg-white mb-5">
                            <div class="card-header">Statping Settings</div>
                            <div class="card-body">

                                <CoreSettings :core="core"/>

                                 <form name="updatePercentileRank" @submit.prevent="UpdatePercentileRank">
                                    <div class="form-group">
                                        <label for="percentileRank">Percentile of response time</label> <br>
                                        Current: <strong>{{this.$store.getters.percentileRank}}th percentile</strong> <br>
                                        <input type="text" v-model="percentileRank" id="percentileRank" class="form-control"  placeholder="Percentile Rank, e.g. 95">
                                        <br>
                                    </div>
                                    <button @click.prevent="UpdatePercentileRank" class="btn btn-primary btn-block">Save Percentile Rank</button>
                                </form>

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
              core: this.$store.getters.core,
              percentileRank: this.$store.getters.percentileRank,
          }
      },
      async mounted() {
          this.cache = await Api.cache()
      },
      async created() {
          const c = this.$store.state.core
          this.qrurl = `statping://setup?domain=${c.domain}&api=${c.api_secret}`
          this.qrcode = "https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + encodeURI(this.qrurl)
      },
      async beforeMount() {
        this.core = await Api.core()
      },
      methods: {
          changeTab(e) {
              this.tab = e.target.id
          },
          liClass(id) {
              return this.tab === id
          },
        async UpdatePercentileRank() {
            if (this.validatePercentileRankInput()) {
                const data = { rank: parseInt(this.percentileRank) }
                const response = await Api.percentile_rank_update(data)
                if (response.status != 'error') {
                    await this.$store.dispatch('loadRequired')
                    this.$router.push('/dashboard/settings')
                }   
                else {
                    alert("Percentile update failed. Error: " + response.error)
                }
            } 
        },
        validatePercentileRankInput() {
            var input = document.getElementById("percentileRank").value;
            if (input == "") {
                alert("Percentile field must be filled out");
                return false;
            }
            if (input < 1 || input > 100) {
                alert("Percentile number has to be between 1 and 100.")
                return false;
            }
            var onlyNumbers = new RegExp('^[0-9]+$');
            if (!onlyNumbers.test(input)) {
                alert("Only whole numbers are allowed")
                return false;
            }
            return true;
        },
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
