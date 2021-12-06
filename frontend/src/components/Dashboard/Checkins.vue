<template>
    <div class="col-12">
        <h2>{{service.name}} Checkins</h2>
        <p class="mb-3">Tell your service to send a routine HTTP request to a Statping Checkin.</p>

        <div v-for="(checkin, i) in checkins" class="card text-black-50 bg-white mt-3">
            <div class="card-header text-capitalize">
                {{checkin.name}}
                <button @click="deleteCheckin(checkin)" class="btn btn-sm small btn-danger float-right text-uppercase">Delete</button>
            </div>
            <div class="card-body">
                <div class="input-group">
                    <input type="text" class="form-control" :value="`${core.domain}/checkin/${checkin.api_key}`" readonly>
                    <div class="input-group-append copy-btn">
                        <button @click.prevent="copy(`${core.domain}/checkin/${checkin.api_key}`)" class="btn btn-outline-secondary" type="button">Copy</button>
                    </div>
                </div>

                <span class="small">Send a GET request to this URL every {{checkin.interval}} minutes</span>
                <span class="small float-right mt-1 mr-3 d-none d-md-block">Requested {{ago(checkin.last_hit)}} ago</span>
                <span class="small float-right mt-1 mr-3 d-none d-md-block">Request expected every {{checkin.interval}} minutes</span>

                <div class="card text-black-50 bg-white mt-3">
                    <div class="card-header text-capitalize">
                        <font-awesome-icon @click="expanded = !expanded" :icon="expanded ? 'minus' : 'plus'" class="mr-2 pointer"/>
                        {{checkin.name}} Records
                    </div>
                    <div class="card-body" :class="{'d-none': !expanded}">
                        <div class="alert alert-primary small" :class="{'alert-success': hit.success, 'alert-danger': !hit.success}" v-for="(hit, i) in records(checkin)">
                            Checkin {{hit.success ? "Request" : "Failure"}} at {{hit.created_at}}
                        </div>
                    </div>
                </div>

                <div class="card text-black-50 bg-white mt-3">
                    <div class="card-header text-capitalize">
                        <font-awesome-icon @click="curl_expanded = !curl_expanded" :icon="curl_expanded ? 'minus' : 'plus'" class="mr-2 pointer"/>
                        Cronjob Task
                    </div>
                    <div class="card-body" :class="{'d-none': !curl_expanded}">
                        This cronjob script will request the checkin endpoint every {{checkin.interval}} minutes. Add this cronjob task to the machine running this service.
                        <div class="input-group mt-2">
                            <input type="text" class="form-control" :value="`${checkin.interval} * * * * /usr/bin/curl ${core.domain}/checkin/${checkin.api_key} >/dev/null 2>&1`" readonly>
                            <div class="input-group-append copy-btn">
                                <button @click.prevent="copy(`${checkin.interval} * * * * /usr/bin/curl ${core.domain}/checkin/${checkin.api_key} >/dev/null 2>&1`)" class="btn btn-outline-secondary" type="button">Copy</button>
                            </div>
                        </div>
                        <span class="small d-block">Using CURL</span>
                    </div>
                </div>
            </div>
            <div class="card-footer">
                <span :class="{'text-success': last_record(checkin).success, 'text-danger': !last_record(checkin).success}">
                    {{last_record(checkin).success ? "Checkin is currently working correctly" : "Checkin is currently failing"}}
                </span>
            </div>
        </div>

        <div class="card text-black-50 bg-white mt-4">
            <div class="card-header text-capitalize">Create Checkin</div>
            <div class="card-body">
            <form @submit.prevent="saveCheckin">
                <div class="form-group row">
                    <div class="col-7 col-md-5">
                        <label for="checkin_interval" class="col-form-label">Checkin Name</label>
                        <input v-model="checkin.name" type="text" name="name" class="form-control" id="checkin_name" placeholder="New Checkin">
                    </div>
                    <div class="col-5 col-md-3">
                        <label for="checkin_interval" class="col-form-label">Interval (minutes)</label>
                        <input v-model="checkin.interval" type="number" name="interval" class="form-control" id="checkin_interval" placeholder="1" min="1">
                    </div>
                    <div class="col-12 col-md-4">
                        <label class="col-form-label"></label>
                        <button :disabled="btn_disabled" @click.prevent="saveCheckin" type="submit" id="submit" class="btn btn-primary d-block mt-2">Save Checkin</button>
                    </div>
                </div>
            </form>
            </div>
        </div>
    </div>
</template>

<script>
import Api from "../../API";

export default {
    name: 'Checkins',
    data() {
        return {
            service: {},
          ready: false,
          expanded: false,
          curl_expanded: false,
          checkin: {
            name: "",
            interval: 1,
            service_id: 0,
            hits: [],
            failures: []
          }
        }
    },
      computed: {
        checkins() {
          return this.$store.getters.serviceCheckins(this.service.id)
        },
        core() {
          return this.$store.getters.core
        },
        btn_disabled() {
          if (this.checkin.name === "" || this.checkin.interval <= 0) {
            return true
          }
          return false
        },
      },
      async created() {
          if (this.$route.params) {
            const id = this.$route.params.id
            this.service = await Api.service(id)
            this.checkin.service_id = this.service.id
            this.ready = true
          }
      },
    methods: {
      records(checkin) {
        let hits = []
        let failures = []
        checkin.hits.forEach((hit) => {
          hits.push({success: true, created_at: this.parseISO(hit.created_at), id: hit.id})
        })
        checkin.failures.forEach((failure) => {
          failures.push({success: false, created_at: this.parseISO(failure.created_at), id: failure.id})
        })
        return hits.concat(failures).sort((a, b) => {return a.created_at-b.created_at}).reverse().slice(0,32)
      },
      last_record(checkin) {
        const r = this.records(checkin)
        if (r.length === 0) {
          return {success: false}
        }
        return r[0]
      },
      fixInts() {
        const c = this.checkin
        this.checkin.interval = parseInt(c.interval)
        return this.checkin
      },
      async saveCheckin() {
        const c = this.fixInts()
        await Api.checkin_create(c)
        this.checkin.name = ""
        await this.load()
      },
      async deleteCheckin(checkin) {
        await Api.checkin_delete(checkin)
        await this.load()
      },
      async load() {
        const checkins = await Api.checkins()
        this.$store.commit('setCheckins', checkins)
      }
    }
}
</script>
