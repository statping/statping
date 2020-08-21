<template>
    <div v-if="services.length > 0" class="col-12 full-col-12">
        <h4 v-if="group.name !== 'Empty Group'" class="group_header mb-3 mt-4">{{group.name}}</h4>
        <div class="list-group online_list mb-4">

            <div v-for="(service, index) in services" v-bind:key="index" class="list-group-item list-group-item-action">
                <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge text-uppercase float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>

                <GroupServiceFailures :service="service"/>

                <IncidentsBlock :service="service"/>

            </div>

        </div>
    </div>
</template>

<script>
    const GroupServiceFailures = () => import(/* webpackChunkName: "index" */ './GroupServiceFailures');
    const IncidentsBlock = () => import(/* webpackChunkName: "index" */ './IncidentsBlock');

export default {
  name: 'Group',
  components: {
      IncidentsBlock,
      GroupServiceFailures
  },
  props: {
    group: {
      type: Object,
      required: true,
    }
  },
  computed: {
    services() {
      return this.$store.getters.servicesInGroup(this.group.id)
    }
  }
}
</script>
