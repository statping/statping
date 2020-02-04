<template >
    <form method="POST" action="settings/css">
        <ul class="nav nav-pills mb-3" id="pills-tab" role="tablist">
            <li class="nav-item col text-center">
                <a class="nav-link active" id="pills-vars-tab" data-toggle="pill" href="#pills-vars" role="tab" aria-controls="pills-vars" aria-selected="true">Variables</a>
            </li>
            <li class="nav-item col text-center">
                <a class="nav-link" id="pills-theme-tab" data-toggle="pill" href="#pills-theme" role="tab" aria-controls="pills-theme" aria-selected="false">Base Theme</a>
            </li>
            <li class="nav-item col text-center">
                <a class="nav-link" id="pills-mobile-tab" data-toggle="pill" href="#pills-mobile" role="tab" aria-controls="pills-mobile" aria-selected="false">Mobile</a>
            </li>
        </ul>
        <div class="tab-content" id="pills-tabContent">
            <div class="tab-pane show active" id="pills-vars" role="tabpanel" aria-labelledby="pills-vars-tab">
                <codemirror v-if="loaded"  v-model="base" :options="cmOptions"></codemirror>
            </div>
        </div>
        <button type="submit" class="btn btn-primary btn-block mt-2">Save Style</button>
        <a href="settings/delete_assets" class="btn btn-danger btn-block confirm-btn">Delete All Assets</a>
    </form>
</template>

<script>
  import Api from "../API";

  // require component
  import { codemirror } from 'vue-codemirror'
  // require styles
  import 'codemirror/lib/codemirror.css'

  export default {
  name: 'ThemeEditor',
  components: {
    codemirror
  },
  props: {
    core: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      base: "",
      loaded: false,
      cmOptions: {
        // codemirror options
        tabSize: 4,
        mode: 'text/javascript',
        theme: 'base16-dark',
        lineNumbers: true,
        line: true,
        // more codemirror options, 更多 codemirror 的高级配置...
      }
    }
  },
  async mounted() {
    this.base = await Api.scss_base()
     window.console.log(this.base)
    this.loaded = true
  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
