<template>
    <div class="card contain-card mb-3">
        <div class="card-header"> {{user.id ? `${$t('update')} ${user.username}` : $t('user_create')}}
            <transition name="slide-fade">
                <button @click.prevent="removeEdit" v-if="user.id" class="btn btn-sm float-right btn-danger btn-sm">Close</button>
            </transition>
        </div>
        <div class="card-body">
    <form @submit="saveUser">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{$t('username')}}</label>
            <div class="col-6 col-md-4">
                <input v-model="user.username" type="text" class="form-control" id="username" placeholder="Username" required autocorrect="off" autocapitalize="none" v-bind:readonly="user.id">
            </div>
            <div class="col-6 col-md-4">
                  <span id="admin_switch" @click="user.admin = !!user.admin" class="switch">
                    <input v-model="user.admin" type="checkbox" class="switch" id="user_admin_switch" v-bind:checked="user.admin">
                    <label for="user_admin_switch">{{$t('administrator')}}</label>
                  </span>
            </div>
        </div>
        <div class="form-group row">
            <label for="email" class="col-sm-4 col-form-label">{{$t('email')}}</label>
            <div class="col-sm-8">
                <input v-model="user.email" type="email" class="form-control" id="email" placeholder="user@domain.com" required autocapitalize="none" spellcheck="false">
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{$t('password')}}</label>
            <div class="col-sm-8">
                <input v-model="user.password" type="password" id="password" class="form-control" placeholder="Password" required>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{$t('confirm_password')}}</label>
            <div class="col-sm-8">
                <input v-model="user.confirm_password" type="password" id="password_confirm" class="form-control" placeholder="Confirm Password" required>
            </div>
        </div>
        <div v-if="user.api_key" class="form-group row">
            <label for="user_key_key" class="col-sm-4 col-form-label">API Key</label>
            <div class="col-sm-8">
                <div class="input-group">
                    <input v-bind:value="user.api_key" type="text" class="form-control" id="user_key_key" readonly>
                    <div class="input-group-append copy-btn">
                        <button @click.prevent="copy(user.api_key)" class="btn btn-outline-secondary" type="button">Copy</button>
                    </div>
                </div>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <LoadButton
                        class="btn-primary"
                        :disabled="loading || !user.username || !user.email || !user.password || !user.confirm_password || (user.password !== user.confirm_password)"
                        :action="saveUser"
                        :label="user.id ? $t('user_update'): $t('user_create')"
                />
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
    </div>
    </div>
</template>

<script>
  import Api from "../API";
  const LoadButton = () => import(/* webpackChunkName: "index" */ "@/components/Elements/LoadButton");

  export default {
  name: 'FormUser',
    components: {LoadButton},
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
        confirm_password: "",
        api_key: "",
      }
    }
  },
  watch: {
    in_user() {
        this.user = this.in_user
    }
  },
  methods: {
    removeEdit() {
      this.user = {}
      this.edit(false)
    },
    async saveUser() {
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
      await this.update()
      this.user = {}
    },
    async updateUser() {
      let user = this.user
      if (!user.password) {
        delete user.password
      }
      delete user.confirm_password
      await Api.user_update(user)
      await this.update()
      this.edit(false)
    },
    async update() {
      const users = await Api.users()
      this.$store.commit('setUsers', users)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
