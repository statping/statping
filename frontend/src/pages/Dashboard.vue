<template>
    <div>
        <Login v-show="token === null"/>

        <div v-show="token !== null" class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <TopNav :changeView="changeView"/>

            <DashboardIndex v-show="view === 'DashboardIndex'" :services="services"/>

            <DashboardServices v-show="view === 'DashboardServices'" :services="services"/>

            <DashboardUsers v-show="view === 'DashboardUsers'" :services="services"/>

            <DashboardMessages v-show="view === 'DashboardMessages'" :services="services"/>

            <Settings v-show="view === 'Settings'" :services="services"/>

    </div>
    </div>
</template>

<script>
  import Api from "../components/API"
  import Login from "./Login";
  import TopNav from "../components/Dashboard/TopNav";
  import DashboardIndex from "../components/Dashboard/DashboardIndex";
  import DashboardServices from "../components/Dashboard/DashboardServices";
  import DashboardUsers from "../components/Dashboard/DashboardUsers";
  import DashboardMessages from "../components/Dashboard/DashboardMessages";
  import Settings from "./Settings";

  export default {
  name: 'Dashboard',
  components: {
    Settings,
    DashboardMessages,
    DashboardUsers,
    DashboardServices,
    DashboardIndex,
    TopNav,
    Login,
  },
  data () {
    return {
      services: null,
      groups: null,
      core: null,
      token: null,
      view: "DashboardIndex",
    }
  },
  created() {
    this.pathView(this.$route.path)
    this.token = Api.token()
    this.loadAll()
  },
  methods: {
    pathView (path) {
        switch (path) {
          case "/dashboard/settings":
            this.view = "Settings"
            break
          case "/dashboard/users":
            this.view = "DashboardUsers"
            break
          case "/dashboard/messages":
            this.view = "DashboardMessages"
            break
          case "/dashboard/services":
            this.view = "DashboardServices"
            break
          default:
            this.view = "DashboardIndex"
        }
    },
    changeView (v, name) {
        this.view = v
      this.$router.push('/'+name)
    },
    async loadAll () {
      this.token = await Api.token()
      this.core = await Api.root()
      this.groups = await Api.groups()
      this.services = await Api.services()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
