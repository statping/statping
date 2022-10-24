<template>
    <div class="container col-md-7 col-sm-12 sm-container">

      <Header/>

      <div v-if="!loaded" class="row mt-5 mb-5">
        <div class="col-12 mt-5 mb-2 text-center">
          <font-awesome-icon icon="circle-notch" class="text-dim" size="2x" spin/>
        </div>
        <div class="col-12 text-center mt-3 mb-3">
          <span class="text-dim">{{loading_text}}</span>
        </div>
      </div>

      <div class="col-12 full-col-12">
          <MessageBlock v-for="message in messages" v-bind:key="message.id" :message="message" />
      </div>

      <div class="col-12 full-col-12">
          <div v-for="service in services_no_group" v-bind:key="service.id" class="list-group online_list mb-4">
              <div class="list-group-item list-group-item-action">
                  <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                  <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                  <GroupServiceFailures :service="service"/>
                  <IncidentsBlock :service="service"/>
              </div>
          </div>
      </div>

      <Group v-for="group in groups" v-bind:key="group.id" :group=group />

      <div class="col-12 full-col-12">
          <div v-for="service in services" :ref="service.id" v-bind:key="service.id">
              <ServiceBlock :service="service" />
          </div>
      </div>

    </div>
</template>

<script>
import Api from "@/API";

const Group = () => import(/* webpackChunkName: "index" */ '@/components/Index/Group')
const Header = () => import(/* webpackChunkName: "index" */ '@/components/Index/Header')
const MessageBlock = () => import(/* webpackChunkName: "index" */ '@/components/Index/MessageBlock')
const ServiceBlock = () => import(/* webpackChunkName: "index" */ '@/components/Service/ServiceBlock')
const GroupServiceFailures = () => import(/* webpackChunkName: "index" */ '@/components/Index/GroupServiceFailures')
const IncidentsBlock = () => import(/* webpackChunkName: "index" */ '@/components/Index/IncidentsBlock')

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
            logged_in: false,
        }
    },
    computed: {
      loading_text() {
        if (this.$store.getters.groups.length === 0) {
          return "Loading Groups"
        } else if (this.$store.getters.services.length === 0) {
          return "Loading Services"
        } else if (this.$store.getters.messages == null) {
          return "Loading Announcements"
        }
      },
      loaded() {
        return this.$store.getters.services.length !== 0
      },
        core() {
          return this.$store.getters.core
        },
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
    methods: {
        async checkLogin() {
          const token = this.$cookies.get('statping_auth')
          if (!token) {
            this.$store.commit('setLoggedIn', false)
            return
          }
          try {
            const jwt = await Api.check_token(token)
            this.$store.commit('setAdmin', jwt.admin)
            if (jwt.username) {
              this.$store.commit('setLoggedIn', true)
            }
          } catch (e) {
            console.error(e)
          }
        },
        inRange(message) {
            return this.isBetween(this.now(), message.start_on, message.start_on === message.end_on ? this.maxDate().toISOString() : message.end_on)
        }
    }
}
</script>
