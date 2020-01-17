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
                                    <label for="update_notify" class="d-inline d-sm-none">Send Updates only</label>
                                    <label for="update_notify" class="d-none d-sm-block">Send Updates only</label>

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

                        <form class="ajax_form command">
                            <h4 class="text-capitalize">{{notifier.title}}</h4>
                            <p class="small text-muted" v-html="notifier.description"></p>

                            <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
                                <label class="text-capitalize">{{form.title}}</label>

                                <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[notifier.field]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

                                <small class="form-text text-muted" v-html="form.small_text"></small>

                            </div>

                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_command" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
                                    <span class="switch">
                                        <input type="checkbox" name="enabled-option" class="switch" id="switch-Command" >
                                        <label v-bind:for="`switch-${notifier.method}`"></label>
                                        <input type="hidden" name="enabled" v-bind:id="`switch-${notifier.method}`" value="false">
                                    </span>
                                </div>

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>

                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="command-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> {{notifier.method}} has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="command-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The {{notifier.method}} notifier is working correctly!
                                    </div>
                                </div>
                            </div>

                            <span class="d-block small text-center mt-3 mb-5">
                                <span class="text-capitalize">{{notifier.title}}</span> Notifier created by <a :href="notifier.author_url" target="_blank">{{notifier.author}}</a>
                            </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-discord" role="tabpanel" aria-labelledby="v-pills-discord-tab">



                        <form class="ajax_form discord" data-func="SaveNotifier" action="api/notifier/discord" method="POST">
                            <h4 class="text-capitalize">discord</h4>
                            <p class="small text-muted">Send notifications to your discord channel using discord webhooks. Insert your discord channel Webhook URL to receive notifications. Based on the <a href="https://discordapp.com/developers/docs/resources/Webhook">discord webhooker API</a>.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="discord_webhooker_url">discord webhooker URL</label>

                                <input type="text" name="host" class="form-control" value="https://discordapp.com/api/webhooks/****/*****" id="discord_webhooker_url" placeholder="Insert your Webhook URL here" >


                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_discord" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-discord" >
                <label for="switch-discord"></label>
                <input type="hidden" name="enabled" id="switch-discord-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="discord">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="discord-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> discord has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="discord-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The discord notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">discord</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-email" role="tabpanel" aria-labelledby="v-pills-email-tab">



                        <form class="ajax_form email" data-func="SaveNotifier" action="api/notifier/email" method="POST">
                            <h4 class="text-capitalize">email</h4>
                            <p class="small text-muted">Send emails via SMTP when services are online or offline.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="smtp_host">SMTP Host</label>

                                <input type="text" name="host" class="form-control" value="" id="smtp_host" placeholder="Insert your SMTP Host here." >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="smtp_username">SMTP Username</label>

                                <input type="text" name="username" class="form-control" value="" id="smtp_username" placeholder="Insert your SMTP Username here." >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="smtp_password">SMTP Password</label>

                                <input type="password" name="password" class="form-control" value="" id="smtp_password" placeholder="Insert your SMTP Password here." >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="smtp_port">SMTP Port</label>

                                <input type="number" name="port" class="form-control" value="0" id="smtp_port" placeholder="Insert your SMTP Port here." >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="outgoing_email_address">Outgoing Email Address</label>

                                <input type="text" name="var1" class="form-control" value="" id="outgoing_email_address" placeholder="outgoing@email.com" >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="send_alerts_to">Send Alerts To</label>

                                <input type="email" name="var2" class="form-control" value="" id="send_alerts_to" placeholder="sendto@email.com" >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="disable_tls_ssl">Disable TLS/SSL</label>

                                <input type="text" name="api_key" class="form-control" value="" id="disable_tls_ssl" placeholder="" >


                                <small class="form-text text-muted">To Disable TLS/SSL insert 'true'</small>

                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_email" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-email" >
                <label for="switch-email"></label>
                <input type="hidden" name="enabled" id="switch-email-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="email">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="email-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> email has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="email-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The email notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">email</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-line_notify" role="tabpanel" aria-labelledby="v-pills-line_notify-tab">



                        <form class="ajax_form line_notify" data-func="SaveNotifier" action="api/notifier/line%20notify" method="POST">
                            <h4 class="text-capitalize">LINE Notify</h4>
                            <p class="small text-muted">LINE Notify will send notifications to your LINE Notify account when services are offline or online. Based on the <a href="https://notify-bot.line.me/doc/en/">LINE Notify API</a>.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="access_token">Access Token</label>

                                <input type="text" name="api_secret" class="form-control" value="" id="access_token" placeholder="Insert your Line Notify Access Token here." >


                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_line_notify" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-line notify" >
                <label for="switch-line notify"></label>
                <input type="hidden" name="enabled" id="switch-line notify-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="line_notify">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>



                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">LINE Notify</span> Notifier created by <a href="https://github.com/dogrocker" target="_blank">Kanin Peanviriyakulkit</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-mobile" role="tabpanel" aria-labelledby="v-pills-mobile-tab">



                        <form class="ajax_form mobile" data-func="SaveNotifier" action="api/notifier/mobile" method="POST">
                            <h4 class="text-capitalize">Mobile Notifications</h4>
                            <p class="small text-muted">Receive push notifications on your Mobile device using the Statping App. You can scan the Authentication QR Code found in Settings to get the Mobile app setup in seconds.
                            <p align="center"><a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a><a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a></p>


                            <div class="form-group">
                                <label class="text-capitalize d-none" for="device_identifiers">Device Identifiers</label>

                                <input type="text" name="var1" class="form-control d-none" value="" id="device_identifiers" placeholder="A list of your Mobile device push notification ID&#39;s." >


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize d-none" for="array_of_device_numbers">Array of device numbers</label>

                                <input type="number" name="var2" class="form-control d-none" value="" id="array_of_device_numbers" placeholder="1 for iphone 2 for android" >


                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_mobile" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-mobile" >
                <label for="switch-mobile"></label>
                <input type="hidden" name="enabled" id="switch-mobile-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="mobile">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="mobile-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> mobile has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="mobile-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The mobile notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">Mobile Notifications</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-slack" role="tabpanel" aria-labelledby="v-pills-slack-tab">



                        <form class="ajax_form slack" data-func="SaveNotifier" action="api/notifier/slack" method="POST">
                            <h4 class="text-capitalize">slack</h4>
                            <p class="small text-muted">Send notifications to your slack channel when a service is offline. Insert your Incoming webhooker URL for your channel to receive notifications. Based on the <a href="https://api.slack.com/incoming-webhooks">slack API</a>.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="incoming_webhooker_url">Incoming webhooker Url</label>

                                <input type="text" name="host" class="form-control" value="https://webhooksurl.slack.com/***" id="incoming_webhooker_url" placeholder="Insert your slack Webhook URL here." required>


                                <small class="form-text text-muted">Incoming webhooker URL from <a href="https://api.slack.com/apps" target="_blank">slack Apps</a></small>

                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_slack" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-slack" >
                <label for="switch-slack"></label>
                <input type="hidden" name="enabled" id="switch-slack-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="slack">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="slack-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> slack has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="slack-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The slack notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">slack</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-telegram" role="tabpanel" aria-labelledby="v-pills-telegram-tab">



                        <form class="ajax_form telegram" data-func="SaveNotifier" action="api/notifier/telegram" method="POST">
                            <h4 class="text-capitalize">Telegram</h4>
                            <p class="small text-muted">Receive notifications on your Telegram channel when a service has an issue. You must get a Telegram API token from the /botfather. Review the <a target="_blank" href="http://techthoughts.info/how-to-create-a-telegram-bot-and-send-messages-via-api">Telegram API Tutorial</a> to learn how to generate a new API Token.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="telegram_api_token">Telegram API Token</label>

                                <input type="text" name="api_secret" class="form-control" value="" id="telegram_api_token" placeholder="383810182:EEx829dtCeufeQYXG7CUdiQopqdmmxBPO7-s" required>


                                <small class="form-text text-muted">Enter the API Token given to you from the /botfather chat.</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="channel_or_user_id">Channel or User ID</label>

                                <input type="text" name="var1" class="form-control" value="" id="channel_or_user_id" placeholder="789325392" required>


                                <small class="form-text text-muted">Insert your Telegram Channel ID or User ID here.</small>

                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_telegram" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-telegram" >
                <label for="switch-telegram"></label>
                <input type="hidden" name="enabled" id="switch-telegram-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="telegram">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="telegram-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> telegram has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="telegram-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The telegram notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">Telegram</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-twilio" role="tabpanel" aria-labelledby="v-pills-twilio-tab">



                        <form class="ajax_form twilio" data-func="SaveNotifier" action="api/notifier/twilio" method="POST">
                            <h4 class="text-capitalize">Twilio</h4>
                            <p class="small text-muted">Receive SMS text messages directly to your cellphone when a service is offline. You can use a Twilio test account with limits. This notifier uses the <a href="https://www.twilio.com/docs/usage/api">Twilio API</a>.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="account_sid">Account SID</label>

                                <input type="text" name="api_key" class="form-control" value="" id="account_sid" placeholder="Insert your Twilio Account SID" required>


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="account_token">Account Token</label>

                                <input type="text" name="api_secret" class="form-control" value="" id="account_token" placeholder="Insert your Twilio Account Token" required>


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="sms_to_phone_number">SMS to Phone Number</label>

                                <input type="text" name="var1" class="form-control" value="" id="sms_to_phone_number" placeholder="18555555555" required>


                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="from_phone_number">From Phone Number</label>

                                <input type="text" name="var2" class="form-control" value="" id="from_phone_number" placeholder="18555555555" required>


                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_twilio" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-twilio" >
                <label for="switch-twilio"></label>
                <input type="hidden" name="enabled" id="switch-twilio-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="twilio">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="twilio-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> twilio has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="twilio-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The twilio notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">Twilio</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



                    </div>


                    <div class="tab-pane fade" id="v-pills-webhook" role="tabpanel" aria-labelledby="v-pills-webhook-tab">



                        <form class="ajax_form webhook" data-func="SaveNotifier" action="api/notifier/Webhook" method="POST">
                            <h4 class="text-capitalize">HTTP webhooker</h4>
                            <p class="small text-muted">Send a custom HTTP request to a specific URL with your own body, headers, and parameters.</p>


                            <div class="form-group">
                                <label class="text-capitalize" for="http_endpoint">HTTP Endpoint</label>

                                <input type="text" name="host" class="form-control" value="" id="http_endpoint" placeholder="http://webhookurl.com/JW2MCP4SKQP" required>


                                <small class="form-text text-muted">Insert the URL for your HTTP Requests.</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="http_method">HTTP Method</label>

                                <input type="text" name="var1" class="form-control" value="" id="http_method" placeholder="POST" required>


                                <small class="form-text text-muted">Choose a HTTP method for example: GET, POST, DELETE, or PATCH.</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="http_body">HTTP Body</label>

                                <textarea rows="3" class="form-control" name="var2" id="http_body"></textarea>


                                <small class="form-text text-muted">Optional HTTP body for a POST request. You can insert variables into your body request.<br>%service.Id, %service.Name, %service.Online<br>%failure.Issue</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="content_type">Content Type</label>

                                <input type="text" name="api_key" class="form-control" value="" id="content_type" placeholder="application/json" >


                                <small class="form-text text-muted">Optional content type for example: application/json or text/plain</small>

                            </div>

                            <div class="form-group">
                                <label class="text-capitalize" for="header">Header</label>

                                <input type="text" name="api_secret" class="form-control" value="" id="header" placeholder="Authorization=Token12345" >


                                <small class="form-text text-muted">Optional Headers for request use format: KEY=Value,Key=Value</small>

                            </div>


                            <div class="row">
                                <div class="col-9 col-sm-6">
                                    <div class="input-group mb-2">
                                        <div class="input-group-prepend">
                                            <div class="input-group-text">Limit</div>
                                        </div>
                                        <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_webhook" value="3" placeholder="7">
                                        <div class="input-group-append">
                                            <div class="input-group-text">Per Minute</div>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-3 col-sm-2 mt-1">
            <span class="switch">
                <input type="checkbox" name="enabled-option" class="switch" id="switch-Webhook" >
                <label for="switch-Webhook"></label>
                <input type="hidden" name="enabled" id="switch-Webhook-value" value="false">
            </span>
                                </div>

                                <input type="hidden" name="method" value="webhook">

                                <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                                    <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
                                </div>


                                <div class="col-12 col-sm-12">
                                    <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
                                </div>

                                <div class="col-12 col-sm-12 mt-2">
                                    <div class="alert alert-danger d-none" id="webhook-error" role="alert">
                                        <i class="fa fa-exclamation-triangle"></i> Webhook has an error!
                                    </div>

                                    <div class="alert alert-success d-none" id="webhook-success" role="alert">
                                        <i class="fa fa-smile-beam"></i> The Webhook notifier is working correctly!
                                    </div>
                                </div>


                            </div>


                            <span class="d-block small text-center mt-3 mb-5">
        <span class="text-capitalize">HTTP webhooker</span> Notifier created by <a href="https://github.com/hunterlong" target="_blank">Hunter Long</a>
                        </span>

                            <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                        </form>



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
  import Api from "../components/API"
  import CoreSettings from '../forms/CoreSettings';

  export default {
  name: 'Settings',
  components: {
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
