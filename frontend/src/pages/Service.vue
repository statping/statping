<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <TopNav/>

        <div class="col-12 mb-4">

            <span class="mt-3 mb-3 text-white d-md-none btn bg-success d-block d-md-none">ONLINE</span>


            <h4 class="mt-2"><a href="">{{service.name}}</a> - {{service.name}}

                <span class="badge bg-success float-right d-none d-md-block">ONLINE</span>
            </h4>

            <div class="row stats_area mt-5 mb-5">
                <div class="col-4">
                    <span class="lg_number">100%</span>
                    Online last 24 Hours
                </div>
                <div class="col-4">
                    <span class="lg_number">31ms</span>
                    Average Response
                </div>
                <div class="col-4">
                    <span class="lg_number">85.70%</span>
                    Total Uptime
                </div>
            </div>



            <div class="service-chart-container">
                <div id="service"></div>
                <div id="service-bar"></div>
            </div>

            <div class="service-chart-heatmap">
                <div id="service_heatmap"></div>
            </div>

            <form id="service_date_form" class="col-12 mt-2 mb-3">
                <input type="text" class="d-none" name="start" id="service_start" data-input>
                <span data-toggle title="toggle" id="start_date" class="text-muted small float-left pointer mt-2">Thu, 09 Jan 2020 to Thu, 16 Jan 2020</span>
                <button type="submit" class="btn btn-light btn-sm mt-2">Set Timeframe</button>
                <input type="text" class="d-none" name="end" id="service_end" data-input>

                <div id="start_container"></div>
                <div id="end_container"></div>
            </form>


            <nav class="nav nav-pills flex-column flex-sm-row mt-3" id="service_tabs" role="serviceLists">
                <a class="flex-sm-fill text-sm-center nav-link active" id="edit-tab" data-toggle="tab" href="#edit" role="tab" aria-controls="edit" aria-selected="false">Edit Service</a>
                <a class="flex-sm-fill text-sm-center nav-link" id="failures-tab" data-toggle="tab" href="#failures" role="tab" aria-controls="failures" aria-selected="true">Failures</a>
                <a class="flex-sm-fill text-sm-center nav-link disabled" id="incidents-tab" data-toggle="tab" href="#incidents" role="tab" aria-controls="incidents" aria-selected="true">Incidents</a>
                <a class="flex-sm-fill text-sm-center nav-link" id="checkins-tab" data-toggle="tab" href="#checkins" role="tab" aria-controls="checkins" aria-selected="false">Checkins</a>
                <a class="flex-sm-fill text-sm-center nav-link" id="response-tab" data-toggle="tab" href="#response" role="tab" aria-controls="response" aria-selected="false">Response</a>
            </nav>
            <div class="tab-content" id="myTabContent">

                <div class="tab-pane fade" id="failures" role="serviceLists" aria-labelledby="failures-tab">

                    <div class="list-group mt-3 mb-4">

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Incorrect HTTP Status Code</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Status Code 502 did not match 200</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: dial tcp 162.248.92.36:443: connect: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: dial tcp 162.248.92.36:443: connect: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: dial tcp 162.248.92.36:443: connect: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: dial tcp 162.248.92.36:443: connect: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: dial tcp 162.248.92.36:443: connect: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Reset</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: read tcp 172.27.0.8:46586-&gt;162.248.92.36:443: read: connection reset by peer</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:50890-&gt;127.0.0.11:53: read: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:35222-&gt;127.0.0.11:53: read: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:49817-&gt;127.0.0.11:53: read: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Failed</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:57247-&gt;127.0.0.11:53: read: connection refused</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Reset</h5>
                                <small>4 weeks ago</small>
                            </div>
                            <p class="mb-1">HTTP Error Get https://api.statping.com: read tcp 172.27.0.10:52594-&gt;162.248.92.36:443: read: connection reset by peer</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Timed Out</h5>
                                <small>Last month</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:51238-&gt;127.0.0.11:53: i/o timeout</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Connection Timed Out</h5>
                                <small>Last month</small>
                            </div>
                            <p class="mb-1">Could not get IP address for domain https://api.statping.com, lookup api.statping.com on 127.0.0.11:53: read udp 127.0.0.1:47323-&gt;127.0.0.11:53: i/o timeout</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Incorrect HTTP Status Code</h5>
                                <small>Last month</small>
                            </div>
                            <p class="mb-1">HTTP Status Code 503 did not match 200</p>
                        </a>

                        <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1">Incorrect HTTP Status Code</h5>
                                <small>Last month</small>
                            </div>
                            <p class="mb-1">HTTP Status Code 503 did not match 200</p>
                        </a>

                    </div>

                </div>

                <div class="tab-pane fade" id="incidents" role="serviceLists" aria-labelledby="incidents-tab">

                </div>

                <div class="tab-pane fade" id="checkins" role="serviceLists" aria-labelledby="checkins-tab">



                    <div class="card">
                        <div class="card-body">
                            <form class="ajax_form" action="api/checkin" data-redirect="/service/7" method="POST">
                                <div class="form-group row">
                                    <div class="col-md-3">
                                        <label for="checkin_interval" class="col-form-label">Checkin Name</label>
                                        <input type="text" name="name" class="form-control" id="checkin_name" placeholder="New Checkin">
                                    </div>
                                    <div class="col-3">
                                        <label for="checkin_interval" class="col-form-label">Interval (seconds)</label>
                                        <input type="number" name="interval" class="form-control" id="checkin_interval" placeholder="60">
                                    </div>
                                    <div class="col-3">
                                        <label for="grace_period" class="col-form-label">Grace Period</label>
                                        <input type="number" name="grace" class="form-control" id="grace_period" placeholder="10">
                                    </div>
                                    <div class="col-3">
                                        <label for="submit" class="col-form-label"></label>
                                        <input type="hidden" name="service_id" class="form-control" id="service_id" value="7">
                                        <button type="submit" id="submit" class="btn btn-success d-block" style="margin-top: 14px;">Save Checkin</button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>



                </div>

                <div class="tab-pane fade" id="response" role="serviceLists" aria-labelledby="response-tab">
                    <div class="col-12 mt-4">
                        <h3>Last Response</h3>
                        <textarea rows="8" class="form-control" readonly>invalid route</textarea>
                        <div class="form-group row mt-2">
                            <label for="last_status_code" class="col-sm-3 col-form-label">HTTP Status Code</label>
                            <div class="col-sm-2">
                                <input type="text" id="last_status_code" class="form-control" value="200" readonly>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="tab-pane fade show active" id="edit" role="serviceLists" aria-labelledby="edit-tab">

                    <div class="card">
                        <div class="card-body">

                            <form class="ajax_form" action="api/services/7" data-redirect="services" method="POST">
                                <h4 class="mb-5 text-muted">Basic Information</h4>
                                <div class="form-group row">
                                    <label for="service_name" class="col-sm-4 col-form-label">Service Name</label>
                                    <div class="col-sm-8">
                                        <input type="text" name="name" class="form-control" id="service_name" value="Statping API" placeholder="Name" required spellcheck="false" autocorrect="off">
                                        <small class="form-text text-muted">Give your service a name you can recognize</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_type" class="col-sm-4 col-form-label">Service Type</label>
                                    <div class="col-sm-8">
                                        <select name="type" class="form-control" id="service_type" value="http" readonly>
                                            <option value="http" selected>HTTP Service</option>
                                            <option value="tcp" >TCP Service</option>
                                            <option value="udp" >UDP Service</option>
                                            <option value="icmp" >ICMP Ping</option>
                                        </select>
                                        <small class="form-text text-muted">Use HTTP if you are checking a website or use TCP if you are checking a server</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_url" class="col-sm-4 col-form-label">Application Endpoint (URL)</label>
                                    <div class="col-sm-8">
                                        <input type="text" name="domain" class="form-control" id="service_url" value="https://api.statping.com" placeholder="https://google.com" required autocapitalize="none" spellcheck="false">
                                        <small class="form-text text-muted">Statping will attempt to connect to this URL</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_type" class="col-sm-4 col-form-label">Group</label>
                                    <div class="col-sm-8">
                                        <select name="group_id" class="form-control" id="group_id">
                                            <option value="0" >None</option>

                                            <option value="1" >JSON Test Servers</option>

                                            <option value="2" >Google Servers</option>

                                            <option value="3" selected>Statping Servers</option>

                                        </select>
                                        <small class="form-text text-muted">Attach this service to a group</small>
                                    </div>
                                </div>

                                <h4 class="mt-5 mb-5 text-muted">Request Details</h4>

                                <div class="form-group row">
                                    <label for="service_check_type" class="col-sm-4 col-form-label">Service Check Type</label>
                                    <div class="col-sm-8">
                                        <select name="method" class="form-control" id="service_check_type" value="GET">
                                            <option value="GET" selected>GET</option>
                                            <option value="POST" >POST</option>
                                            <option value="DELETE" >DELETE</option>
                                            <option value="PATCH" >PATCH</option>
                                            <option value="PUT" >PUT</option>
                                        </select>
                                        <small class="form-text text-muted">A GET request will simply request the endpoint, you can also send data with POST.</small>
                                    </div>
                                </div>
                                <div class="form-group row d-none">
                                    <label for="post_data" class="col-sm-4 col-form-label">Optional Post Data (JSON)</label>
                                    <div class="col-sm-8">
                                        <textarea name="post_data" class="form-control" id="post_data" rows="3" autocapitalize="none" spellcheck="false" placeholder='{"data": { "method": "success", "id": 148923 } }'></textarea>
                                        <small class="form-text text-muted">Insert a JSON string to send data to the endpoint.</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="headers" class="col-sm-4 col-form-label">HTTP Headers</label>
                                    <div class="col-sm-8">
                                        <input name="headers" class="form-control" id="headers" autocapitalize="none" spellcheck="false" placeholder='Authorization=1010101,Content-Type=application/json' value="">
                                        <small class="form-text text-muted">Comma delimited list of HTTP Headers (KEY=VALUE,KEY=VALUE)</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_response" class="col-sm-4 col-form-label">Expected Response (Regex)</label>
                                    <div class="col-sm-8">
                                        <textarea name="expected" class="form-control" id="service_response" rows="3" autocapitalize="none" spellcheck="false" placeholder='(method)": "((\\"|[success])*)"'></textarea>
                                        <small class="form-text text-muted">You can use plain text or insert <a target="_blank" href="https://regex101.com/r/I5bbj9/1">Regex</a> to validate the response</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_response_code" class="col-sm-4 col-form-label">Expected Status Code</label>
                                    <div class="col-sm-8">
                                        <input type="number" name="expected_status" class="form-control" value="200" placeholder="200" id="service_response_code">
                                        <small class="form-text text-muted">A status code of 200 is success, or view all the <a target="_blank" href="https://www.restapitutorial.com/httpstatuscodes.html">HTTP Status Codes</a></small>
                                    </div>
                                </div>
                                <div class="form-group row d-none">
                                    <label for="port" class="col-sm-4 col-form-label">TCP Port</label>
                                    <div class="col-sm-8">
                                        <input type="number" name="port" class="form-control" value="" id="service_port" placeholder="8080">
                                    </div>
                                </div>

                                <h4 class="mt-5 mb-5 text-muted">Additional Options</h4>

                                <div class="form-group row">
                                    <label for="service_interval" class="col-sm-4 col-form-label">Check Interval (Seconds)</label>
                                    <div class="col-sm-8">
                                        <input type="number" name="check_interval" class="form-control" value="60" min="1" id="service_interval" required>
                                        <small id="interval" class="form-text text-muted">10,000+ will be checked in Microseconds (1 millisecond = 1000 microseconds).</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="service_timeout" class="col-sm-4 col-form-label">Timeout in Seconds</label>
                                    <div class="col-sm-8">
                                        <input type="number" name="timeout" class="form-control" value="15" placeholder="15" id="service_timeout" min="1">
                                        <small class="form-text text-muted">If the endpoint does not respond within this time it will be considered to be offline</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="post_data" class="col-sm-4 col-form-label">Permalink URL</label>
                                    <div class="col-sm-8">
                                        <input type="text" name="permalink" class="form-control" value="" id="permalink" autocapitalize="none" spellcheck="true" placeholder='awesome_service'>
                                        <small class="form-text text-muted">Use text for the service URL rather than the service number.</small>
                                    </div>
                                </div>
                                <div class="form-group row d-none">
                                    <label for="order" class="col-sm-4 col-form-label">List Order</label>
                                    <div class="col-sm-8">
                                        <input type="number" name="order" class="form-control" min="0" value="0" id="order">
                                        <small class="form-text text-muted">You can also drag and drop services to reorder on the Services tab.</small>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="order" class="col-sm-4 col-form-label">Verify SSL</label>
                                    <div class="col-8 mt-1">
            <span class="switch float-left">
                <input type="checkbox" name="verify_ssl-option" class="switch" id="switch-verify-ssl" >
                <label for="switch-verify-ssl">Verify SSL Certificate for this service</label>
                <input type="hidden" name="verify_ssl" id="switch-verify-ssl-value" value="false">
            </span>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="order" class="col-sm-4 col-form-label">Notifications</label>
                                    <div class="col-8 mt-1">
            <span class="switch float-left">
                <input type="checkbox" name="allow_notifications-option" class="switch" id="switch-notifications" checked>
                <label for="switch-notifications">Allow notifications to be sent for this service</label>
                <input type="hidden" name="allow_notifications" id="switch-notifications-value" value="true">
            </span>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="order" class="col-sm-4 col-form-label">Visible</label>
                                    <div class="col-8 mt-1">
            <span class="switch float-left">
                <input type="checkbox" name="public-option" class="switch" id="switch-public" checked>
                <label for="switch-public">Show service details to the public</label>
                <input type="hidden" name="public" id="switch-public-value" value="true">
            </span>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <div class="col-6">
                                        <button type="submit" class="btn btn-success btn-block">Update Service</button>
                                    </div>

                                    <div class="col-6">
                                        <a href="service/7/delete_failures" data-method="GET" data-redirect="/service/7" class="btn btn-danger btn-block confirm-btn">Delete All Failures</a>
                                    </div>

                                </div>
                                <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                            </form>
                        </div>
                    </div>

                </div>

            </div>

        </div>
    </div>

</template>

<script>
  import Api from "../components/API"
  import TopNav from "../components/Dashboard/TopNav";

  export default {
  name: 'Service',
  components: {
    TopNav
  },
  data () {
    return {
      id: null,
      service: null,
    }
  },
  mounted() {
    this.service = this.$route.params.service
    this.id = this.$route.params.id
    if (!this.service) {
      this.getService(this.id)
    }
  },
  beforeMount() {

  },
  methods: {
    async getService(id) {
        this.service = await Api.service(id)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
