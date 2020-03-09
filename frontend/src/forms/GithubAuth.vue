<template>
    <form @submit="updateCore">
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Github Client ID</label>
            <div class="col-sm-8">
                <input v-model="clientId" type="text" class="form-control" placeholder="" required>
            </div>
        </div>
        <div class="form-group row">
            <label class="col-sm-4 col-form-label">Github Client Secret</label>
            <div class="col-sm-8">
                <input v-model="clientSecret" type="text" class="form-control" placeholder="" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="switch-group-public" class="col-sm-4 col-form-label">Enabled</label>
            <div class="col-md-8 col-xs-12 mt-1">
            <span @click="enabled = !!enabled" class="switch float-left">
                <input v-model="enabled" type="checkbox" class="switch" id="switch-group-public" :checked="enabled">
                <label for="switch-group-public">Enabled Github Auth</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button @click="updateCore" type="submit" :disabled="loading || group.name === ''" class="btn btn-block" :class="{'btn-primary': !group.id, 'btn-secondary': group.id}">
                    {{loading ? "Loading..." : group.id ? "Update Group" : "Create Group"}}
                </button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
</template>

<script>
  import Api from "../API";

  export default {
      name: 'FormGroup',
      props: {
          in_group: {
              type: Object
          },
          edit: {
              type: Function
          }
      },
      data() {
          return {
              loading: false,
              clientId: "",
              clientSecret: "",
              enabled: true,
          }
      },
      watch: {
          in_group() {
              this.group = this.in_group
          }
      },
      methods: {
          removeEdit() {
              this.group = {}
              this.edit(false)
          },
          async updateCore() {
              const g = this.group
              const data = {id: g.id, name: g.name, public: g.public}
              await Api.core_save(data)
              const groups = await Api.groups()
              this.$store.commit('setGroups', groups)
              this.edit(false)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
