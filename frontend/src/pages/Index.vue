<template>
    <div class="container col-md-7 col-sm-12 mt-4 sm-container index_container">

        <Header/>

        <div v-for="(group, index) in $store.getters.groupsInOrder" v-bind:key="index">
            <Group :group=group />
        </div>

        <div v-for="(message, index) in $store.getters.messages" v-bind:key="index" v-if="inRange(message) && message.service === 0">
            <MessageBlock :message="message"/>
        </div>

        <div class="col-12 full-col-12">
            <div v-for="(service, index) in $store.getters.servicesInOrder" :ref="service.id" v-bind:key="index">
                <ServiceBlock :service=service />
            </div>
        </div>

    </div>
</template>

<script>
import Api from '../API';
import Group from '../components/Index/Group';
import Header from '../components/Index/Header';
import MessageBlock from '../components/Index/MessageBlock';
import ServiceBlock from '../components/Service/ServiceBlock';


export default {
    name: 'Index',
    components: {
        ServiceBlock,
        MessageBlock,
        Group,
        Header
    },
    data() {
        return {
            logged_in: false
        }
    },
    async created() {
        this.logged_in = this.loggedIn()
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
    .fade-enter-active, .fade-leave-active {
        transition: opacity .5s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
        opacity: 0;
    }
</style>
