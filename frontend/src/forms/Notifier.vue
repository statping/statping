<template>
    <form class="ajax_form command">
        <h4 class="text-capitalize">{{notifier.title}}</h4>
        <p class="small text-muted" v-html="notifier.description"></p>

        <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>

            <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[notifier.field]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <small class="form-text text-muted" v-html="form.small_text"></small>

        </div>

        <div class="row">
            <div class="col-9 col-sm-6">
                <div class="input-group mb-2">
                    <div class="input-group-prepend">
                        <div class="input-group-text">Limit</div>
                    </div>
                    <input type="number" class="form-control" name="limits" min="1" max="60" id="limits_per_hour_command" value="3" placeholder="7">
                    <div class="input-group-append">
                        <div class="input-group-text">Per Minute</div>
                    </div>
                </div>
            </div>

            <div class="col-3 col-sm-2 mt-1">
                <span class="switch">
                    <input @change="notifier.enabled = !notifier.enabled" type="checkbox" name="enabled-option" class="switch" v-bind:id="`switch-${notifier.method}`" >
                    <label v-bind:for="`switch-${notifier.method}`"></label>
                    <input type="hidden" name="enabled" v-bind:id="`switch-${notifier.method}`">
                </span>
            </div>

            <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                <button type="submit" class="btn btn-primary btn-block text-capitalize"><i class="fa fa-check-circle"></i> Save</button>
            </div>

            <div class="col-12 col-sm-12">
                <button class="test_notifier btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
            </div>

            <div class="col-12 col-sm-12 mt-2">
                <div class="alert alert-danger d-none" id="command-error" role="alert">
                    <i class="fa fa-exclamation-triangle"></i> {{notifier.method}} has an error!
                </div>

                <div class="alert alert-success d-none" id="command-success" role="alert">
                    <i class="fa fa-smile-beam"></i> The {{notifier.method}} notifier is working correctly!
                </div>
            </div>
        </div>

        <span class="d-block small text-center mt-3 mb-5">
                                <span class="text-capitalize">{{notifier.title}}</span> Notifier created by <a :href="notifier.author_url" target="_blank">{{notifier.author}}</a>
                            </span>

        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
  import Api from "../components/API";

  export default {
  name: 'Notifier',
  props: {
    notifier: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      notifier: {

      }
    }
  },
  mounted() {

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
