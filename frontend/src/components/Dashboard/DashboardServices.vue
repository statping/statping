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
        <h1 class="text-black-50">Services <a href="service/create" class="btn btn-outline-success mt-1 float-right">
            <i class="fas fa-plus"></i> Create</a>
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

            <tr v-for="(service, index) in services" v-bind:key="index">
                <td><span class="drag_icon d-none d-md-inline"><i class="fas fa-bars"></i></span> {{service.name}}</td>
                <td class="d-none d-md-table-cell"><span class="badge badge-success">ONLINE</span>
                    <i class="toggle-service fas fa-toggle-on text-success" data-online="true" data-id="1"></i>
                </td>
                <td class="d-none d-md-table-cell"><span class="badge badge-primary">PUBLIC</span></td>
                <td class="d-none d-md-table-cell"><span class="badge badge-secondary">Main Services</span></td>
                <td class="text-right">
                    <div class="btn-group">
                        <router-link :to="{path: `/service/${service.id}`, params: {service: service} }" class="btn btn-outline-secondary"><i class="fas fa-chart-area"></i> View</router-link>
                        <a href="api/services/1" class="ajax_delete btn btn-danger" data-method="DELETE" data-obj="service_1" data-id="1"><i class="fas fa-times"></i></a>
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

            <tr v-for="(group, index) in groups" v-bind:key="index">
                <td><span class="drag_icon d-none d-md-inline"><i class="fas fa-bars"></i></span>{{group.name}}</td>
                <td></td>
                <td><span class="badge badge-secondary">PRIVATE</span></td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="group/2" class="btn btn-outline-secondary"><i class="fas fa-chart-area"></i> Edit</a>
                        <a href="api/groups/2" class="ajax_delete btn btn-danger" data-method="DELETE" data-obj="group_2" data-id="2"><i class="fas fa-times"></i></a>
                    </div>
                </td>
            </tr>

            </tbody>
        </table>


        <h1 class="text-muted mt-5">Create Group</h1>

        <div class="card">
            <div class="card-body">

                <form class="ajax_form" action="api/groups" data-redirect="services" method="POST">
                    <div class="form-group row">
                        <label for="title" class="col-sm-4 col-form-label">Group Name</label>
                        <div class="col-sm-8">
                            <input type="text" name="name" class="form-control" value="" id="title" placeholder="Group Name" required>
                        </div>
                    </div>
                    <div class="form-group row">
                        <label for="switch-group-public" class="col-sm-4 col-form-label">Public Group</label>
                        <div class="col-8 mt-1">
            <span class="switch float-left">
                <input type="checkbox" name="public" class="switch" id="switch-group-public" >
                <label for="switch-group-public">Show group services to the public</label>
                <input type="hidden" name="public" id="switch-group-public-value" value="false">
            </span>
                        </div>
                    </div>
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <button type="submit" class="btn btn-primary btn-block">Create Group</button>
                        </div>
                    </div>
                    <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                </form>
            </div>
        </div>

    </div>
    </div>
</template>

<script>

export default {
  name: 'DashboardServices',
  props: {
    services: Array,
    groups: Array,
  },
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
