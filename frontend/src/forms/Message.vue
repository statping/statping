<template>
    <form class="ajax_form" action="api/messages" data-redirect="messages" method="POST">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Title</label>
            <div class="col-sm-8">
                <input type="text" name="title" class="form-control" value="" id="title" placeholder="Message Title" required>
            </div>
        </div>

        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Description</label>
            <div class="col-sm-8">
                <textarea rows="5" name="description" class="form-control" id="description" required></textarea>
            </div>
        </div>

        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Message Date Range</label>
            <div class="col-sm-4">
                <input type="text" name="start_on" class="form-control form-control-plaintext" id="start_on" value="0001-01-01T00:00:00Z" required>
            </div>
            <div class="col-sm-4">
                <input type="text" name="end_on" class="form-control form-control-plaintext" id="end_on" value="0001-01-01T00:00:00Z" required>
            </div>
        </div>

        <div class="form-group row">
            <label for="service_id" class="col-sm-4 col-form-label">Service</label>
            <div class="col-sm-8">
                <select class="form-control" name="service" id="service_id">
                    <option value="0" selected>Global Message</option>


                    <option value="7" >Statping API</option>


                    <option value="6" >Push Notification Server</option>


                    <option value="1" >Google</option>


                    <option value="2" >Statping Github</option>


                    <option value="3" >JSON Users Test</option>


                    <option value="4" >JSON API Tester</option>


                    <option value="5" >Google DNS</option>

                </select>
            </div>
        </div>

        <div class="form-group row">
            <label for="notify_method" class="col-sm-4 col-form-label">Notification Method</label>
            <div class="col-sm-8">
                <input type="text" name="notify_method" class="form-control" id="notify_method" value="" placeholder="email">
            </div>
        </div>

        <div class="form-group row">
            <label for="notify_method" class="col-sm-4 col-form-label">Notify Users</label>
            <div class="col-sm-8">
                <span class="switch">
                    <input @click="message.notify = !!message.notify" type="checkbox" class="switch" id="switch-normal">
                    <label for="switch-normal">Notify Users Before Scheduled Time</label>
                </span>
            </div>
        </div>

        <div class="form-group row">
            <label for="notify_before" class="col-sm-4 col-form-label">Notify Before</label>
            <div class="col-sm-8">
                <div class="form-inline">
                    <input type="number" name="notify_before" class="col-4 form-control" id="notify_before" value="0">
                    <select class="ml-2 col-7 form-control" name="notify_before_scale" id="notify_before_scale">
                        <option value="minute">Minutes</option>
                        <option value="hour">Hours</option>
                        <option value="day">Days</option>
                    </select>
                </div>
            </div>
        </div>

        <div class="form-group row">
            <div class="col-sm-12">
                <button type="submit" class="btn btn-primary btn-block">Create Message</button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
import Api from "../components/API";

export default {
  name: 'FormMessage',
  props: {

  },
  data () {
    return {
      group: {
        name: "",
        public: true
      }
    }
  },
  mounted() {
    if (this.props.group) {
      this.group = this.props.group
    }
  },
  methods: {
    async saveGroup(e) {
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
