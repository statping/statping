<template>
    <div class="card text-black-50 bg-white mb-5">
        <div class="card-header text-capitalize">{{notifier.title}}</div>
        <div class="card-body">
    <form @submit.prevent="saveNotifier">

        <div v-if="error" class="alert alert-danger col-12" role="alert">{{error}}</div>

        <div v-if="ok" class="alert alert-success col-12" role="alert">
            <i class="fa fa-smile-beam"></i> The {{notifier.method}} notifier is working correctly!
        </div>

        <p class="small text-muted" v-html="notifier.description"/>

        <div v-for="(form, index) in notifier.form" v-bind:key="index" class="form-group">
            <label class="text-capitalize">{{form.title}}</label>
            <input v-if="form.type === 'text' || 'number' || 'password'" v-model="notifier[form.field.toLowerCase()]" :type="form.type" class="form-control" :placeholder="form.placeholder" >

            <small class="form-text text-muted" v-html="form.small_text"></small>
        </div>

        <div class="row mt-4">
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
                <button @click.prevent="saveNotifier" type="submit" class="btn btn-block text-capitalize" :class="{'btn-primary': !saved, 'btn-success': saved}">
                    <i class="fa fa-check-circle"></i> {{loading ? "Loading..." : saved ? "Saved" : "Save"}}
                </button>
            </div>

            <div class="col-12 col-sm-12 mt-3">
                <button @click.prevent="testNotifier" class="btn btn-secondary btn-block text-capitalize col-12 float-right"><i class="fa fa-vial"></i>
                    {{loading ? "Loading..." : "Test Notifier"}}</button>
            </div>

        </div>

    </form>
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
            error: null,
            saved: false,
            ok: false,
            form: {},
        }
    },
    mounted() {

    },
    methods: {
        async saveNotifier() {
            this.loading = true
            this.form.enabled = this.notifier.enabled
            this.form.limits = parseInt(this.notifier.limits)
            this.form.method = this.notifier.method
            this.notifier.form.forEach((f) => {
                let field = f.field.toLowerCase()
                this.form[field] = this.notifier[field]
            });
            await Api.notifier_save(this.form)
            const notifiers = await Api.notifiers()
            await this.$store.commit('setNotifiers', notifiers)
            this.saved = true
            this.loading = false
            setTimeout(() => {
                this.saved = false
            }, 2000)
        },
        async testNotifier() {
            this.ok = false
            this.loading = true
            let form = {}
            this.notifier.form.forEach((f) => {
                form[f.field] = this.notifier[f.field]
            });
            form.enabled = this.notifier.enabled
            form.limits = parseInt(this.notifier.limits)
            form.method = this.notifier.method
            const tested = await Api.notifier_test(form)
            if (tested === 'ok') {
                this.ok = true
            } else {
                this.error = tested
            }
            this.loading = false
        },
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
