<template>
    <form @submit="updateIntegration">
        <h4 class="text-capitalize">{{integration.full_name}}</h4>
        <p class="small text-muted" v-html="integration.description"></p>

        <div v-for="(field, index) in integration.fields" v-bind:key="index" class="form-group">

            <label class="text-capitalize">{{field.name}}</label>

            <textarea v-if="field.type === 'textarea'" v-model="field.value" rows="3" class="form-control"></textarea>

            <input v-else :type="field.type" v-model="field.value" class="form-control">

            <small class="form-text text-muted" v-html="field.description"></small>
        </div>

        <div class="col-12">
            <div class="col-3">
                <span @click="integration.enabled = !!integration.enabled" class="switch">
                    <input type="checkbox" name="enabled-option" class="switch" v-model="integration.enabled" v-bind:id="`switch-${integration.name}`" v-bind:checked="integration.enabled">
                    <label v-bind:for="`switch-${integration.name}`"></label>
                </span>
            </div>

            <div v-if="services.length !== 0" class="col-12">
                <table class="table">
                    <thead>
                    <tr>
                        <th scope="col">Name</th>
                        <th scope="col">Domain</th>
                        <th scope="col">Port</th>
                        <th scope="col">Interval</th>
                        <th scope="col">Timeout</th>
                        <th scope="col">Type</th>
                        <th scope="col"></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="(service, index) in services" v-bind:key="index">
                        <td><input v-model="service.name" type="text" style="width: 80pt"></td>
                        <td>{{service.domain}}</td>
                        <td>{{service.port}}</td>
                        <td><input v-model="service.check_interval" type="number" style="width: 35pt"></td>
                        <td><input v-model="service.timeout" type="number" style="width: 35pt"></td>
                        <td>{{service.type}}</td>
                        <td><button @click.prevent="addService(service)" v-bind:disabled="service.added" :disabled="service.added" class="btn btn-sm btn-outline-primary">Add</button></td>
                    </tr>
                    </tbody>
                </table>
            </div>

            <div class="col-12">
                <button @click.prevent="updateIntegration" type="submit" class="btn btn-block btn-info">Fetch Services</button>
            </div>
        </div>

        <div class="alert alert-danger d-none" role="alert"></div>
    </form>
</template>

<script>
  import Api from "../API";

  export default {
      name: 'FormIntegration',
      props: {
          integration: {
              type: Object
          }
      },
      data() {
          return {
              out: {},
              services: []
          }
      },
      watch: {},
      methods: {
          async addService(s) {
              const data = {
                  name: s.name,
                  type: s.type,
                  domain: s.domain,
                  port: s.port,
                  check_interval: s.check_interval,
                  timeout: s.timeout
              }
              const out = await Api.service_create(data)
              const services = await Api.services()
              this.$store.commit('setServices', services)
              s.added = true
          },
          async updateIntegration() {
              const i = this.integration
              const data = {name: i.name, enabled: i.enabled, fields: i.fields}
              this.out = data
              const out = await Api.integration_save(data)
              if (out != null) {
                  this.services = out
              }
              const integrations = await Api.integrations()
              this.$store.commit('setIntegrations', integrations)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
