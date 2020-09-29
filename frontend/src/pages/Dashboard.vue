<template>
    <div class="container col-md-7 col-sm-12 mt-md-5">
      <div v-if="modal" class="modal-backdrop"></div>
      <Modal/>
        <TopNav :admin="admin"/>
        <router-view :admin="admin"/>
    </div>
</template>

<script>
  import Modal from "@/components/Elements/Modal";
  const TopNav = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/TopNav')

  export default {
  name: 'Dashboard',
  components: {
    Modal,
    TopNav,
  },
  data () {
      return {
          authenticated: false,
          loaded: false,
      }
  },
    computed: {
      modal() {
        return this.$store.getters.modal.visible
      },
      admin() {
        return this.$store.getters.admin
      },
      user() {
        return this.$store.getters.user
      }
    },
    mounted() {
      // if (!this.user || !this.admin) {
      //   this.$router.push('/login')
      // }
    }
  }
</script>

