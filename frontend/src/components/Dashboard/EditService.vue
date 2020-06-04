<template>

    <div v-if='ready'>

        <div v-if='errorCode==404' class="col-12">
           <div class="alert alert-warning" role="alert">
                Service {{ this.$route.params.id }} not found!
                <router-link v-if="$store.state.admin" to="/dashboard/create_service" class="btn btn-sm btn-outline-success float-right">
                    <font-awesome-icon icon="plus"/>  Create One?
                </router-link>
            </div>
        </div>

        <div v-else-if='errorCode==401' class="col-12">
           <div class="alert alert-danger" role="alert">
                Unauthorized! Perhaps your session has expired?
            </div>
        </div>

        <div v-else class="col-12">
            <FormService :in_service="service"/>
        </div>

    </div>
    <div v-else>

      <div class="text-center">
        <div class="spinner-border text-primary" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>

    </div>

</template>

<script>
  import Api from "@/API";
  import FormService from "@/forms/Service";

  export default {
  name: 'EditService',
  components: {
    FormService,
  },
  props: {

  },
  data () {
    return {
      ready: false,
      errorCode: 'none',
      service: {
          name: "",
          type: "http",
          domain: "",
          group_id: 0,
          method: "GET",
          post_data: "",
          headers: "",
          expected: "",
          expected_status: 200,
          port: 80,
          check_interval: 60,
          timeout: 15,
          permalink: "",
          order: 1,
          verify_ssl: true,
          redirect: true,
          allow_notifications: true,
          notify_all_changes: true,
          notify_after: 2,
          public: true,
          tls_cert: "",
          tls_cert_key: "",
          tls_cert_root: "",
      },
    }
  },


  // because route changes within the same component are re-used

  watch: {
    $route(to, from) {
      this.errorCode = 'none';
      this.ready = true;
    }
  },

  // beforeCreated() causes sync issues with mounted() as they are executed
  // one after the other regardless of async/await methods inside.

  mounted() {

    const id = this.$route.params.id

    if (id) {
      
      this.loadService(id);

    } else {

      this.errorCode = 'none';
      this.ready = true;

    };

  },

  methods: {

    async loadService(id){

	this.ready = false;

        // api still responds if session is invalid
        // specifically, if statping is restarted and an existing session exists
        // theres a further check to not display the data in the form component it seems ??

        await Api.service(id).then(
          response => {
              this.service = response;
              this.errorCode = 'none';
              this.ready = true;
          },
          error => {

              const respStatus = error.response.status;

              if ( respStatus == '404' ) {
                  this.errorCode = 404;
              } else if ( respStatus == '401' ) {
                   this.errorCode = 401 ;
              } else {
                  this.errorCode = 'none';
              };

              this.ready = true;

          }
      );

    }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
