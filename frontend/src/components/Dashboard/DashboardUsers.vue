<template>
    <div class="col-12">
        <div class="card contain-card text-black-50 bg-white mb-4">
            <div class="card-header">{{ $t('top_nav.users') }}</div>
            <div class="card-body pt-0">
        <table class="table table-striped">
            <thead>
                <tr>
                    <th scope="col">{{$t('username')}}</th>
                    <th scope="col">{{$t('type')}}</th>
                    <th scope="col" class="d-none d-md-table-cell">{{ $t('last_login') }}</th>
                    <th scope="col"></th>
                </tr>
            </thead>
            <tbody id="users_table">

            <tr v-for="(user, index) in users" v-bind:key="user.id" >
                <td>{{user.username}}</td>
                <td>
                    <span class="badge text-uppercase" :class="{'badge-danger': user.admin, 'badge-primary': !user.admin}">
                        {{user.admin ? $t('admin') : $t('user')}}
                    </span>
                </td>
                <td class="d-none d-md-table-cell">{{niceDate(user.updated_at)}}</td>
                <td class="text-right">
                    <div class="btn-group">
                        <a @click.prevent="editUser(user, edit)" href="#" class="btn btn-outline-secondary edit-user">
                            <font-awesome-icon icon="user" /> {{$t('edit') }}
                        </a>
                        <a @click.prevent="deleteUser(user)" v-if="index !== 0" href="#" class="btn btn-danger delete-user">
                            <font-awesome-icon icon="times" />
                        </a>
                    </div>
                </td>
            </tr>
            </tbody>
        </table>
            </div>
        </div>

                <FormUser v-if="$store.state.admin" :edit="editChange" :in_user="user"/>

    </div>
</template>

<script>
  import Api from "../../API"
  const FormUser = () => import('@/forms/User')

  export default {
  name: 'DashboardUsers',
    components: {FormUser},
    data () {
    return {
      edit: false,
      user: {}
    }
  },
      computed: {
        users() {
            return this.$store.getters.users
        }
      },
  methods: {
    editChange(v) {
      this.user = {}
      this.edit = v
    },
    editUser(u, mode) {
      delete(u.password)
      delete(u.confirm_password)
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
