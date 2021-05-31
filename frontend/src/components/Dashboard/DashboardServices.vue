<template>
    <div class="col-12">

        <div class="card contain-card mb-4">
            <div class="card-header">{{ $t('services') }}
                <router-link v-if="$store.state.admin" to="/dashboard/create_service" class="btn btn-sm btn-success float-right">
                    <font-awesome-icon icon="plus"/>  {{$t('create')}}
                </router-link>
            </div>
            <div class="card-body pt-0">
                <ServicesList/>
            </div>
        </div>

        <div class="card contain-card mb-4">
            <div class="card-header">{{ $t('groups') }}</div>
            <div class="card-body pt-0">

                <div v-if="groupsList.length === 0">
                    <div class="alert alert-dark d-block mt-3 mb-0">
                        You currently don't have any groups! Create one using the form below.
                    </div>
                </div>

        <table v-else class="table">
            <thead>
            <tr>
                <th scope="col">{{ $t('name') }}</th>
                <th scope="col" class="d-none d-md-table-cell">{{ $tc('service', 2) }}</th>
                <th scope="col">{{ $t('visibility') }}</th>
                <th scope="col"></th>
            </tr>
            </thead>

            <draggable tag="tbody" v-model="groupsList" class="sortable_groups" handle=".drag_icon">
            <tr v-for="(group, index) in groupsList" v-bind:key="group.id">
                <td><span class="drag_icon d-none d-md-inline">
                    <font-awesome-icon icon="bars" class="mr-3" /></span> {{group.name}}
                </td>
                <td class="d-none d-md-table-cell">{{$store.getters.servicesInGroup(group.id).length}}</td>
                <td>
                    <span class="badge text-uppercase" :class="{'badge-primary': group.public, 'badge-secondary': !group.public}">
                        {{group.public ? $t('public') : $t('private')}}
                    </span>
                </td>
                <td class="text-right">
                    <div v-if="$store.state.admin" class="btn-group">
                        <button @click.prevent="editGroup(group, edit)" href="#" class="btn btn-sm btn-outline-secondary">
                            <font-awesome-icon icon="edit" />
                        </button>
                        <button @click.prevent="deleteGroup(group)" href="#" class="btn btn-sm btn-danger">
                            <font-awesome-icon icon="times" />
                        </button>
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
  const Modal = () => import(/* webpackChunkName: "dashboard" */ "@/components/Elements/Modal")
  const FormGroup = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Group')
  const ToggleSwitch = () => import(/* webpackChunkName: "dashboard" */ '@/forms/ToggleSwitch')
  const ServicesList = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/ServicesList')
  import Api from "../../API";
  const draggable = () => import(/* webpackChunkName: "dashboard" */ 'vuedraggable')

  export default {
      name: 'DashboardServices',
      components: {
        Modal,
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
      methods: {
          editChange(v) {
              this.group = {}
              this.edit = v
          },
          editGroup(g, mode) {
              this.group = g
              this.edit = !mode
          },
        confirm_delete(service) {

        },
        async delete(g) {
          await Api.group_delete(g.id)
          const groups = await Api.groups()
          this.$store.commit('setGroups', groups)
        },
          async deleteGroup(g) {
            const modal = {
              visible: true,
              title: "Delete Group",
              body: `Are you sure you want to delete group ${g.name}? All services attached will be removed from this group.`,
              btnColor: "btn-danger",
              btnText: "Delete Group",
              func: () => this.delete(g),
            }
            this.$store.commit("setModal", modal)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
