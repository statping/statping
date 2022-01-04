<template>
    <button v-html="loading ? loadLabel : label" @click.prevent="runAction" type="submit" :disabled="loading || disabled" class="btn btn-block" :class="{'btn-outline-light': loading}">
    </button>
</template>

<script>
  export default {
      name: 'LoadButton',
      props: {
        action: {
          type: Function,
          required: true
        },
        label: {
          type: String,
          required: true
        },
        disabled: {
          type: Boolean,
          default: false
        },
      },
      data() {
          return {
              loading: false,
              loadLabel: "<div class=\"spinner-border text-dark\"><span class=\"sr-only\">Loading</span></div>"
          }
      },
      methods: {
        async runAction() {
          this.loading = true;
          await this.action();
          this.loading = false;
        }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
