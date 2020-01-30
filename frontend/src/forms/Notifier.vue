<template>
    <form @submit="saveNotifier">
        <h4 class="text-capitalize">{{notifier.title}}</h4>
        <p class="small text-muted" v-html="notifier.description"></p>

        <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>
            <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[form.field]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <small class="form-text text-muted" v-html="form.small_text"></small>
        </div>

        <div class="row">
            <div class="col-9 col-sm-6">
                <div class="input-group mb-2">
                    <div class="input-group-prepend">
                        <div class="input-group-text">Limit</div>
                    </div>
                    <input v-model="notifier.limits" type="number" class="form-control" name="limits" min="1" max="60" placeholder="7">
                    <div class="input-group-append">
                        <div class="input-group-text">Per Minute</div>
                    </div>
                </div>
            </div>

            <div class="col-3 col-sm-2 mt-1">
                <span @click="notifier.enabled = !!notifier.enabled" class="switch">
                    <input type="checkbox" name="enabled-option" class="switch" v-model="notifier.enabled" v-bind:id="`switch-${notifier.method}`" v-bind:checked="notifier.enabled">
                    <label v-bind:for="`switch-${notifier.method}`"></label>
                </span>
            </div>

            <div class="col-12 col-sm-4 mb-2 mb-sm-0 mt-2 mt-sm-0">
                <button @click="saveNotifier" type="submit" class="btn btn-block text-capitalize" :class="{'btn-primary': !saved, 'btn-success': saved}">
                    <i class="fa fa-check-circle"></i> {{saved ? "Saved" : "Save"}}
                </button>
            </div>

            <div class="col-12 col-sm-12">
                <button @click="testNotifier" class="btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i> Test Notifier</button>
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

        <div v-if="error" class="alert alert-danger d-none" id="alerter" role="alert"></div>
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
        error: null,
      saved: false,
    }
  },
  mounted() {

  },
  methods: {
    async saveNotifier(e) {
      e.preventDefault();
      let form = {}
      this.notifier.form.forEach((f) => {
        form[f.field] = this.notifier[f.field]
      });
      form.enabled = this.notifier.enabled
      form.limits = parseInt(this.notifier.limits)
      form.method = this.notifier.method
      await Api.notifier_save(form)
      const notifiers = await Api.notifiers()
      this.$store.commit('setNotifiers', notifiers)
      this.saved = true
      setTimeout(() => {
        this.saved = false
      }, 2000)
    },
    async testNotifier(e) {
      e.preventDefault();
      let form = {}
      this.notifier.form.forEach((f) => {
        form[f.field] = this.notifier[f.field]
      });
      form.enabled = this.notifier.enabled
      form.limits = parseInt(this.notifier.limits)
      form.method = this.notifier.method
      alert(JSON.stringify(form))
      const tested = await Api.notifier_test(form)
      if (tested === "ok") {
        alert('This notifier seems to be working!')
      } else {
        this.error = tested
      }
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
