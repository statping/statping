<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">

        <Header :core="core"/>

        <Group/>
        <Group/>
        <Group/>
        <div v-for="(group, index) in groups" v-bind:key="index">
            <Group :group=group />
        </div>

        <div class="col-12">
            <MessageBlock/>
        </div>

        <div class="col-12 full-col-12">

            <div v-for="(service, index) in services" v-bind:key="index">
                <ServiceBlock :service=service />
            </div>

        </div>
    </div>
</template>

<script>
import ServiceBlock from '../components/Service/ServiceBlock.vue'
import MessageBlock from "../components/Index/MessageBlock";
import Group from "../components/Index/Group";
import Header from "../components/Index/Header";

export default {
  name: 'Dashboard',
  components: {
    Header,
    Group,
    MessageBlock,
    ServiceBlock,
  },
  data () {
    return {
      services: null,
      groups: null,
      core: null,
    }
  },
  beforeMount() {
    this.getAPI()
    this.getGroups()
    this.getServices()
  },
  methods: {
    getAPI: function() {
      axios
        .get('/api')
        .then(response => (this.core = response.data))
    },
    getServices: function() {
      axios
        .get('/api/services')
        .then(response => (this.services = response.data))
    },
    getGroups: function() {
      axios
        .get('/api/groups')
        .then(response => (this.groups = response.data))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
