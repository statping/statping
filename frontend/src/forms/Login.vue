<template>
    <form @submit="login">
        <div class="form-group row">
            <label for="username" class="col-sm-2 col-form-label">Username</label>
            <div class="col-sm-10">
                <input type="text" v-model="username" name="username" class="form-control" id="username" placeholder="Username" autocorrect="off" autocapitalize="none">
            </div>
        </div>
        <div class="form-group row">
            <label for="password" class="col-sm-2 col-form-label">Password</label>
            <div class="col-sm-10">
                <input type="password" v-model="password" name="password" class="form-control" id="password" placeholder="Password">
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button v-on:click="login" type="submit" class="btn btn-primary btn-block mb-3">Sign in</button>
            </div>
        </div>
    </form>
</template>

<script>
  import Api from "../components/API";

  export default {
  name: 'FormLogin',
  props: {

  },
  data () {
    return {
        username: "",
        password: "",
        auth: null
      }
  },
  mounted() {

  },
  methods: {
    async login (e) {
      e.preventDefault();
      const auth = await Api.login(this.username, this.password)
      if (auth.token !== null) {
        this.auth = Api.saveToken(this.username, auth.token)
        this.$store.commit('setToken', auth)
        this.$router.push('/dashboard')
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
