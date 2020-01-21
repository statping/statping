<template>
    <div class="col-12">
        <div class="row">
            <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
                <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
                    <h6 class="text-muted">Main Settings</h6>

                    <a v-on:click="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-home-tab')}" id="v-pills-home-tab" data-toggle="pill" href="#v-pills-home" role="tab" aria-controls="v-pills-home" aria-selected="true"><i class="fa fa-cogs"></i> Settings</a>
                    <a v-on:click="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-notifications-tab')}" id="v-pills-notifications-tab" data-toggle="pill" href="#v-pills-notifications" role="tab" aria-controls="v-pills-notifications" aria-selected="true"><i class="fa fa-bell"></i> Notifications</a>
                    <a v-on:click="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-style-tab')}" id="v-pills-style-tab" data-toggle="pill" href="#v-pills-style" role="tab" aria-controls="v-pills-style" aria-selected="false"><i class="fa fa-image"></i> Theme Editor</a>
                    <a v-on:click="changeTab" class="nav-link" v-bind:class="{active: liClass('v-pills-cache-tab')}" id="v-pills-cache-tab" data-toggle="pill" href="#v-pills-cache" role="tab" aria-controls="v-pills-cache" aria-selected="false"><i class="fa fa-paperclip"></i> Cache</a>

                    <h6 class="mt-4 text-muted">Notifiers</h6>

                    <a v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="index" v-on:click="changeTab" class="nav-link text-capitalize" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" data-toggle="pill" v-bind:href="`#v-pills-${notifier.method.toLowerCase()}`" role="tab" v-bind:aria-controls="`v-pills-${notifier.method.toLowerCase()}`" aria-selected="false"><i class="fas fa-terminal"></i> {{notifier.method}}</a>

                    <h6 class="mt-4 text-muted">Integrations (beta)</h6>

                    <a class="nav-link text-capitalize" id="v-pills-integration-csv-tab" data-toggle="pill" href="#v-pills-integration-csv" role="tab" aria-controls="v-pills-integration-csv" aria-selected="false"><i class="fas fa-file-csv"></i> CSV File</a>

                    <a class="nav-link text-capitalize" id="v-pills-integration-docker-tab" data-toggle="pill" href="#v-pills-integration-docker" role="tab" aria-controls="v-pills-integration-docker" aria-selected="false"><i class="fab fa-docker"></i> Docker</a>

                    <a class="nav-link text-capitalize" id="v-pills-integration-traefik-tab" data-toggle="pill" href="#v-pills-integration-traefik" role="tab" aria-controls="v-pills-integration-traefik" aria-selected="false"><i class="fas fa-network-wired"></i> Traefik</a>

                </div>
            </div>
            <div class="col-md-9 col-sm-12">

                <div class="tab-content" id="v-pills-tabContent">
                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}" id="v-pills-home" role="tabpanel" aria-labelledby="v-pills-home-tab">

                        <CoreSettings/>

                        <h2 class="mt-5">Bulk Import Services</h2>
                        You can import multiple services based on a CSV file with the format shown on the <a href="https://github.com/hunterlong/statping/wiki/Bulk-Import-Services" target="_blank">Bulk Import Wiki</a>.

                        <div class="card mt-2">
                            <div class="card-body">
                                <form action="settings/bulk_import" method="POST" enctype="multipart/form-data" class="form-inline">
                                    <div class="form-group col-10">
                                        <input type="file" name="file" class="form-control-file" accept=".csv">
                                    </div>
                                    <div class="form-group">
                                        <button type="submit" class="btn btn-outline-success right">Import</button>
                                    </div>
                                </form>
                            </div>
                        </div>


                        <h2 class="mt-5">Additional Settings</h2>
                        <div class="row">
                            <div class="col-12">
                                <div class="row align-content-center">
                                    <img class="rounded text-center" width="300" height="300" src="https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=statping%3a%2f%2fsetup%3fdomain%3dhttps%3a%2f%2fdemo.statping.com%26api%3d6b05b48f4b3a1460f3864c31b26cab6a27dbaff9">
                                </div>
                                <a class="btn btn-sm btn-primary" href=statping://setup?domain&#61;https://demo.statping.com&amp;api&#61;6b05b48f4b3a1460f3864c31b26cab6a27dbaff9>Open in Statping App</a>
                                <a href="settings/export" class="btn btn-sm btn-secondary">Export Settings</a>
                            </div>
                        </div>

                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-notifications-tab'), show: liClass('v-pills-notifications-tab')}" id="v-pills-notifications" role="tabpanel" aria-labelledby="v-pills-notifications-tab">
                        <h3>Notifications</h3>

                        <form method="POST" action="settings">

                            <div class="form-group">
                                <div class="col-12">
                                    <label class="d-inline d-sm-none">Send Updates only</label>
                                    <label class="d-none d-sm-block">Send Updates only</label>

                                    <span class="switch">
                                        <input type="checkbox" name="update_notify-option" class="switch" id="switch-update_notify">
                                        <label for="switch-update_notify" class="mt-2 mt-sm-0"></label>
                                        <small class="form-text text-muted">Enabling this will send only notifications when the status of a services changes.</small>
                                    </span>

                                    <input type="hidden" name="update_notify" id="switch-update_notify-value" value="false">
                                </div>
                            </div>
                            <button type="submit" class="btn btn-primary btn-block">Save Settings</button>

                        </form>
                    </div>

                    <div class="tab-pane fade" v-bind:class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}" id="v-pills-style" role="tabpanel" aria-labelledby="v-pills-style-tab">


                        <form method="POST" action="settings/css">
                            <ul class="nav nav-pills mb-3" id="pills-tab" role="tablist">
                                <li class="nav-item col text-center">
                                    <a class="nav-link active" id="pills-vars-tab" data-toggle="pill" href="#pills-vars" role="tab" aria-controls="pills-vars" aria-selected="true">Variables</a>
                                </li>
                                <li class="nav-item col text-center">
                                    <a class="nav-link" id="pills-theme-tab" data-toggle="pill" href="#pills-theme" role="tab" aria-controls="pills-theme" aria-selected="false">Base Theme</a>
                                </li>
                                <li class="nav-item col text-center">
                                    <a class="nav-link" id="pills-mobile-tab" data-toggle="pill" href="#pills-mobile" role="tab" aria-controls="pills-mobile" aria-selected="false">Mobile</a>
                                </li>
                            </ul>
                            <div class="tab-content" id="pills-tabContent">
                                <div class="tab-pane show active" id="pills-vars" role="tabpanel" aria-labelledby="pills-vars-tab">
                                    <textarea name="variables" id="sass_vars">
                                        </textarea>
                                </div>
                            </div>
                            <button type="submit" class="btn btn-primary btn-block mt-2">Save Style</button>
                            <a href="settings/delete_assets" class="btn btn-danger btn-block confirm-btn">Delete All Assets</a>
                        </form>

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

                            <tr>
                                <td>/api/services/7/data?start=1577937580&amp;end=9999999999&amp;group=hour</td>
                                <td>13951</td>
                                <td>2020-01-15 20:00:10 -0800 -0800</td>
                            </tr>

                            </tbody>
                        </table>
                        <a href="api/clear_cache" class="btn btn-danger btn-block">Clear Cache</a>
                    </div>

                    <div v-for="(notifier, index) in $store.getters.notifiers" v-bind:key="index" class="tab-pane fade" v-bind:class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), show: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}" v-bind:id="`v-pills-${notifier.method.toLowerCase()}-tab`" role="tabpanel" v-bind:aria-labelledby="`v-pills-${notifier.method.toLowerCase()}-tab`">

                        <Notifier :notifier="notifier"/>
                    </div>



                    <div class="tab-pane fade" id="v-pills-integration-csv" role="tabpanel" aria-labelledby="v-pills-integration-csv-tab">


                        <form class="integration_csv" action="settings/integrator/csv" method="POST">
                            <input type="hidden" name="integrator" class="form-control" value="csv">
                            <h4 class="text-capitalize">csv</h4>
                            <p class="small text-muted">Import multiple services from a CSV file. Please have your CSV file formatted with the correct amount of columns based on the <a href="https://raw.githubusercontent.com/hunterlong/statping/master/source/tmpl/bulk_import.csv">example file on Github</a>.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="input">input</label>

                                <textarea rows="3" class="form-control" name="input" id="input"></textarea>


                            </div>


                            <button type="submit" class="btn btn-block btn-info fetch_integrator">Fetch Services</button>

                            <div class="alert alert-danger d-none" id="integration_alerter" role="alert"></div>
                        </form>

                    </div>


                    <div class="tab-pane fade" id="v-pills-integration-docker" role="tabpanel" aria-labelledby="v-pills-integration-docker-tab">


                        <form class="integration_docker" action="settings/integrator/docker" method="POST">
                            <input type="hidden" name="integrator" class="form-control" value="docker">
                            <h4 class="text-capitalize">docker</h4>
                            <p class="small text-muted">Import multiple services from Docker by attaching the unix socket to Statping.
                                You can also do this in Docker by setting <u>-v /var/run/docker.sock:/var/run/docker.sock</u> in the Statping Docker container.
                                All of the containers with open TCP/UDP ports will be listed for you to choose which services you want to add. If you running Statping inside of a container,
                                this container must be attached to all networks you want to communicate with.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="path">path</label>

                                <input type="text" name="path" class="form-control" value="unix:///var/run/docker.sock" id="path">


                                <small class="form-text text-muted">The absolute path to the Docker unix socket</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="version">version</label>

                                <input type="text" name="version" class="form-control" value="1.25" id="version">


                                <small class="form-text text-muted">Version number of Docker server</small>

                            </div>


                            <button type="submit" class="btn btn-block btn-info fetch_integrator">Fetch Services</button>

                            <div class="alert alert-danger d-none" id="integration_alerter" role="alert"></div>
                        </form>

                    </div>


                    <div class="tab-pane fade" id="v-pills-integration-traefik" role="tabpanel" aria-labelledby="v-pills-integration-traefik-tab">


                        <form class="integration_traefik" action="settings/integrator/traefik" method="POST">
                            <input type="hidden" name="integrator" class="form-control" value="traefik">
                            <h4 class="text-capitalize">traefik</h4>



                            <div class="form-group">
                                <label class="text-capitalize" for="endpoint">endpoint</label>

                                <input type="text" name="endpoint" class="form-control" value="http://localhost:8080" id="endpoint">


                                <small class="form-text text-muted">The URL for the traefik API Endpoint</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="username">username</label>

                                <input type="text" name="username" class="form-control" value="" id="username">


                                <small class="form-text text-muted">Username for HTTP Basic Authentication</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="password">password</label>

                                <input type="password" name="password" class="form-control" value="" id="password">


                                <small class="form-text text-muted">Password for HTTP Basic Authentication</small>

                            </div>


                            <button type="submit" class="btn btn-block btn-info fetch_integrator">Fetch Services</button>

                            <div class="alert alert-danger d-none" id="integration_alerter" role="alert"></div>
                        </form>

                    </div>



                    <div class="tab-pane fade" id="v-pills-browse" role="tabpanel" aria-labelledby="v-pills-browse-tab">

                    </div>


                    <div class="tab-pane fade" id="v-pills-backups" role="tabpanel" aria-labelledby="v-pills-backups-tab">
                        <a href="backups/create" class="btn btn-primary btn-block">Backup Database</a>
                    </div>



                </div>
            </div>

        </div>
    </div>
</template>

<script>
  import CoreSettings from '../forms/CoreSettings';
  import Notifier from "../forms/Notifier";

  export default {
  name: 'Settings',
  components: {
    Notifier,
      CoreSettings

  },
  data () {
    return {
      tab: "v-pills-home-tab",
    }
  },
  created() {

  },
  beforeMount() {

  },
  methods: {
    changeTab (e) {
      this.tab = e.target.id
    },
    liClass (id) {
      return this.tab === id
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
