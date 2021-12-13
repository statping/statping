<template>
  <div class="col-12">
    <div class="card contain-card mb-4">
      <div class="card-header">
        {{ $t('downtimes') }}
        <!-- <router-link
          v-if="$store.state.admin"
          to="/dashboard/create_service"
          class="btn btn-sm btn-success float-right"
        >
          <FontAwesomeIcon icon="plus" />  {{ $t('create') }}
        </router-link> -->
      </div>
      <div class="card-body pt-0">
        <div
          v-if="isLoading"
          class="loader d-flex align-items-center justify-content-center"
        >
          <div
            class="spinner-border"
            role="status"
          >
            <span class="sr-only">
              Loading...
            </span>
          </div>
        </div>
        <div v-else>
          <DowntimesList />
          <Pagination
            :get-next-downtimes="getNextDowntimes"
            :get-prev-downtimes="getPrevDowntimes"
            :skip="params.skip"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DowntimesList from './DowntimesList.vue';
import Pagination from '../Elements/Pagination.vue';

export default {
    name: 'DashboardDowntimes',
    components: {
        DowntimesList,
        Pagination
    },
    data: function () {
        return {
            isLoading: false,
            params: { 
                serviceId: null,
                start: null,
                end: null,
                skip: 0,
                count: 10,
                subStatus: ''
            }
        };
    },
    async mounted () {
        this.getDowntimes(this.params);
    },
    methods: {
        getDowntimes: async function (params) {
            this.isLoading = true;
            await this.$store.dispatch({ type: 'getDowntimes', payload: params });
            this.isLoading = false;
        },
        getNextDowntimes: function () {
            this.params = { ...this.params, skip: this.params.skip + 1 };
            this.getDowntimes(this.params);
        },
        getPrevDowntimes: function () {
            this.params = { ...this.params, skip: this.params.skip + 1 };
            this.getDowntimes(this.params);
        }
    }
};
</script>

<style scoped>
.loader {
  min-height: 100px;
}
</style>