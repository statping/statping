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
    <div class="col-12">
        <h1 class="text-black-50">Users</h1>
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">Username</th>
                <th scope="col"></th>
            </tr>
            </thead>
            <tbody id="users_table">

            <tr v-for="(user, index) in users" v-bind:key="index" >
                <td>{{user.username}}</td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="user/1" class="btn btn-outline-secondary"><i class="fas fa-user-edit"></i> Edit</a>
                        <a href="api/users/1" class="ajax_delete btn btn-danger"><i class="fas fa-times"></i></a>
                    </div>
                </td>
            </tr>

            </tbody>
        </table>

        <h1 class="text-black-50 mt-5">Create User</h1>

        <div class="card">
            <div class="card-body">
                <form class="ajax_form" action="api/users" data-redirect="users" method="POST">
                    <div class="form-group row">
                        <label for="username" class="col-sm-4 col-form-label">Username</label>
                        <div class="col-6 col-md-4">
                            <input type="text" name="username" class="form-control" value="" id="username" placeholder="Username" required autocorrect="off" autocapitalize="none">
                        </div>
                        <div class="col-6 col-md-4">
                          <span class="switch">
                            <input type="checkbox" name="admin" class="switch" id="switch-normal">
                            <label for="switch-normal">Administrator</label>
                            <input type="hidden" name="admin" id="switch-normal-value" value="false">
                          </span>
                        </div>
                    </div>
                    <div class="form-group row">
                        <label for="email" class="col-sm-4 col-form-label">Email Address</label>
                        <div class="col-sm-8">
                            <input type="email" name="email" class="form-control" id="email" value="" placeholder="user@domain.com" required autocapitalize="none" spellcheck="false">
                        </div>
                    </div>
                    <div class="form-group row">
                        <label for="password" class="col-sm-4 col-form-label">Password</label>
                        <div class="col-sm-8">
                            <input type="password" name="password" class="form-control" id="password"  placeholder="Password" required>
                        </div>
                    </div>
                    <div class="form-group row">
                        <label for="password_confirm" class="col-sm-4 col-form-label">Confirm Password</label>
                        <div class="col-sm-8">
                            <input type="password" name="password_confirm" class="form-control" id="password_confirm"  placeholder="Confirm Password" required>
                        </div>
                    </div>
                    <div class="form-group row">
                        <div class="col-sm-12">
                            <button type="submit" class="btn btn-primary btn-block">Create User</button>
                        </div>
                    </div>
                    <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
  import Api from "../API"

  export default {
  name: 'DashboardUsers',
  data () {
    return {
        users: null
    }
  },
  created() {
    this.getUsers()
  },
  methods: {
    async getUsers () {
      this.users = await Api.users()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
