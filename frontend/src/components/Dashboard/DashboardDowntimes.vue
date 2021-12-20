<template>
  <div class="col-12">
    <DowntimesFilterForm
      :handle-clear-filters="handleClearFilters"
      :params="params"
      :handle-filter-search="handleFilterSearch"
      :filter-errors="filterErrors"
      :handle-filter-change="handleFilterChange"
    />
    
    <div class="card contain-card mb-4">
      <div class="card-header">
        {{ $t('downtimes') }}
        <router-link
          v-if="$store.state.admin"
          to="/dashboard/create_downtime"
          class="btn btn-sm btn-success float-right"
        >
          <FontAwesomeIcon icon="plus" />  {{ $t('create') }}
        </router-link>
      </div>
      <div class="card-body pt-0">
        <div
          v-if="isLoading"
          class="mt-5"
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
              Loading Downtimes
            </span>
          </div>
        </div>
        <div v-else>
          <DowntimesList :get-downtimes="getDowntimes" />
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

const convertToSec = (val) => {
    return +new Date(val)/1000;
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
            params: { ...initialParams },
            filterErrors: {}
        };
    },
    computed: {
        ...mapState([ 'downtimes' ]),
    },
    async mounted () {
        this.getDowntimes(this.params);
    },
    methods: {
        getDowntimes: async function (params = this.params) {
            const { start, end } = params;

            this.checkFilterErrors();

            if (Object.keys(this.filterErrors).length > 0) {
                return;
            }

            const startSec = convertToSec(start);
            const endSec = convertToSec(end) + (60 * 60 * 23 + 59 * 60 + 59); // adding end of time for that particular date.
        

            this.isLoading = true;
            await this.$store.dispatch({ type: 'getDowntimes', payload: { ...params, start: startSec, end: endSec } });
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

            this.getDowntimes();
        },
        checkFilterErrors: function () {
            const { start, end } = this.params;
            const errors = {};

            // Converting into millisec
            const startSec = convertToSec(start);
            const endSec = convertToSec(end) + (60 * 60 * 23 + 59 * 60 + 59);

            if (!start && end) {
                errors.start = 'Need to enter Start Date';
            } else if (start && !end) {
                errors.end = 'Need to enter End Date';
            } else if ( startSec > endSec ) {
                errors.end = 'End Date should be greater than Start Date';
            }

            this.filterErrors = Object.assign({}, errors);
        },
        handleFilterChange: function (e) {
            // reset all the errors
            const { name } = e.target;

            delete this.filterErrors[name];
        }
    }
};
</script>