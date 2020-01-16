<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <div class="col-10 offset-1 col-md-8 offset-md-2 mt-md-2">
            <div class="col-12 col-md-8 offset-md-2 mb-4">
                <img class="col-12 mt-5 mt-md-0" :src="require(`@/assets/banner.png`)">
            </div>
            {{auth}}
            <form id="login_form" @submit="login" method="post">
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
                        <button type="submit" class="btn btn-primary btn-block mb-3">Sign in</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
  import Api from "../components/API"

  export default {
  name: 'Login',
  components: {

  },
  data () {
    return {
      username: "",
      password: "",
      auth: null
    }
  },
  methods: {
    async login (e) {
      e.preventDefault();
      const auth = await Api.login(this.username, this.password)
      if (auth.token !== null) {
        this.auth = Api.saveToken(this.username, auth.token)
        await this.$router.push('/dashboard')
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
