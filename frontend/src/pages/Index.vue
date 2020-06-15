<template>
    <div class="container col-md-7 col-sm-12 sm-container index_container">

        <Header/>

        <div class="col-12 full-col-12">
            <div v-for="service in services_no_group" v-bind:key="service.id" class="list-group online_list mb-4">
                <a class="service_li list-group-item list-group-item-action">
                    <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                    <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    <GroupServiceFailures :service="service"/>
                    <IncidentsBlock :service="service"/>
                </a>
            </div>
        </div>

        <div>
            <Group v-for="group in groups" v-bind:key="group.id" :group=group />
        </div>

        <div class="col-12 full-col-12">
            <MessageBlock v-for="message in messages" v-bind:key="message.id" :message="message" />
        </div>

        <div class="col-12 full-col-12">
            <div v-for="service in services" :ref="service.id" v-bind:key="service.id">
                <ServiceBlock :in_service=service />
            </div>
        </div>

    </div>
</template>

<script>
const Group = () => import('@/components/Index/Group')
const Header = () => import('@/components/Index/Header')
const MessageBlock = () => import('@/components/Index/MessageBlock')
const ServiceBlock = () => import('@/components/Service/ServiceBlock')
const GroupServiceFailures = () => import('@/components/Index/GroupServiceFailures')
const IncidentsBlock = () => import('@/components/Index/IncidentsBlock')

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
            return this.$store.getters.messages.filter(m => this.inRange(m) && m.service === 0)
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
            return this.isBetween(this.now(), message.start_on, message.start_on === message.end_on ? this.maxDate().toISOString() : message.end_on)
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
