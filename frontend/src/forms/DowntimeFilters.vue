<template>
  <div class="card contain-card mb-4">
    <div class="card-header">
      {{ $t('filters') }}
    </div>
    <div class="card-body">
      <form>
        <div class="form-row">
          <div class="form-group col-md-3">
            <label class="col-form-label">
              {{ $t('service') }}
            </label>
            <select
              v-model="params.serviceId"
              name="service"
              class="form-control"
            >
              <option value="">
                Select Service
              </option>
              <option
                v-for="service in services"
                :key="service.id"
                :value="service.id"
              >
                {{ service.name }}
              </option>
            </select>
          </div>
          <div class="form-group col-md-4">
            <label class="col-form-label">
              {{ $t('downtime_date_range') }}
            </label>
            <div class="form-row">
              <div class="col-sm-6">
                <FlatPickr
                  id="start"
                  v-model="params.start"
                  type="text"
                  name="start"
                  class="form-control form-control-plaintext"
                  value=""
                  :config="config"
                  placeholder="Select Start Date"
                />
              </div>
              <div class="col-sm-6">
                <FlatPickr
                  id="end"
                  v-model="params.end"
                  type="text"
                  name="end"
                  class="form-control form-control-plaintext"
                  value=""
                  :config="config"
                  placeholder="Select End Date"
                />
              </div>
            </div>
          </div>
          <div class="form-group col-md-2">
            <label class="col-form-label">
              {{ $t('status') }}
            </label>
            <select
              v-model="params.subStatus"
              name="status"
              class="form-control"
            >
              <option value="">
                Select Status
              </option>
              <option value="degraded">
                Degraded
              </option>
              <option value="down">
                Down
              </option>
            </select>
          </div>

          <div class="form-group col-md-3 d-flex align-items-end">
            <div
              class="ml-auto"
              role="group"
            >
              <button
                type="submit"
                class="btn btn-primary mr-1"
                @click.prevent="handleFilterSearch"
              >
                {{ $t('search') }}
              </button>
              <button
                type="button"
                class="btn btn-outline-secondary"
                @click.prevent="handleClearFilters"
              >
                {{ $t('clear') }}
              </button>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import FlatPickr from 'vue-flatpickr-component';
import 'flatpickr/dist/flatpickr.css';
import { initialParams } from '../components/Dashboard/DashboardDowntimes.vue';

export default {
    name: 'DashboardDowntimeFilters',
    components: {
        FlatPickr
    },
    props: {
        params: {
            type: Object,
            default: initialParams
        },
        handleClearFilters: {
            type: Function,
            default: function () {}
        },
        handleFilterSearch: {
            type: Function,
            default: function () {}
        }
    },
    data: function () {
        return {
            config: {
                altFormat: 'D, J M Y',
                altInput: true,
                dateFormat: 'Z',
                maxDate: new Date()
            },
        };
    },
    computed: {
        ...mapState([ 'services' ]) 
    },
};
</script>