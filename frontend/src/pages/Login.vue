<template>
    <div class="offset-md-3 offset-lg-4 offset-0 col-lg-4 col-md-6 mt-5">

      <div class="offset-1 offset-lg-2 col-lg-8 col-10 mb-4 mb-md-3">
        <img alt="Statping Login" class="embed-responsive" src="banner.png">
      </div>

      <div class="login_container col-12 p-4">
            <FormLogin/>
        </div>
    </div>
</template>

<script>
  const FormLogin = () => import(/* webpackChunkName: "index" */ '@/forms/Login')

  export default {
  name: 'Login',
  components: {
      FormLogin
  },
  data () {
    return {

    }
  },
    mounted() {

    },
    methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
