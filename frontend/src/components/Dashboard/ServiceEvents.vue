<template>
  <div>
    <Loading :loading="!loaded"/>
    <div v-if="loaded && !service.online" class="bg-white shadow-sm mt-3 p-3 pr-4 pl-4 col-12">
      <font-awesome-icon icon="exclamation" class="mr-3" size="1x"/>
      Last failure was {{ago(service.last_error)}} ago.
      <code v-if="failure" class="d-block bg-light p-3 mt-3">
        {{failure.issue}}
        <span class="d-block text-dim float-right small mt-3 mb-1">Failure #{{failure.id}}</span>
      </code>
    </div>
    <div v-if="loaded" v-for="message in messages" class="bg-light shadow-sm p-3 pr-4 pl-4 col-12 mt-3">
      <font-awesome-icon icon="calendar" class="mr-3" size="1x"/> {{message.description}}
      <span class="d-block small text-muted mt-3">
        Starts at <strong>{{niceDate(message.start_on)}}</strong> till <strong>{{niceDate(message.end_on)}}</strong>
        ({{dur(parseISO(message.start_on), parseISO(message.end_on))}})
      </span>
    </div>
    <div v-if="loaded" v-for="incident in incidents" class="bg-light shadow-sm p-3 pr-4 pl-4 col-12 mt-3">
      <font-awesome-icon icon="calendar" class="mr-3" size="1x"/>
      {{incident.title}} - {{incident.description}}
      <div v-for="update in incident.updates" class="d-block small">
        <span class="font-weight-bold text-capitalize">{{update.type}}</span> - {{update.message}}
      </div>
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
      messages: null,
      incidents: null,
      failure: null,
      loaded: false,
    }
  },
  mounted() {
   this.load()
  },
  methods: {
    async load() {
      this.loaded = false
      if (!this.service.online) {
        await this.getFailure()
      }
      await this.getMessages()
      await this.getIncidents()
      this.loaded = true
    },
    async getMessages() {
      this.messages = await Api.messages()
    },
    async getFailure() {
      const f = await Api.service_failures(this.service.id, null, null, 1)
      this.failure = f[0]
    },
    async getIncidents() {
      this.incidents = await Api.incidents_service(this.service.id)
    },
  }
}
</script>

<style scoped>

</style>
