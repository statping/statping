<template>
    <div class="container col-md-7 col-sm-12 mt-md-5 bg-light">
        <div class="col-10 offset-1 col-md-8 offset-md-2 mt-md-2">
            <div class="col-12 col-md-8 offset-md-2 mb-4">
                <img class="col-12 mt-5 mt-md-0" src="/public/banner.png">
            </div>

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
                        <button @click="login" type="submit" class="btn btn-block mb-3" :class="{'btn-primary': !loading, 'btn-default': loading}" v-bind:disabled="loading">
                            {{loading ? "Loading" : "Sign in"}}
                        </button>
                        <div v-if="error" class="alert alert-danger" role="alert">
                            Incorrect username or password
                        </div>
                    </div>
                </div>
            </form>

        </div>
    </div>
</template>

<script>
  import Api from "../components/API";

  export default {
  name: 'Login',
  components: {
  },
  data () {
    return {
      username: "",
      password: "",
      auth: {},
      loading: false,
      error: false
    }
  },
  methods: {
    async login (e) {
      e.preventDefault();
      this.loading = true
      this.error = false
      const auth = await Api.login(this.username, this.password)
      if (auth.error) {
        this.error = true
      } else if (auth.token) {
        this.auth = Api.saveToken(this.username, auth.token)
        await this.$store.dispatch('loadRequired')
        this.$router.push('/dashboard')
      }
      this.loading = false
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
