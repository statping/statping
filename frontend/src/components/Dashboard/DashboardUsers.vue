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
                    <div class="btn-group">
                        <a @click.prevent="editUser(user, edit)" href="" class="btn btn-outline-secondary"><font-awesome-icon icon="user" /> Edit</a>
                        <a v-if="index !== 0" @click.prevent="deleteUser(user)" href="" class="btn btn-danger"><font-awesome-icon icon="times" /></a>
                    </div>
                </td>
            </tr>
            </tbody>
        </table>

        <FormUser :edit="editChange" :in_user="user"/>
    </div>
</template>

<script>
  import Api from "../API"
  import FormUser from "../../forms/User";

  export default {
  name: 'DashboardUsers',
    components: {FormUser},
    data () {
    return {
      edit: false,
      user: {}
    }
  },
  methods: {
    editChange(v) {
      this.user = {}
      this.edit = v
    },
    editUser(u, mode) {
      this.user = u
      this.edit = !mode
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
