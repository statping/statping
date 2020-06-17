<template>
    <div class="col-12 full-col-12">
        <h4 v-if="group.name !== 'Empty Group'" class="group_header mb-3 mt-4">{{group.name}}</h4>
        <div class="list-group online_list mb-4">

            <a v-for="(service, index) in $store.getters.servicesInGroup(group.id)" v-bind:key="index" class="service_li list-group-item list-group-item-action">
                <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                <span class="badge text-uppercase float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">
                    {{service.online ? $t('online') : $t('offline')}}
                </span>

                <GroupServiceFailures :service="service"/>

                <IncidentsBlock :service="service"/>
            </a>

        </div>
    </div>
</template>

<script>
    import Api from '../../API';
    import GroupServiceFailures from './GroupServiceFailures';
    import IncidentsBlock from './IncidentsBlock';

export default {
  name: 'Group',
  components: {
      IncidentsBlock,
      GroupServiceFailures
  },
  props: {
    group: Object
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
