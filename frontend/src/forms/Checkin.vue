<template>
    <form @submit="saveCheckin">
        <div class="form-group row">
            <div class="col-md-3">
                <label for="checkin_interval" class="col-form-label">Checkin Name</label>
                <input v-model="checkin.name" type="text" name="name" class="form-control" id="checkin_name" placeholder="New Checkin">
            </div>
            <div class="col-3">
                <label for="checkin_interval" class="col-form-label">Interval (seconds)</label>
                <input v-model="checkin.interval" type="number" name="interval" class="form-control" id="checkin_interval" placeholder="60">
            </div>
            <div class="col-3">
                <label for="grace_period" class="col-form-label">Grace Period</label>
                <input v-model="checkin.grace" type="number" name="grace" class="form-control" id="grace_period" placeholder="10">
            </div>
            <div class="col-3">
                <button @click="saveCheckin" type="submit" id="submit" class="btn btn-success d-block" style="margin-top: 14px;">Save Checkin</button>
            </div>
        </div>
    </form>
</template>

<script>
  import Api from "../components/API";

  export default {
  name: 'Checkin',
  props: {
    service: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      checkin: {
        name: "",
        interval: 60,
        grace: 60,
        service: this.service.id
      }
    }
  },
  mounted() {

  },
  methods: {
    async saveCheckin(e) {
      e.preventDefault();
      const data = {name: this.group.name, public: this.group.public}
      await Api.group_create(data)
      const groups = await Api.groups()
      this.$store.commit('setGroups', groups)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
