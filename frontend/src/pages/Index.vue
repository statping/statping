<template>
    <div class="container col-md-7 col-sm-12 sm-container index_container">

        <Header/>

        <div v-for="(service, i) in services_no_group" class="col-12 full-col-12">
            <div class="list-group online_list mb-4">
                <a class="service_li list-group-item list-group-item-action">
                    <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                    <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">{{service.online ? "ONLINE" : "OFFLINE"}}</span>

                    <GroupServiceFailures :service="service"/>

                    <IncidentsBlock :service="service"/>
                </a>
            </div>
        </div>

        <div v-for="(group, index) in groups" v-bind:key="index">
            <Group :group=group />
        </div>

        <div v-for="(message, index) in messages" v-bind:key="index" v-if="inRange(message) && message.service === 0">
            <MessageBlock :message="message"/>
        </div>

        <div class="col-12 full-col-12">
            <div v-for="(service, index) in services" :ref="service.id" v-bind:key="index">
                <ServiceBlock :in_service=service />
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
import GroupServiceFailures from "../components/Index/GroupServiceFailures";
import IncidentsBlock from "../components/Index/IncidentsBlock";


export default {
    name: 'Index',
    components: {
      IncidentsBlock,
      GroupServiceFailures,
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
    computed: {
        messages() {
            return this.$store.getters.messages
        },
        groups() {
            return this.$store.getters.groupsInOrder
        },
        services() {
            return this.$store.getters.servicesInOrder
        },
        services_no_group() {
            return this.$store.getters.servicesNoGroup
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
