<template>
    <div class="col-12 full-col-12">
        <h4 class="group_header">{{group.name}}</h4>
        <div class="list-group online_list mb-3">

            <GroupedService/>

        </div>
    </div>
</template>

<script>
import GroupedService from './GroupedService'

export default {
  name: 'Group',
  components: {
    GroupedService
  },
  props: {
    group: Object
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
