<template>
    <div class="card contain-card mb-3">
        <div class="card-header">{{group.id ? `${$t('update')} ${group.name}` : $t('group_create')}}
            <transition name="slide-fade">
                <button @click="removeEdit" v-if="group.id" class="btn float-right btn-danger btn-sm">
                    {{ $t('close') }}
                </button>
            </transition></div>
        <div class="card-body">

    <form @submit.prevent="saveGroup">
        <div class="form-group row">
            <label for="title" class="col-sm-4 col-form-label">{{ $t('group') }} {{ $t('name') }}</label>
            <div class="col-sm-8">
                <input v-model="group.name" type="text" class="form-control" id="title" placeholder="Group Name" required>
            </div>
        </div>
        <div class="form-group row">
            <label for="switch-group-public" class="col-sm-4 col-form-label text-capitalize">{{ $t('public') }} {{ $t('group') }}</label>
            <div class="col-md-8 col-xs-12 mt-1">
            <span @click="group.public = !!group.public" class="switch float-left">
                <input v-model="group.public" type="checkbox" class="switch" id="switch-group-public" :checked="group.public">
                <label for="switch-group-public">{{$t('group_public_desc')}}</label>
            </span>
            </div>
        </div>
        <div class="form-group row">
            <div class="col-sm-12">
                <button @click.prevent="saveGroup" type="submit" :disabled="loading || group.name === ''" class="btn btn-block" :class="{'btn-primary': !group.id, 'btn-secondary': group.id}">
                    {{loading ? "Loading..." : group.id ? $t('group_update') : $t('group_create')}}
                </button>
            </div>
        </div>
        <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
    </form>
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
          async saveGroup() {
              this.loading = true
              if (this.group.id) {
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
              await this.update()
              this.group = {}
          },
          async updateGroup() {
              const g = this.group
              const data = {id: g.id, name: g.name, public: g.public}
              await Api.group_update(data)
              await this.update()
              this.edit(false)
          },
          async update() {
              const groups = await Api.groups()
              this.$store.commit('setGroups', groups)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
