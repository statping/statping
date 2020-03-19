<template>
    <font-awesome-icon @click="toggleChecking" class="toggle-service" :class="{'text-success': running, 'text-muted': !running}" :icon="icon" />
</template>

<script>

  export default {
      name: 'ToggleSwitch',
      props: {
          service: {
              type: Object,
              required: true
          }
      },
      data() {
          return {
              icon: "toggle-on",
              running: true
          }
      },
      mounted() {
          if (this.service.online) {
              this.running = true
              this.icon = "toggle-on"
          } else {
              this.running = false
              this.icon = "toggle-off"
          }
      },
      methods: {
          toggleChecking() {
              if (this.running) {
                  this.icon = "toggle-off"
              } else {
                  this.icon = "toggle-on"
              }
              this.running = !this.running
          },
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
