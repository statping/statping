<template>
    <form @submit="saveGroup" action="api/groups" method="POST">
        <div class="form-group row">
            <label for="title" class="col-sm-4 col-form-label">Group Name</label>
            <div class="col-sm-8">
                <input v-model="group.name" type="text" name="name" class="form-control" value="" id="title" placeholder="Group Name" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="switch-group-public" class="col-sm-4 col-form-label">Public Group</label>
            <div class="col-8 mt-1">
            <span class="switch float-left">
                <input v-model="group.public" type="checkbox" name="public" class="switch" id="switch-group-public" >
                <label for="switch-group-public">Show group services to the public</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button v-on:click="saveGroup" type="submit" class="btn btn-primary btn-block">Create Group</button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
export default {
  name: 'FormGroup',
  props: {
    group: Object
  },
  data () {
    return {
      group: {}
    }
  },
  mounted() {
    if (this.props.group) {
      this.group = this.props.group
    }
  },
  methods: {
    saveGroup: {

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
