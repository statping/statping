<template>
    <div v-show="core" class="container col-md-7 col-sm-12 mt-2 sm-container">

        <Header :core="core"/>

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
import Api from "../components/API"

export default {
  name: 'Index',
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
    this.loadAll()
  },
  methods: {
    async loadAll () {
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
