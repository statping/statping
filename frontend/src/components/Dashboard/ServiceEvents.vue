<template>
  <div>
    <Loading :loading="!loaded"/>
    <div v-if="loaded && isBefore(nowSubtract(86400), service.last_error)" class="text-danger font-2 p-0 m-0 mb-2">
      <font-awesome-icon icon="exclamation" class="mr-1 text-danger" size="1x"/> Recent Failure<br>
      <span class="font-italic d-inline-block text-dim" style="max-width: 270px">
      Last failure was {{ago(service.last_error)}} ago. {{service.failures[0].issue}}
      </span>
    </div>


    <div v-if="loaded" v-for="message in messages" class="font-2 p-0 m-0 mb-2">
      <font-awesome-icon icon="calendar" class="mr-1" size="1x"/> Upcoming Announcement<br>
      <span class="font-italic font-weight-light">{{message.description}}</span>
      <span class="font-0 text-dim float-right font-weight-light">@ <strong>{{niceDate(message.start_on)}}</strong>
      </span>
    </div>
    <div v-if="loaded" v-for="incident in incidents" class="font-2 p-0 m-0 mb-2">
      <font-awesome-icon icon="bullhorn" class="mr-1" size="1x"/>Recent Incident<br>
      <span class="font-italic d-inline-block text-truncate" style="max-width: 270px">{{incident.title}} - {{incident.description}}</span>
      <span class="font-0 text-dim float-right font-weight-light">@ <strong>{{niceDate(incident.created_at)}}</strong></span>
    </div>

    <div v-if="success_event" class="font-2 p-0 m-0 mt-1 mb-3">
      <span class="text-success"><font-awesome-icon icon="check" class="mr-1" size="1x"/>No New Events</span>
      <span class="font-italic d-inline-block text-truncate text-dim" style="max-width: 270px">Last failure was {{ago(service.last_error)}} ago.</span>
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
      failure: null,
      loaded: false,
    }
  },
  mounted() {
   this.load()
  },
  computed: {
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
      if (!this.service.online) {
        this.failure = this.service.failures[0]
      }
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
