<template>
    <div class="container col-md-7 col-sm-12 mt-2 sm-container">

        <Header/>

        <div v-for="(group, index) in $store.getters.groupsInOrder" v-bind:key="index">
            <Group :group=group />
        </div>

        <div v-for="(message, index) in $store.getters.messages" v-bind:key="index" v-if="inRange(message) && message.service === 0">
            <MessageBlock :message="message"/>
        </div>

        <div class="col-12 full-col-12">

            <div v-for="(service, index) in $store.getters.services" :ref="service.id" v-bind:key="index">
                <ServiceBlock :service=service />
            </div>

        </div>
    </div>
</template>

<script>
import Api from '../components/API';

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
  async created() {
      const core = await Api.core()
      this.$store.commit("setCore", core);
      if (!core.setup) {
          this.$router.push('/setup')
      }
  },
  async mounted() {

  },
  methods: {
      inRange(message) {
          const start = this.isBetween(new Date(), message.start_on)
          const end = this.isBetween(message.end_on, new Date())
          return start && end
      },
    clickService(s) {
      this.$nextTick(() => {
        this.$refs.s.scrollTop = 0;
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
