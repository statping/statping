<template>
    <form @submit="saveUser">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Username</label>
            <div class="col-6 col-md-4">
                <input v-model="user.username" type="text" class="form-control" placeholder="Username" required autocorrect="off" autocapitalize="none">
            </div>
            <div class="col-6 col-md-4">
                          <span class="switch">
                            <input type="checkbox" name="admin" class="switch" id="switch-normal">
                            <label for="switch-normal">Administrator</label>
                            <input v-model="user.admin" type="hidden" name="admin" id="switch-normal-value" value="false">
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
  methods: {
    saveUser () {

    },
    async saveGroup(e) {
      e.preventDefault();
      const data = {name: this.group.name, public: this.group.public}
      await Api.group_create(data)
      const groups = await Api.groups()
      this.$store.commit('setGroups', groups)
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
