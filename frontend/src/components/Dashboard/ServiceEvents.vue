<template>
  <div class="row p-2">

    <div v-if="loaded && last_failure && failureBefore" class="col-12 text-danger font-2 m-0 mb-2">
      <font-awesome-icon icon="exclamation" class="mr-1 text-danger font-weight-bold" size="1x"/> Recent Failure<br>
      <span class="font-italic font-weight-light text-dim mt-1" style="max-width: 270px">
      Last failure was {{ago(last_failure.created_at)}} ago. {{last_failure.issue}}
      </span>
    </div>

    <div v-if="loaded" v-for="message in messages" class="col-12 font-2 m-0 mb-2">
      <font-awesome-icon icon="calendar" class="mr-1" size="1x"/> Upcoming Announcement<br>
      <span class="font-italic font-weight-light text-dim mt-1">{{message.description}}</span>
      <span class="font-0 text-dim float-right font-weight-light mt-1">@ <strong>{{niceDate(message.start_on)}}</strong>
      </span>
    </div>

    <div v-if="loaded" v-for="incident in incidents" class="col-12 font-2 m-0 mb-2">
      <font-awesome-icon icon="bullhorn" class="mr-1" size="1x"/>Recent Incident<br>
      <span class="font-italic font-weight-light text-dim mt-1" style="max-width: 270px">{{incident.title}} - {{incident.description}}</span>
      <span class="font-0 text-dim float-right font-weight-light mt-1">@ <strong>{{niceDate(incident.created_at)}}</strong></span>
    </div>

    <div v-if="success_event && !failureBefore" class="col-12 font-2 m-0 mb-2">
      <span class="text-success"><font-awesome-icon icon="check" class="mr-1" size="1x"/>No New Events</span>
      <span class="font-italic d-inline-block text-truncate text-dim mt-1" style="max-width: 270px">
        Last failure was {{ago(service.last_error)}} ago.
      </span>
    </div>

  </div>
</template>

<script>
import Api from "../../API";
const Loading = () => import(/* webpackChunkName: "index" */ "@/components/Elements/Loading");

export default {
name: "ServiceEvents",
  components: {
    Loading
  },
  props: {
    service: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      incidents: null,
      loaded: false,
    }
  },
  mounted() {
   this.load()
  },
  computed: {
  last_failure() {
    if (!this.service.failures) {
      return null
    }
    return this.service.failures[0]
  },
  failureBefore() {
    return this.isAfter(this.parseISO(this.service.last_error), this.nowSubtract(43200).toISOString())
  },
    messages() {
      return this.$store.getters.serviceMessages(this.service.id)
    },
    success_event() {
      if (this.service.online && this.service.messages.length === 0 && this.service.incidents.length === 0) {
        return true
      }
      return false
    }
  },
  methods: {
    async load() {
      this.loaded = false
      await this.getMessages()
      await this.getIncidents()
      this.loaded = true
    },
    async getMessages() {
      // this.messages = this.$store.getters.serviceMessages(this.service.id)
    },
    async getIncidents() {
      this.incidents = await Api.incidents_service(this.service.id)
    },
  }
}
</script>
