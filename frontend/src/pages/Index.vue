<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">

        <Header/>

        <div v-for="(group, index) in $store.getters.groups" v-bind:key="index">
            <Group :group=group />
        </div>

        <div class="col-12">
            <MessageBlock/>
        </div>

        <div class="col-12 full-col-12">

            <div v-for="(service, index) in $store.getters.services" v-bind:key="index">
                <ServiceBlock :service=service />
            </div>

        </div>
    </div>
</template>

<script>
const Header = () => import("@/components/Index/Header");
const ServiceBlock = () => import("@/components/Service/ServiceBlock.vue");
const MessageBlock = () => import("@/components/Index/MessageBlock");
const Group = () => import("@/components/Index/Group");

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

    }
  },
  created() {

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
