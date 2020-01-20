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

            <tr v-for="(user, index) in $store.getters.users" v-bind:key="index" >
                <td>{{user.username}}</td>
                <td class="text-right">
                    <div class="btn-group"><font-awesome-icon icon="user-edit" />
                        <a href="user/1" class="btn btn-outline-secondary"><i class="fas fa-user-edit"></i> Edit</a>
                        <a @click="deleteUser(user)" href="#" class="btn btn-danger"><font-awesome-icon icon="times" /></a>
                    </div>
                </td>
            </tr>
            </tbody>
        </table>

        <h1 class="text-black-50 mt-5">Create User</h1>

        <div class="card">
            <div class="card-body">

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
      user: {
        username: "",
        admin: false,
        email: "",
        password: "",
        confirm_password: ""
      }
    }
  },
  created() {

  },
  methods: {
    async saveUser(e) {
      e.preventDefault();
      let u = this.user;
      if (u.password === u.confirm_password) {
        alert("Both password do not match")
        return
      }
      const data = {name: this.group.name, public: this.group.public}
      await Api.user_create(data)
      const users = await Api.users()
      this.$store.commit('setUsers', users)
    },
    async deleteUser(u) {
      let c = confirm(`Are you sure you want to delete user '${u.username}'?`)
      if (c) {
        await Api.user_delete(u.id)
        const users = await Api.users()
        this.$store.commit('setUsers', users)
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
