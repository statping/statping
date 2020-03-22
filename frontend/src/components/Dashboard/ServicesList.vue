<template>
    <table class="table">
        <thead>
        <tr>
            <th scope="col">Name</th>
            <th scope="col" class="d-none d-md-table-cell"></th>
            <th scope="col" class="d-none d-md-table-cell">Visibility</th>
            <th scope="col" class="d-none d-md-table-cell">Group</th>
            <th scope="col"></th>
        </tr>
        </thead>
        <draggable tag="tbody" v-model="servicesList" handle=".drag_icon">
            <tr v-for="(service, index) in $store.getters.servicesInOrder" :key="service.id">
                <td>
                    <span v-if="$store.state.admin" class="drag_icon d-none d-md-inline">
                        <font-awesome-icon icon="bars" class="mr-3"/>
                    </span> {{service.name}}
                </td>
                <td v-if="$store.state.admin" class="d-none d-md-table-cell">
                    <ToggleSwitch v-if="service.online" :service="service"/>
                </td>
                <td class="d-none d-md-table-cell">
                    <span class="badge" :class="{'badge-primary': service.public, 'badge-secondary': !service.public}">
                        {{service.public ? "PUBLIC" : "PRIVATE"}}
                    </span>
                </td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.group_id !== 0">
                        <span class="badge badge-secondary">{{serviceGroup(service)}}</span>
                    </div>
                </td>
                <td class="text-right">
                    <div class="btn-group">
                        <router-link v-if="$store.state.admin" :to="{path: `/dashboard/edit_service/${service.id}`, params: {service: service} }" class="btn btn-outline-secondary">
                            <i class="fas fa-chart-area"></i> Edit
                        </router-link>
                        <router-link :to="{path: serviceLink(service), params: {service: service} }" class="btn btn-outline-secondary">
                            <i class="fas fa-chart-area"></i> View
                        </router-link>
                        <a v-if="$store.state.admin" @click.prevent="deleteService(service)" href="#" class="btn btn-danger">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>
        </draggable>
    </table>
</template>

<script>
import Api from "../../API";
import draggable from 'vuedraggable'
import ToggleSwitch from '../../forms/ToggleSwitch';

export default {
      name: 'ServicesList',
    components: {
        ToggleSwitch,
          draggable
    },
    computed: {
        servicesList: {
            get () {
                return this.$store.getters.servicesInOrder
            },
            set (value) {
                this.updateOrder(value)
            }
        }
    },
      data() {
          return {

          }
      },
      methods: {
          async updateOrder(value) {
              let data = [];
              value.forEach((s, k) => {
                  data.push({ service: s.id, order: k + 1 })
              });
              const reorder = await Api.services_reorder(data)
              window.console.log('reorder', reorder)
              const services = await Api.services()
              this.$store.commit('setServices', services)
          },
          async deleteService(s) {
              let c = confirm(`Are you sure you want to delete '${s.name}'?`)
              if (c) {
                  await Api.service_delete(s.id)
                  const services = await Api.services()
                  this.$store.commit('setServices', services)
              }
          },
          serviceGroup(s) {
              let group = this.$store.getters.groupById(s.group_id)
              if (group) {
                  return group.name
              }
              return ""
          },
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
