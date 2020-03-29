<template>
    <form @submit.prevent="saveService">
        <div class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">Basic Information</div>
            <div class="card-body">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Service Name</label>
            <div class="col-sm-8">
                <input v-model="service.name" @keypress="service.permalink=service.name.split(' ').join('_')" type="text" name="name" class="form-control" placeholder="Name" required spellcheck="false" autocorrect="off">
                <small class="form-text text-muted">Give your service a name you can recognize</small>
            </div>
        </div>
        <div class="form-group row">
            <label for="service_type" class="col-sm-4 col-form-label">Service Type</label>
            <div class="col-sm-8">
                <select v-model="service.type" class="form-control" id="service_type" >
                    <option value="http">HTTP Service</option>
                    <option value="grpc">gRPC Service</option>
                    <option value="tcp">TCP Service</option>
                    <option value="udp">UDP Service</option>
                    <option value="icmp">ICMP Ping</option>
                </select>
                <small class="form-text text-muted">Use HTTP if you are checking a website or use TCP if you are checking a server</small>
            </div>
        </div>
        <div class="form-group row">
            <label for="service_url" class="col-sm-4 col-form-label">Application Endpoint (URL)</label>
            <div class="col-sm-8">
                <input v-model="service.domain" type="text" class="form-control" id="service_url" placeholder="https://google.com" required autocapitalize="none" spellcheck="false">
                <small class="form-text text-muted">Statping will attempt to connect to this URL</small>
            </div>
        </div>
        <div class="form-group row">
            <label for="service_type" class="col-sm-4 col-form-label">Group</label>
            <div class="col-sm-8">
                <select v-model="service.group_id" class="form-control">
                    <option value="0" >No Group</option>
                    <option v-for="(group, index) in $store.getters.cleanGroups()" :value="group.id">{{group.name}}</option>
                </select>
                <small class="form-text text-muted">Attach this service to a group</small>
            </div>
        </div>
            </div>
        </div>

        <div v-if="service.type !== 'icmp'" class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">Request Details</div>
            <div class="card-body">

        <div v-if="service.type.match(/^(http)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">Service Check Type</label>
            <div class="col-sm-8">
                <select v-model="service.method" name="method" class="form-control">
                    <option value="GET" >GET</option>
                    <option value="POST" >POST</option>
                    <option value="DELETE" >DELETE</option>
                    <option value="PATCH" >PATCH</option>
                    <option value="PUT" >PUT</option>
                </select>
                <small class="form-text text-muted">A GET request will simply request the endpoint, you can also send data with POST.</small>
            </div>
        </div>
        <div v-if="service.type.match(/^(http)$/) && service.method.match(/^(POST|PATCH|DELETE|PUT)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">Optional Post Data (JSON)</label>
            <div class="col-sm-8">
                <textarea v-model="service.post_data" class="form-control" rows="3" autocapitalize="none" spellcheck="false" placeholder='{"data": { "method": "success", "id": 148923 } }'></textarea>
                <small class="form-text text-muted">Insert a JSON string to send data to the endpoint.</small>
            </div>
        </div>
        <div v-if="service.type.match(/^(http)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">HTTP Headers</label>
            <div class="col-sm-8">
                <input v-model="service.headers" class="form-control" autocapitalize="none" spellcheck="false" placeholder='Authorization=1010101,Content-Type=application/json'>
                <small class="form-text text-muted">Comma delimited list of HTTP Headers (KEY=VALUE,KEY=VALUE)</small>
            </div>
        </div>
        <div v-if="service.type.match(/^(http)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">Expected Response (Regex)</label>
            <div class="col-sm-8">
                <textarea v-model="service.expected" class="form-control" rows="3" autocapitalize="none" spellcheck="false" placeholder='(method)": "((\\"|[success])*)"'></textarea>
                <small class="form-text text-muted">You can use plain text or insert <a target="_blank" href="https://regex101.com/r/I5bbj9/1">Regex</a> to validate the response</small>
            </div>
        </div>
        <div v-if="service.type.match(/^(http)$/)" class="form-group row">
            <label for="service_response_code" class="col-sm-4 col-form-label">Expected Status Code</label>
            <div class="col-sm-8">
                <input v-model="service.expected_status" type="number" name="expected_status" class="form-control" placeholder="200" id="service_response_code">
                <small class="form-text text-muted">A status code of 200 is success, or view all the <a target="_blank" href="https://www.restapitutorial.com/httpstatuscodes.html">HTTP Status Codes</a></small>
            </div>
        </div>
        <div v-if="service.type.match(/^(tcp|udp)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">{{service.type.toUpperCase()}} Port</label>
            <div class="col-sm-8">
                <input v-model="service.port" type="number" name="port" class="form-control" id="service_port" placeholder="8080">
            </div>
        </div>
            </div>
        </div>

        <div class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">Additional Options</div>
            <div class="card-body">

        <div class="form-group row">
            <label for="service_interval" class="col-sm-4 col-form-label">Check Interval (Seconds)</label>
            <div class="col-sm-8">
                <input v-model="service.check_interval" type="number" class="form-control" min="1" id="service_interval" required>
                <small id="interval" class="form-text text-muted">10,000+ will be checked in Microseconds (1 millisecond = 1000 microseconds).</small>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Timeout in Seconds</label>
            <div class="col-sm-8">
                <input v-model="service.timeout" type="number" name="timeout" class="form-control" placeholder="15" min="1">
                <small class="form-text text-muted">If the endpoint does not respond within this time it will be considered to be offline</small>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Permalink URL</label>
            <div class="col-sm-8">
                <input v-model="service.permalink" type="text" name="permalink" class="form-control" id="permalink" autocapitalize="none" spellcheck="true" placeholder='awesome_service'>
                <small class="form-text text-muted">Use text for the service URL rather than the service number.</small>
            </div>
        </div>
        <div v-if="service.type.match(/^(http)$/)" class="form-group row">
            <label class="col-sm-4 col-form-label">Verify SSL</label>
            <div class="col-8 mt-1">
            <span @click="service.verify_ssl = !!service.verify_ssl" class="switch float-left">
                <input v-model="service.verify_ssl" type="checkbox" name="verify_ssl-option" class="switch" id="switch-verify-ssl" v-bind:checked="service.verify_ssl">
                <label for="switch-verify-ssl">Verify SSL Certificate for this service</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Notifications</label>
            <div class="col-8 mt-1">
            <span @click="service.allow_notifications = !!service.allow_notifications" class="switch float-left">
                <input v-model="service.allow_notifications" type="checkbox" name="allow_notifications-option" class="switch" id="switch-notifications" v-bind:checked="service.allow_notifications">
                <label for="switch-notifications">Allow notifications to be sent for this service</label>
            </span>
            </div>
        </div>
        <div v-if="service.allow_notifications"  class="form-group row">
            <label class="col-sm-4 col-form-label">Notify After Failures</label>
            <div class="col-sm-8">
                <input v-model="service.notify_after" type="number" name="notify_after" class="form-control" id="notify_after" autocapitalize="none">
                <small class="form-text text-muted">Send Notification after {{service.notify_after === 0 ? 'the first Failure' : service.notify_after+' Failures'}} </small>
            </div>
        </div>
        <div v-if="service.allow_notifications" class="form-group row">
            <label class="col-sm-4 col-form-label">Notify All Changes</label>
            <div class="col-8 mt-1">
            <span @click="service.notify_all_changes = !!service.notify_all_changes" class="switch float-left">
                <input v-model="service.notify_all_changes" type="checkbox" name="notify_all-option" class="switch" id="notify_all" v-bind:checked="service.notify_all_changes">
                <label for="notify_all">Continuously notify when service is failing.</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Visible</label>
            <div class="col-8 mt-1">
            <span @click="service.public = !!service.public" class="switch float-left">
                <input v-model="service.public" type="checkbox" name="public-option" class="switch" id="switch-public" v-bind:checked="service.public">
                <label for="switch-public">Show service details to the public</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-12">
                <button :disabled="loading" @click.prevent="saveService" type="submit" class="btn btn-success btn-block">
                    {{service.id ? "Update Service" : "Create Service"}}
                </button>
            </div>
        </div>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
  import Api from "../API";

  export default {
  name: 'FormService',
  data () {
    return {
        loading: false,
      service: {
        name: "",
        type: "http",
        domain: "",
        group_id: 0,
        method: "GET",
        post_data: "",
        headers: "",
        expected: "",
        expected_status: 200,
        port: 80,
        check_interval: 60,
        timeout: 15,
        permalink: "",
        order: 1,
        verify_ssl: true,
        allow_notifications: true,
        notify_all_changes: true,
        notify_after: 2,
        public: true,
      },
      groups: [],
    }
  },
    props: {
      in_service: {
        type: Object
      }
    },
    watch: {
      in_service() {
        this.service = this.in_service
      }
    },
  async mounted() {
    if (!this.$store.getters.groups) {
      const groups = await Api.groups()
      this.$store.commit('setGroups', groups)
    }
  },
  methods: {
    async saveService() {
      let s = this.service
        this.loading = true
      delete s.failures
      delete s.created_at
      delete s.updated_at
      delete s.last_success
      delete s.latency
      delete s.online_24_hours
        s.check_interval = parseInt(s.check_interval)

        window.console.log(s)

        if (s.id) {
            await this.updateService(s)
        } else {
            await this.createService(s)
        }
        const services = await Api.services()
        this.$store.commit('setServices', services)
        this.loading = false
        this.$router.push('/dashboard/services')
    },
    async createService(s) {
        await Api.service_create(s)
    },
      async updateService(s) {
          await Api.service_update(s)
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
