<template>
    <div class="card contain-card text-black-50 bg-white mb-3">
        <div class="card-header"> {{user.id ? `Update ${user.username}` : "Create User"}}
            <transition name="slide-fade">
                <button @click.prevent="removeEdit" v-if="user.id" class="btn btn-sm float-right btn-danger btn-sm">Close</button>
            </transition>
        </div>
        <div class="card-body">
    <form @submit="saveUser">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Username</label>
            <div class="col-6 col-md-4">
                <input v-model="user.username" type="text" class="form-control" placeholder="Username" required autocorrect="off" autocapitalize="none" v-bind:readonly="user.id">
            </div>
            <div class="col-6 col-md-4">
                  <span @click="user.admin = !!user.admin" class="switch">
                    <input v-model="user.admin" type="checkbox" class="switch" id="switch-normal" v-bind:checked="user.admin">
                    <label for="switch-normal">Administrator</label>
                  </span>
            </div>
        </div>
        <div class="form-group row">
            <label for="email" class="col-sm-4 col-form-label">Email Address</label>
            <div class="col-sm-8">
                <input v-model="user.email" type="email" class="form-control" id="email" placeholder="user@domain.com" required autocapitalize="none" spellcheck="false">
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Password</label>
            <div class="col-sm-8">
                <input v-model="user.password" type="password" class="form-control" placeholder="Password" required>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Confirm Password</label>
            <div class="col-sm-8">
                <input v-model="user.confirm_password" type="password" class="form-control" placeholder="Confirm Password" required>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button @click="saveUser"
                        :disabled="loading || !user.username || !user.email || !user.password || !user.confirm_password || (user.password !== user.confirm_password)"
                        class="btn btn-block" :class="{'btn-primary': !user.id, 'btn-secondary': user.id}">
                    {{loading ? "Loading..." : user.id ? "Update User" : "Create User"}}
                </button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
    </div>
    </div>
</template>

<script>
  import Api from "../API";

  export default {
  name: 'FormUser',
  props: {
    in_user: {
      type: Object
    },
    edit: {
      type: Function
    }
  },
  data () {
    return {
      loading: false,
      user: {
        username: "",
        admin: false,
        email: "",
        password: "",
        confirm_password: ""
      }
    }
  },
  watch: {
    in_user() {
        const u = this.in_user
        this.user = u
    }
  },
  methods: {
    removeEdit() {
      this.user = {}
      this.edit(false)
    },
    async saveUser(e) {
      e.preventDefault();
      this.loading = true
      if (this.user.id) {
        await this.updateUser()
      } else {
        await this.createUser()
      }
        this.loading = false
    },
    async createUser() {
      let user = this.user
      delete user.confirm_password
      await Api.user_create(user)
      const users = await Api.users()
      this.$store.commit('setUsers', users)
      this.user = {}
    },
    async updateUser() {
      let user = this.user
      if (!user.password) {
        delete user.password
      }
      delete user.confirm_password
      await Api.user_update(user)
      const users = await Api.users()
      this.$store.commit('setUsers', users)
      this.edit(false)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
