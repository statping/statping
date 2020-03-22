<template>
    <div class="col-12">
        <div class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">Services
                <router-link v-if="$store.state.admin" to="/dashboard/create_service" class="btn btn-sm btn-outline-success float-right">
                <font-awesome-icon icon="plus"/>  Create
            </router-link></div>
            <div class="card-body">
                <ServicesList/>
            </div>
        </div>

        <div class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">Groups</div>
            <div class="card-body">
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
            <tr v-for="(group, index) in $store.getters.groupsCleanInOrder" v-bind:key="group.id">
                <td><span class="drag_icon d-none d-md-inline">
                    <font-awesome-icon icon="bars" class="mr-3" /></span> {{group.name}}
                </td>
                <td>{{$store.getters.servicesInGroup(group.id).length}}</td>
                <td>
                    <span v-if="group.public" class="badge badge-primary">PUBLIC</span>
                    <span v-if="!group.public" class="badge badge-secondary">PRIVATE</span>
                </td>
                <td class="text-right">
                    <div v-if="$store.state.admin" class="btn-group">
                        <a @click.prevent="editGroup(group, edit)" href="#" class="btn btn-outline-secondary"><font-awesome-icon icon="chart-area" /> Edit</a>
                        <a @click.prevent="deleteGroup(group)" href="#" class="btn btn-danger">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>

                </draggable>
        </table>

            </div>
        </div>


        <FormGroup v-if="$store.state.admin" :edit="editChange" :in_group="group"/>

    </div>

</template>

<script>
  import FormGroup from "../../forms/Group";
  import Api from "../../API";
  import ToggleSwitch from "../../forms/ToggleSwitch";
  import draggable from 'vuedraggable'
  import ServicesList from './ServicesList';

  export default {
      name: 'DashboardServices',
      components: {
          ServicesList,
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
          groupsList: {
              get() {
                  return this.$store.getters.groupsCleanInOrder
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
          async deleteGroup(g) {
              let c = confirm(`Are you sure you want to delete '${g.name}'?`)
              if (c) {
                  await Api.group_delete(g.id)
                  const groups = await Api.groups()
                  this.$store.commit('setGroups', groups)
              }
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
