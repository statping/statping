<template>
<div class="mb-5">
  <h3 v-if="!loaded" >Import and Export</h3>
  <p v-if="!loaded">
    You can export your current Statping services, groups, notifiers, and other settings to a JSON file.
  </p>

  <div v-if="!loaded" class="mt-4 row">
    <div class="col-8 custom-file">
      <input @change="onFileChange" type="file" class="custom-file-input pointer" id="customFile" accept=".json,application/json">
      <label class="custom-file-label" for="customFile">Choose exported Statping JSON file</label>
    </div>
    <div class="col-4">
      <a class="btn btn-block btn-light btn-outline-secondary" href="/api/settings/export">Export</a>
    </div>
  </div>

  <div v-if="loaded" class="col-12 mb-4">
    <h3>Core Settings
      <span @click="file.core.enabled = !!file.core.enabled" class="switch switch-sm float-right">
            <input @change="update" v-model="file.core.enabled" type="checkbox" class="switch" :id="`switch-core`">
            <label :for="`switch-core`"></label>
          </span>
    </h3>

    <div class="row mb-2"><span class="col-4">Name</span><span class="col-8 text-right font-weight-bold">{{file.core.name}}</span></div>
    <div class="row mb-2"><span class="col-4">Description</span><span class="col-8 text-right font-weight-bold">{{file.core.description}}</span></div>
    <div class="row mb-2"><span class="col-4">Domain</span><span class="col-8 text-right font-weight-bold">{{file.core.domain}}</span></div>

  </div>

  <div v-if="loaded" class="col-12 mb-4">
    <h3>Users
      <button @click.prevent="toggle_all(file.users)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.users" class="alert alert-link">
      No Users in file
    </div>
    <div v-for="user in file.users" v-bind:key="user.id" class="row">
      <div class="col-4 font-weight-bold">
        {{user.username}}
      </div>
      <div class="col-6">
        {{user.email}}
      </div>
      <div class="col-2 text-right">
          <span @click="user.enabled = !!user.enabled" class="switch switch-sm">
            <input @change="update" v-model="user.enabled" type="checkbox" class="switch" :id="`switch-user-${user.id}`">
            <label :for="`switch-user-${user.id}`"></label>
          </span>
      </div>
    </div>
  </div>


  <div v-if="loaded" class="col-12 mb-4">
    <h3>Checkins
      <button @click.prevent="toggle_all(file.checkins)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.checkins" class="alert alert-link">
      No Checkins in file
    </div>
    <div v-for="checkin in file.checkins" v-bind:key="checkin.id" class="row">
      <div class="col-4 font-weight-bold">
        {{checkin.name}}
      </div>
      <div class="col-6">
        Service #{{checkin.service_id}}
      </div>
      <div class="col-2 text-right">
          <span @click="checkin.enabled = !!checkin.enabled" class="switch switch-sm">
            <input @change="update" v-model="checkin.enabled" type="checkbox" class="switch" :id="`switch-checkin-${checkin.id}`">
            <label :for="`switch-checkin-${checkin.id}`"></label>
          </span>
      </div>
    </div>
  </div>

  <div v-if="loaded" class="col-12 mb-4">
    <h3>Services
      <button @click.prevent="toggle_all(file.services)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.services" class="alert alert-link">
      No Services in file
    </div>
    <div v-for="service in file.services" v-bind:key="service.id" class="row">
        <div class="col-4 font-weight-bold">
          {{service.name}}
        </div>
      <div class="col-6">
        {{service.domain}}
      </div>
        <div class="col-2 text-right">
          <span @click="service.enabled = !!service.enabled" class="switch switch-sm">
            <input @change="update" v-model="service.enabled" type="checkbox" class="switch" :id="`switch-service-${service.id}`">
            <label :for="`switch-service-${service.id}`"></label>
          </span>
        </div>
    </div>
  </div>

  <div v-if="loaded" class="col-12 mb-4">
    <h3>Groups
      <button @click.prevent="toggle_all(file.groups)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.groups" class="alert alert-link">
      No Groups in file
    </div>
    <div v-for="group in file.groups" v-bind:key="group.id" class="row">
      <div class="col-4 font-weight-bold">
        {{group.name}}
      </div>
      <div class="col-8 text-right">
          <span @click="group.enabled = !!group.enabled" class="switch switch-sm">
            <input @change="update" v-model="group.enabled" type="checkbox" class="switch" :id="`switch-group-${group.id}`">
            <label :for="`switch-group-${group.id}`"></label>
          </span>
      </div>
    </div>
  </div>

  <div v-if="loaded" class="col-12 mb-4">
    <h3>Incidents
      <button @click.prevent="toggle_all(file.incidents)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.incidents" class="alert alert-link">
      No Incidents in file
    </div>
    <div v-for="incident in file.incidents" v-bind:key="incident.id" class="row">
      <div class="col-4 font-weight-bold">
        {{incident.name}}
      </div>
      <div class="col-8 text-right">
          <span @click="incident.enabled = !!incident.enabled" class="switch switch-sm">
            <input @change="update" v-model="incident.enabled" type="checkbox" class="switch" :id="`switch-incident-${incident.id}`">
            <label :for="`switch-incident-${incident.id}`"></label>
          </span>
      </div>
    </div>
  </div>

  <div v-if="loaded" class="col-12 mb-3">
    <h3>Notifiers
      <button @click.prevent="toggle_all(file.notifiers)" class="btn btn-sm btn-outline-dark float-right mt-1">Select All</button>
    </h3>
    <div v-if="!file.notifiers" class="alert alert-link">
      No Notifiers in file
    </div>
    <div v-for="notifier in file.notifiers" v-bind:key="notifier.id" class="row">
      <div class="col-4">
        {{notifier.title}}
      </div>
      <div class="col-8 text-right">
          <span @click="notifier.enabled = !!notifier.enabled" class="switch">
            <input @change="update" v-model="notifier.enabled" type="checkbox" class="switch" :id="`switch-notifier-${notifier.id}`">
            <label :for="`switch-notifier-${notifier.id}`"></label>
          </span>
      </div>
    </div>
  </div>

  <div v-if="error" class="alert alert-danger">
    {{error}}
  </div>

  <div class="col-12">
    <button v-if="loaded" @click.prevent="import_all" class="btn btn-block btn-success">Import</button>
  </div>

</div>
</template>

<script>
import Api from '../../API';

export default {
name: "Importer",
  data () {
    return {
      error: null,
      file: null,
      loaded: false,
      output: null,
      all: {
        notifiers: false,
        services: false,
        groups: false,
      }
    }
  },
  methods: {
  clean_elem(elem) {
    if (!elem) {
      return null
    }
    elem.map(e => delete(e.enabled) && delete(e.id))
    return elem
  },
  async import_all() {
    this.error = null
    const outgoing = {
      core: this.output.core,
      users: this.clean_elem(this.output.users),
      services: this.clean_elem(this.output.services),
      groups: this.clean_elem(this.output.groups),
      notifiers: this.clean_elem(this.output.notifiers),
      checkins: this.clean_elem(this.output.checkins),
    }
    try {
      await Api.import(outgoing)
    } catch(e) {
      this.error = e
    }
  },
  toggle_all(elem) {
    elem.map(s => s.enabled = true)
    this.update()
  },
    update() {
      this.output = {
        core: this.file.core.enabled ? this.file.core : null,
        users: this.file.users.filter(s => s.enabled),
        services: this.file.services.filter(s => s.enabled),
        groups: this.file.groups.filter(s => s.enabled),
        notifiers: this.file.notifiers.filter(s => s.enabled),
        checkins: this.file.checkins.filter(s => s.enabled),
      }
    },
    onFileChange(e) {
      let files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      this.processJSON(files[0]);
    },
    processJSON(file) {
      let reader = new FileReader();
      reader.onload = (e) => {
        this.file = JSON.parse(e.target.result);
        this.file.core.enabled = false
      };
      reader.readAsText(file);
      this.loaded = true
    },
  }
}
</script>

<style scoped>

</style>
