<template>
    <div>
        <h1 class="text-muted mt-5">
            {{group.id ? `Update ${group.name}` : "Create Group"}}
            <button @click="removeEdit" v-if="group.id" class="mt-3 btn float-right btn-danger btn-sm">Close</button>
        </h1>

        <div class="card">
            <div class="card-body">
    <form @submit="saveGroup">
        <div class="form-group row">
            <label for="title" class="col-sm-4 col-form-label">Group Name</label>
            <div class="col-sm-8">
                <input v-model="group.name" type="text" class="form-control" id="title" placeholder="Group Name" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="switch-group-public" class="col-sm-4 col-form-label">Public Group</label>
            <div class="col-8 mt-1">
            <span @click="group.public = !!group.public" class="switch float-left">
                <input v-model="group.public" type="checkbox" class="switch" id="switch-group-public" :checked="group.public">
                <label for="switch-group-public">Show group services to the public</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button @click="saveGroup" type="submit" :disabled="loading || group.name === ''" class="btn btn-block" :class="{'btn-primary': !group.id, 'btn-secondary': group.id}">
                    {{loading ? "Loading..." : group.id ? "Update Group" : "Create Group"}}
                </button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
            </div>
        </div>
    </div>
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
              group: {
                  name: "",
                  public: true
              }
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
          async saveGroup(e) {
              e.preventDefault();
              this.loading = true
              if (this.in_group) {
                  await this.updateGroup()
              } else {
                  await this.createGroup()
              }
              this.loading = false
          },
          async createGroup() {
              const g = this.group
              const data = {name: g.name, public: g.public}
              await Api.group_create(data)
              const groups = await Api.groups()
              this.$store.commit('setGroups', groups)
              this.group = {}
          },
          async updateGroup() {
              const g = this.group
              const data = {id: g.id, name: g.name, public: g.public}
              await Api.group_update(data)
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
