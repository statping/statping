<template>
<div>
  <h3>Configuration</h3>
  For security reasons, all database credentials cannot be editted from this page.

  <codemirror v-show="loaded" v-model="configs" ref="configs" :options="cmOptions" class="mt-4 codemirrorInput"/>

</div>
</template>

<script>
import Api from "../../API";

import {codemirror} from 'vue-codemirror'
import('codemirror/lib/codemirror.css')
import('codemirror/mode/yaml/yaml.js')

export default {
name: "Configs",
  components: {
    codemirror
  },
  data() {
    return {
      loaded: false,
      configs: "okkoko: okokoko",
      cmOptions: {
        height: 700,
        tabSize: 4,
        lineNumbers: true,
        matchBrackets: true,
        mode: "text/x-yaml",
        line: true
      }
    }
  },
  mounted() {
    this.update()
  },
  methods: {
    async update() {
      this.loaded = false
      this.configs = await Api.configs()
      this.loaded = true
      this.$refs.configs.codemirror.refresh()
    }
  }
}
</script>

<style scoped>

</style>
