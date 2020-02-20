<template>
    <div>
    <div class="col-12">
        <h1 class="text-black-50">Services
            <router-link to="/dashboard/create_service" class="btn btn-outline-success mt-1 float-right">
                <i class="fas fa-plus"></i> Create
            </router-link>
        </h1>

        <table class="table">
            <thead>
            <tr>
                <th scope="col">Name</th>
                <th scope="col" class="d-none d-md-table-cell">Status</th>
                <th scope="col" class="d-none d-md-table-cell">Visibility</th>
                <th scope="col" class="d-none d-md-table-cell">Group</th>
                <th scope="col"></th>
            </tr>
            </thead>
            <draggable tag="tbody" v-model="servicesList" :key="$store.getters.servicesInOrder.length" class="sortable" handle=".drag_icon">
            <tr v-for="(service, index) in $store.getters.servicesInOrder" :key="index">
                <td>
                    <span class="drag_icon d-none d-md-inline">
                        <font-awesome-icon icon="bars" />
                    </span> {{service.name}}
                </td>
                <td class="d-none d-md-table-cell">
                    <span class="badge" :class="{'animate-fader': !service.online, 'badge-success': service.online, 'badge-danger': !service.online}">
                        {{service.online ? "ONLINE" : "OFFLINE"}}
                    </span>
                    <ToggleSwitch v-if="service.online" :service="service"/>
                </td>
                <td class="d-none d-md-table-cell">
                    <span class="badge" :class="{'badge-primary': service.public, 'badge-secondary': !service.public}">
                        {{service.public ? "PUBLIC" : "PRIVATE"}}
                    </span>
                </td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.group_id !== 0"><span class="badge badge-secondary">{{serviceGroup(service)}}</span></div>
                </td>
                <td class="text-right">
                    <div class="btn-group">
                        <router-link :to="{path: `/dashboard/edit_service/${service.id}`, params: {service: service} }" class="btn btn-outline-secondary">
                            <i class="fas fa-chart-area"></i> Edit
                        </router-link>
                        <router-link :to="{path: serviceLink(service), params: {service: service} }" class="btn btn-outline-secondary">
                            <i class="fas fa-chart-area"></i> View
                        </router-link>
                        <a @click.prevent="deleteService(service)" href="#" class="btn btn-danger">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>
            </draggable>
        </table>

    </div>

    <div class="col-12 mt-5">

        <h1 class="text-muted">Groups</h1>
        <table class="table">
            <thead>
            <tr>
                <th scope="col">Name</th>
                <th scope="col">Services</th>
                <th scope="col">Visibility</th>
                <th scope="col"></th>
            </tr>
            </thead>

            <draggable tag="tbody" v-model="groupsList" class="sortable_groups" handle=".drag_icon">
            <tr v-for="(group, index) in $store.getters.groupsClean" v-bind:key="index">
                <td><span class="drag_icon d-none d-md-inline"><font-awesome-icon icon="bars" /></span> {{group.name}}</td>
                <td>{{$store.getters.servicesInGroup(group.id).length}}</td>
                <td>
                    <span v-if="group.public" class="badge badge-primary">PUBLIC</span>
                    <span v-if="!group.public" class="badge badge-secondary">PRIVATE</span>
                </td>
                <td class="text-right">
                    <div class="btn-group">
                        <a @click.prevent="editGroup(group, edit)" href="#" class="btn btn-outline-secondary"><font-awesome-icon icon="chart-area" /> Edit</a>
                        <a @click.prevent="deleteGroup(group)" href="#" class="btn btn-danger">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>

                </draggable>
        </table>

        <FormGroup :edit="editChange" :in_group="group"/>

    </div>
    </div>
</template>

<script>
  import FormGroup from "../../forms/Group";
  import Api from "../../API";
  import ToggleSwitch from "../../forms/ToggleSwitch";
  import draggable from 'vuedraggable'

  export default {
      name: 'DashboardServices',
      components: {
          ToggleSwitch,
          FormGroup,
          draggable
      },
      data() {
          return {
              edit: false,
              group: {}
          }
      },
      computed: {
          servicesList: {
              get() {
                  return this.$store.state.servicesInOrder
              },
              async set(value) {
                  let data = [];
                  value.forEach((s, k) => {
                      data.push({service: s.id, order: k + 1})
                  });
                  await Api.services_reorder(data)
                  const services = await Api.services()
                  this.$store.commit('setServices', services)
              }
          },
          groupsList: {
              get() {
                  return this.$store.state.groupsInOrder
              },
              async set(value) {
                  let data = [];
                  value.forEach((s, k) => {
                      data.push({group: s.id, order: k + 1})
                  });
                  await Api.groups_reorder(data)
                  const groups = await Api.groups()
                  this.$store.commit('setGroups', groups)
              }
          }
      },
      beforeMount() {

      },
      methods: {
          editChange(v) {
              this.group = {}
              this.edit = v
          },
          editGroup(g, mode) {
              this.group = g
              this.edit = !mode
          },
          reordered_services() {

          },
          saveUpdatedOrder: function (e) {
              window.console.log("saving...");
              window.console.log(this.myViews.array()); // this.myViews.array is not a function
          },
          serviceGroup(s) {
              let group = this.$store.getters.groupById(s.group_id)
              if (group) {
                  return group.name
              }
              return ""
          },
          async deleteGroup(g) {
              let c = confirm(`Are you sure you want to delete '${g.name}'?`)
              if (c) {
                  await Api.group_delete(g.id)
                  const groups = await Api.groups()
                  this.$store.commit('setGroups', groups)
              }
          },
          async deleteService(s) {
              let c = confirm(`Are you sure you want to delete '${s.name}'?`)
              if (c) {
                  await Api.service_delete(s.id)
                  const services = await Api.services()
                  this.$store.commit('setServices', services)
              }
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
