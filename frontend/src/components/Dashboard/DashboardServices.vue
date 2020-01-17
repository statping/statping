<!--
  - Statup
  - Copyright (C) 2020.  Hunter Long and the project contributors
  - Written by Hunter Long <info@socialeck.com> and the project contributors
  -
  - https://github.com/hunterlong/statup
  -
  - The licenses for most software and other practical works are designed
  - to take away your freedom to share and change the works.  By contrast,
  - the GNU General Public License is intended to guarantee your freedom to
  - share and change all versions of a program--to make sure it remains free
  - software for all its users.
  -
  - You should have received a copy of the GNU General Public License
  - along with this program.  If not, see <http://www.gnu.org/licenses/>.
  -->

<template>
    <div>
    <div class="col-12">
        <h1 class="text-black-50">Services
            <router-link to="/service/create" class="btn btn-outline-success mt-1 float-right"><i class="fas fa-plus"></i> Create</router-link>
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
            <tbody class="sortable" id="services_table">

            <tr v-for="(service, index) in $store.getters.services" v-bind:key="index">
                <td><span class="drag_icon d-none d-md-inline"><font-awesome-icon icon="bars" /></span> {{service.name}}</td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.online"><span class="badge badge-success">ONLINE</span><font-awesome-icon class="toggle-service text-success" icon="toggle-on" /></div>
                    <div v-if="!service.online"><span class="badge badge-success">OFFLINE</span><font-awesome-icon class="toggle-service text-danger" icon="toggle-off" /></div>
                </td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.public"><span class="badge badge-primary">PUBLIC</span></div>
                    <div v-if="!service.public"><span class="badge badge-secondary">PRIVATE</span></div>
                </td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.group_id !== 0"><span class="badge badge-secondary">{{$store.getters.groupById(service.group_id).name}}</span></div>
                </td>
                <td class="text-right">
                    <div class="btn-group">
                        <router-link :to="{path: `/service/${service.id}`, params: {service: service} }" class="btn btn-outline-secondary"><i class="fas fa-chart-area"></i> View</router-link>
                        <a href="api/services/1" class="ajax_delete btn btn-danger" data-method="DELETE" data-obj="service_1" data-id="1">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>

            </tbody>
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
            <tbody class="sortable_groups" id="groups_table">

            <tr v-for="(group, index) in $store.getters.cleanGroups()" v-bind:key="index">
                <td><span class="drag_icon d-none d-md-inline"><font-awesome-icon icon="bars" /></span> {{group.name}}</td>
                <td>{{$store.getters.servicesInGroup(group.id).length}}</td>
                <td>
                    <span v-if="group.public" class="badge badge-primary">PUBLIC</span>
                    <span v-if="!group.public" class="badge badge-secondary">PRIVATE</span>
                </td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="group/2" class="btn btn-outline-secondary"> <font-awesome-icon icon="chart-area" /> Edit</a>
                        <a href="api/groups/2" class="btn btn-danger">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>

            </tbody>
        </table>

        <h1 class="text-muted mt-5">Create Group</h1>

        <div class="card">
            <div class="card-body">
                <FormGroup :group="null"/>
            </div>
        </div>

    </div>
    </div>
</template>

<script>
import FormGroup from "../../forms/Group";

export default {
  name: 'DashboardServices',
  components: {FormGroup},
  data () {
    return {

    }
  },
  beforeMount() {

  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
