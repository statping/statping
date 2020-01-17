<template>
    <div class="col-12 full-col-12">
        <h4 v-if="group.name !== 'Empty Group'" class="group_header">{{group.name}}</h4>
        <div class="list-group online_list mb-3">

            <a v-for="(service, index) in $store.getters.servicesInGroup(group.id)" v-bind:key="index" href="#" class="service_li list-group-item list-group-item-action " data-id="7">
                {{service.name}}
                <span class="badge bg-success float-right pulse-glow">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
            </a>

        </div>
    </div>
</template>

<script>

export default {
  name: 'Group',
  components: {

  },
  props: {
    group: Object
  },
    methods: {
      serviceBadge (s) {

      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
