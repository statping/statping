<template>
    <div class="col-12 mb-3 pb-2 border-bottom" role="alert">
        <span class="font-weight-bold text-capitalize" :class="{'text-success': update.type.toLowerCase()==='resolved', 'text-danger': update.type.toLowerCase()==='issue summary', 'text-warning': update.type.toLowerCase()==='update'}">{{update.type}}</span>
        <span class="text-muted">- {{update.message}}
            <button v-if="admin" :disabled="isLoading && incidentId" @click="delete_update(update)" type="button" class="close">
                <FontAwesomeIcon v-if="isLoading && incidentId === update.id" icon="circle-notch" spin size="xs" />
                <FontAwesomeIcon v-else icon="trash" size="xs" />
            </button>
        </span>
        <span class="d-block small">{{ago(update.created_at)}} ago</span>
    </div>
</template>

<script>
  import Api from "@/API";

  export default {
    name: "IncidentUpdate",
    props: {
      update: {
        required: true
      },
      admin: {
        required: true
      },
      onUpdate: {
        required: false
      }
    },
    data() {
      return {
        isLoading: false,
        incidentId: null
      }
    },
    methods: {
      async delete_update(update) {
        this.isLoading = true;
        this.incidentId = update.id;

        const res = await Api.incident_update_delete(update);

        if (res.status === "success") {
          this.onUpdate();

          this.isLoading = false;
          this.incidentId = null;
        }
      },
    }
  }
</script>

<style scoped>

</style>
