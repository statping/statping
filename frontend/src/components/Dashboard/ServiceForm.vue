<template>
    <div class="col-12">

        <div class="card">
            <div class="card-body">
    <form class="ajax_form" action="api/services" data-redirect="services" method="POST">
        <h4 class="mb-5 text-muted">Basic Information</h4>
        <div class="form-group row">
            <label for="service_name" class="col-sm-4 col-form-label">Service Name</label>
            <div class="col-sm-8">
                <input type="text" name="name" class="form-control" id="service_name" value="" placeholder="Name" required spellcheck="false" autocorrect="off">
                <small class="form-text text-muted">Give your service a name you can recognize</small>
            </div>
        </div>
        <div class="form-group row">
            <label for="service_type" class="col-sm-4 col-form-label">Service Type</label>
            <div class="col-sm-8">
                <select name="type" class="form-control" id="service_type" value="" >
                    <option value="http" >HTTP Service</option>
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
                <input type="text" name="domain" class="form-control" id="service_url" value="" placeholder="https://google.com" required autocapitalize="none" spellcheck="false">
                <small class="form-text text-muted">Statping will attempt to connect to this URL</small>
            </div>
        </div>
        <div class="form-group row">
            <label for="service_type" class="col-sm-4 col-form-label">Group</label>
            <div class="col-sm-8">
                <select name="group_id" class="form-control" id="group_id">
                    <option value="0" selected>None</option>

                    <option value="1" >JSON Test Servers</option>

                    <option value="2" >Google Servers</option>

                    <option value="3" >Statping Servers</option>

                </select>
                <small class="form-text text-muted">Attach this service to a group</small>
            </div>
        </div>

        <h4 class="mt-5 mb-5 text-muted">Request Details</h4>

        <div class="form-group row">
            <label for="service_check_type" class="col-sm-4 col-form-label">Service Check Type</label>
            <div class="col-sm-8">
                <select name="method" class="form-control" id="service_check_type" value="">
                    <option value="GET" >GET</option>
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
                <input type="checkbox" name="verify_ssl-option" class="switch" id="switch-verify-ssl" checked>
                <label for="switch-verify-ssl">Verify SSL Certificate for this service</label>
                <input type="hidden" name="verify_ssl" id="switch-verify-ssl-value" value="true">
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
            <div class="col-12">
                <button type="submit" class="btn btn-success btn-block">Create Service</button>
            </div>

        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'ServiceForm',
  props: {
    service: Object
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
