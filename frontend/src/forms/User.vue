<template>
    <form @submit="saveUser">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Username</label>
            <div class="col-6 col-md-4">
                <input v-model="user.username" type="text" class="form-control" placeholder="Username" required autocorrect="off" autocapitalize="none">
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
                <input v-model="user.email" type="email" name="email" class="form-control" id="email" value="" placeholder="user@domain.com" required autocapitalize="none" spellcheck="false">
            </div>
        </div>
        <div class="form-group row">
            <label for="password" class="col-sm-4 col-form-label">Password</label>
            <div class="col-sm-8">
                <input v-model="user.password" type="password" name="password" class="form-control" id="password"  placeholder="Password" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="password_confirm" class="col-sm-4 col-form-label">Confirm Password</label>
            <div class="col-sm-8">
                <input v-model="user.confirm_password" type="password" name="password_confirm" class="form-control" id="password_confirm"  placeholder="Confirm Password" required>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button @click="saveUser" class="btn btn-primary btn-block">Create User</button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
  import Api from "../components/API";

  export default {
  name: 'FormUser',
  props: {

  },
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
  mounted() {

  },
  computed() {

  },
  methods: {
    async saveUser(e) {
      e.preventDefault();
      let user = this.user
      delete user.confirm_password
      await Api.user_create(user)
      const users = await Api.users()
      this.$store.commit('setUsers', users)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
