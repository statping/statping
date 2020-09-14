<template>
 <div v-if="modal.visible" class="modal d-block mt-5 pt-5" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{modal.title}}</h5>
          </div>
          <div class="modal-body">
            <p>{{modal.body}}</p>
          </div>
          <div class="modal-footer">
            <button @click.prevent="close" type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            <button @click.prevent="runFunc" type="button" :class="`btn ${modal.btnColor}`">{{modal.btnText}}</button>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
export default {
name: "Modal",
  data () {
    return {

    }
  },
  computed: {
    modal() {
     return this.$store.getters.modal
    }
  },
  mounted() {

  },
  methods: {
    runFunc() {
      this.$store.getters.modal.func()
      this.close()
    },
    close() {
      this.$store.commit("setModal", {visible: false})
    }
  }
}
</script>

<style scoped>

</style>
