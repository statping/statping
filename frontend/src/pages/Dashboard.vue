<template>
    <div>
        <Login v-show="$store.getters.token === null"/>

        <div v-show="$store.getters.token !== null" class="container col-md-7 col-sm-12 mt-md-5 bg-light">

        <TopNav :changeView="changeView"/>

            <DashboardIndex v-show="view === 'DashboardIndex'"/>

            <DashboardServices v-show="view === 'DashboardServices'"/>

            <DashboardUsers v-show="view === 'DashboardUsers'"/>

            <DashboardMessages v-show="view === 'DashboardMessages'"/>

            <Settings v-show="view === 'Settings'"/>

            <ServiceForm v-show="view === 'ServiceForm'"/>

    </div>
    </div>
</template>

<script>
  import Api from "../components/API"
  import ServiceForm from '../components/Dashboard/ServiceForm';
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
      ServiceForm,
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
      view: "DashboardIndex",
        authenticated: false
    }
  },
  created() {
    this.pathView(this.$route.path)
    this.isAuthenticated()
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
            case "/service/create":
                this.view = "ServiceForm"
                break
          default:
            this.view = "DashboardIndex"
        }
    },
    changeView (v, name) {
        this.view = v
      this.$router.push('/'+name)
    },
      isAuthenticated () {
        const token = this.$store.getters.token
        if (token.token) {
            this.authenticated = true
            if (!this.$store.getters.hasAllData) {
                this.loadAllData()
            }
        }
    },
      async loadAllData () {
        const users = await Api.users()
        const groups = await Api.groups()
        const messages = await Api.messages()
        const notifiers = await Api.notifiers()
          this.$store.commit('setMessages', messages)
          this.$store.commit('setUsers', users)
          this.$store.commit('setGroups', groups)
          this.$store.commit('setNotifiers', notifiers)
          this.$store.commit('setHasAllData', true)
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
