<template>
    <span class="switch">
        <input @change="toggle" type="checkbox" :name="name" class="switch" id="switch-normal" :value="value">
        <label for="switch-normal">{{label}}</label>
        <input type="hidden" :name="name" id="switch-normal-value" :value="value">
      </span>
</template>

<script>

export default {
  name: 'InputSwitch',
  props: {
    label: {
      type: String,
      required: true
    },
    value: {
      type: Boolean,
      required: true
    }
  },
  data () {
    return {

    }
  },
  mounted() {

  },
  methods: {
    toggle() {
      this.$emit('input', {
        value: !this.props.value
      })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
