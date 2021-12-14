<template>
  <div class="row">
    <DowntimesFilterForm
      :handle-clear-filters="handleClearFilters"
      :params="params"
      :handle-filter-search="handleFilterSearch"
    />
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
              v-if="downtimes.length !== 0"
              :get-next-downtimes="getNextDowntimes"
              :get-prev-downtimes="getPrevDowntimes"
              :skip="params.skip"
              :count="params.count"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DowntimesList from './DowntimesList.vue';
import { mapState } from 'vuex';
import DowntimesFilterForm from '../../forms/DowntimeFilters.vue';
import Pagination from '../Elements/Pagination.vue';

export const initialParams = { 
    serviceId: '',
    start: '',
    end: '',
    skip: 0,
    count: 10,
    subStatus: ''
};

export default {
    name: 'DashboardDowntimes',
    components: {
        DowntimesList,
        Pagination,
        DowntimesFilterForm
    },
    data: function () {
        return {
            isLoading: false,
            params: { ...initialParams }
        };
    },
    computed: {
        ...mapState([ 'downtimes' ])
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
            this.params = { ...this.params, skip: this.params.skip - 1 };
            this.getDowntimes(this.params);
        },
        handleClearFilters: function () {
            this.params = { ...initialParams };
        },
        handleFilterSearch: function () {
            this.params = { ...this.params, skip: 0 };
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