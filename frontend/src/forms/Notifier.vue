<template>
    <div>
    <div class="card contain-card text-black-50 bg-white mb-3">
        <div class="card-header text-capitalize">
            {{notifier.title}}
            <span @click="enableToggle" class="switch switch-rd-gr float-right">
                <input v-model="notifier.enabled" type="checkbox" class="switch-sm" :id="`enable_${notifier.method}`" v-bind:checked="notifier.enabled">
                <label class="mb-0" :for="`enable_${notifier.method}`"></label>
            </span>
        </div>
        <div class="card-body">

    <form @submit.prevent="saveNotifier">

        <p class="small text-muted" v-html="notifier.description"/>

        <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>
            <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[form.field.toLowerCase()]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <small class="form-text text-muted" v-html="form.small_text"></small>
        </div>

        <div class="row mt-4">

            <div class="col-sm-12">
                <span class="slider-info">Limit {{notifier.limits}} per hour</span>
                <input v-model="notifier.limits" type="range" name="limits" class="slider" min="1" max="300">
                <small class="form-text text-muted">Notifier '{{notifier.title}}' will send a maximum of {{notifier.limits}} notifications per hour.</small>
            </div>

        </div>

    </form>
        </div>
    </div>

        <div v-if="error && !success" class="alert alert-danger col-12" role="alert">
            {{error}}<p v-if="response">Response:<br>{{response}}</p>
        </div>
        <div v-if="success" class="alert alert-success col-12" role="alert">
            {{notifier.title}} appears to be working!
            <p v-if="response">Response:<br>{{response}}</p>
        </div>

        <div class="card text-black-50 bg-white mb-3">
            <div class="card-body">

                <div class="row">
                    <div class="col-6 col-sm-6 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="saveNotifier" type="submit" class="btn btn-block text-capitalize btn-primary">
                            <i class="fa fa-check-circle"></i> {{loading ? "Loading..." : saved ? "Saved" : "Save Settings"}}
                        </button>
                    </div>
                    <div class="col-6 col-sm-6 mb-2 mb-sm-0 mt-2 mt-sm-0">
                        <button @click.prevent="testNotifier" class="btn btn-outline-dark btn-block text-capitalize"><i class="fa fa-vial"></i>
                            {{loadingTest ? "Loading..." : "Test Notifier"}}</button>
                    </div>
                </div>

            </div>
        </div>

        <span class="d-block small text-center mb-3">
            <span class="text-capitalize">{{notifier.title}}</span> Notifier created by <a :href="notifier.author_url" target="_blank">{{notifier.author}}</a>
        </span>

    </div>
</template>

<script>
import Api from "../API";

export default {
    name: 'Notifier',
    props: {
        notifier: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            loading: false,
            loadingTest: false,
            error: null,
            response: null,
            success: false,
            saved: false,
            form: {},
        }
    },
    methods: {
        async enableToggle() {
            this.notifier.enabled = !!this.notifier.enabled
            const form = {
                enabled: !this.notifier.enabled,
                method: this.notifier.method,
            }
            await Api.notifier_save(form)
        },
        async saveNotifier() {
            this.loading = true
            this.form.enabled = this.notifier.enabled
            this.form.limits = parseInt(this.notifier.limits)
            this.form.method = this.notifier.method
            this.notifier.form.forEach((f) => {
              let field = f.field.toLowerCase()
              let val = this.notifier[field]
              if (this.isNumeric(val)) {
                val = parseInt(val)
              }
                this.form[field] = val
            });
            await Api.notifier_save(this.form)
            const notifiers = await Api.notifiers()
            await this.$store.commit('setNotifiers', notifiers)
            this.saved = true
            this.loading = false
        },
        async testNotifier() {
            this.success = false
            this.loadingTest = true
            this.form.method = this.notifier.method
            this.notifier.form.forEach((f) => {
                let field = f.field.toLowerCase()
                let val = this.notifier[field]
                if (this.isNumeric(val)) {
                    val = parseInt(val)
                }
                this.form[field] = val
            });
            const tested = await Api.notifier_test(this.form)
            if (tested.success) {
                this.success = true
            } else {
                this.error = tested.error
            }
            this.response = tested.response
            this.loadingTest = false
        },
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
