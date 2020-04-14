<template>
    <div class="col-12">
        <h2>{{service.name}} Checkins</h2>
        <p class="mb-3">Tell your service to send a routine HTTP request to a Statping Checkin.</p>
        <div v-for="(checkin, i) in checkins" class="col-12 alert alert-light" role="alert">
            <span class="badge badge-pill badge-info text-uppercase">{{checkin.name}}</span>
            <span class="float-right font-2">Last checkin {{ago(checkin.last_hit)}}</span>
            <span class="float-right font-2 mr-3">Check Every {{checkin.interval}} seconds</span>
            <span class="float-right font-2 mr-3">Grace Period {{checkin.grace}} seconds</span>
            <span class="d-block mt-2">
                <input type="text" class="form-control" :value="`${core.domain}/checkin/${checkin.api_key}`" readonly>
                <span class="small">Send a GET request to this URL every {{checkin.interval}} seconds
                    <button @click="deleteCheckin(checkin)" type="button" class="btn btn-danger btn-xs float-right mt-1">Delete</button>
                </span>
            </span>
        </div>

        <div class="col-12 alert alert-light">
            <form @submit.prevent="saveCheckin">
                <div class="form-group row">
                    <div class="col-5">
                        <label for="checkin_interval" class="col-form-label">Checkin Name</label>
                        <input v-model="checkin.name" type="text" name="name" class="form-control" id="checkin_name" placeholder="New Checkin">
                    </div>
                    <div class="col-2">
                        <label for="checkin_interval" class="col-form-label">Interval</label>
                        <input v-model="checkin.interval" type="number" name="interval" class="form-control" id="checkin_interval" placeholder="60">
                    </div>
                    <div class="col-2">
                        <label for="grace_period" class="col-form-label">Grace Period</label>
                        <input v-model="checkin.grace" type="number" name="grace" class="form-control" id="grace_period" placeholder="10">
                    </div>
                    <div class="col-3">
                        <label class="col-form-label"></label>
                        <button @click.prevent="saveCheckin" type="submit" id="submit" class="btn btn-primary d-block mt-2">Save Checkin</button>
                    </div>
                </div>
            </form>
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
          checkin: {
            name: "",
            interval: 60,
            grace: 60,
            service_id: 0
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
      fixInts() {
        const c = this.checkin
        this.checkin.interval = parseInt(c.interval)
        this.checkin.grace = parseInt(c.grace)
        return this.checkin
      },
      async saveCheckin() {
        const c = this.fixInts()
        await Api.checkin_create(c)
        await this.updateCheckins()
      },
      async deleteCheckin(checkin) {
        await Api.checkin_delete(checkin)
        await this.updateCheckins()
      },
      async updateCheckins() {
        const checkins = await Api.checkins()
        this.$store.commit('setCheckins', checkins)
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .sm {
        font-size: 8pt;
    }
</style>
