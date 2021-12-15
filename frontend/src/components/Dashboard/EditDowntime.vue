<template>
  <div class="col-12">
    <div
      v-if="isLoading"
      class="row mt-5"
    >
      <div class="col-12 text-center">
        <FontAwesomeIcon
          icon="circle-notch"
          size="3x"
          spin
        />
      </div>
      <div class="col-12 text-center mt-3 mb-3">
        <span class="text-muted">
          Loading Downtime
        </span>
      </div>
    </div>

    <FormDowntime />
  </div>
</template>

<script>
import Api from '../../API';

const FormDowntime = () =>
  import(/* webpackChunkName: "dashboard" */ '../../forms/Downtime.vue');

export default {
    name: 'EditDowntime',
    components: {
        FormDowntime
    },
    data: function () {
        return { isLoading: false, downtime: null };
    },
    created () {
        this.getDowntime();
    },
    methods: {
        getDowntime: async function () {
            const id = this.$route.params.id;
            if (!id) {
                return;
            }

            this.isLoading = true;
            const { output } = await Api.getDowntime(id);
            this.isLoading = false;

            this.downtime = output;
        }
    }
};
</script>