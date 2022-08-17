<template>
  <div class="card contain-card mb-4">
    <div class="card-header d-flex align-items-center">
      <button
        class="btn p-0 mr-2"
        @click="$router.push('/dashboard/downtimes');"
      >
        <FontAwesomeIcon icon="arrow-circle-left" />
      </button>
      <div>{{ $t("downtime_info") }}</div>
    </div>
    <div class="card-body">
      <form>
        <div class="form-group row">
          <label class="col-sm-4 col-form-label">
            {{
              $t("service_name")
            }}
          </label>
          <div class="col-sm-8">
            <select
              v-model="downtime.serviceId"
              name="service"
              class="form-control"
              required
              :disabled="$route.params.id"
            >
              <option
                v-for="(service) in services"
                :key="service.id"
                :value="service.id"
              >
                {{ service.name }}
              </option>
            </select>
            <small
              class="form-text text-muted"
            >
              Select Servive you want to have downtime for
            </small>
          </div>
        </div>

        <div class="form-group row">
          <label class="col-sm-4 col-form-label">
            {{
              $t("downtime_status")
            }}
          </label>
          <div class="col-sm-8">
            <select
              v-model="downtime.subStatus"
              name="service"
              class="form-control"
              required
            >
              <option
                value="degraded"
              >
                Degraded
              </option>
              <option value="down">
                Down
              </option>
            </select>
            <small
              class="form-text text-muted"
            >
              Choose status you want to give to the Servive
            </small>
          </div>
        </div>

        <div class="form-group row">
          <label class="col-sm-4 col-form-label">
            {{
              $t("downtime_date_range")
            }}
          </label>
          <div class="col-sm-8">
            <div class="row">
              <div class="col-sm-6">
                <FlatPickr
                  id="start"
                  v-model="downtime.start"
                  type="text"
                  name="start"
                  class="form-control form-control-plaintext"
                  :config="config"
                  placeholder="Select Start Date"
                  @on-change="() => handleFormChange({target: {name: 'start'}})"
                />
                <small
                  v-if="errors.start"
                  class="form-text text-danger"
                >
                  {{ errors.start }}
                </small>
              </div>
              <div class="col-sm-6">
                <FlatPickr
                  id="end"
                  v-model="downtime.end"
                  type="text"
                  name="end"
                  class="form-control form-control-plaintext"
                  :config="config"
                  placeholder="Select End Date"
                  @on-change="() => handleFormChange({target: {name: 'end'}})"
                />
                <small
                  v-if="errors.end"
                  class="form-text text-danger"
                >
                  {{ errors.end }}
                </small>
              </div>
            </div>
            <small
              class="form-text text-muted"
            >
              Enter the Start and End date for which your service will be down/degraded
            </small>
          </div>
        </div>

        <div class="form-group row">
          <label class="col-sm-4 col-form-label">
            {{
              $t("failures")
            }}
          </label>
          <div class="col-sm-8">
            <input
              v-model.number="downtime.failures"
              type="number"
              name="check_interval"
              class="form-control"
              min="0"
              required
            >
            <small
              v-if="errors.failures"
              class="form-text text-danger"
            >
              {{ errors.failures }}
            </small>
            <small
              class="form-text text-muted"
            >
              Select the number of failures you want for your service
            </small>
          </div>
        </div>

        <div class="form-group row">
          <div class="col-12">
            <button
              :disabled="isLoading || !isCreateDowntimeBtnEnabled()"
              type="button"
              class="btn btn-success btn-block"
              @click.prevent="saveDowntime"
            >
              {{ $route.params.id ? $t("downtime_update") : $t("downtime_create") }}
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Api from '../API';
import FlatPickr from 'vue-flatpickr-component';
import 'flatpickr/dist/flatpickr.css';
import { convertToSec } from '../components/Dashboard/DashboardDowntimes.vue';

const checkFormErrors = (value, id) => {
    const { failures, start, end } = value;
    const errors = {};
    let endSec = ''; let startSec = '';
    
    // Converting into millisec
    if (start) {
        startSec = convertToSec(start);
    }

    if (end) {
        endSec = convertToSec(end);
    }

    // Check for valid positive numbers
    if (!(/^\d+$/.test(failures))) {
        errors.failures = 'Enter Valid Positve Number without decimal point';
    } else if (!start && end) {
        errors.start = 'Need to enter Start Date';
    } else if ( endSec && startSec > endSec ) {
        errors.end = 'End Date should be greater than Start Date';
    } 

    return errors;
};

export const removeEmptyParams = (obj) => {
    const updatedObj = {};

    for (const [ key , value ] of Object.entries(obj)) {
        if (value) {
            updatedObj[key] = value;
        }
    }

    return updatedObj;
};

export default {
    name: 'FormDowntime',
    components: {
        FlatPickr,
    },
    props: {
        editDowntime: {
            type: Object,
            default: null,
        }
    },
    data: function () {
        return {
            isLoading: false,
            errors: {},
            downtime: {
                serviceId: '',
                subStatus: 'degraded',
                failures: 20,
                start: new Date().toJSON(),
                end: null,
            },
            config: {
                altFormat: 'J M, Y, h:iK',
                altInput: true,
                enableTime: true,
                dateFormat: 'Z',
                maxDate: new Date().toJSON(),
            },
        };
    },
    computed: {
        ...mapState({
            services : function (state) {
                const { id } = this.$route.params;

                if (!id && state.services.length > 0) {
                    this.downtime.serviceId = state.services[0].id;
                }

                return state.services;
            }
        }),
    },
    mounted: function () {
        if (this.editDowntime) {
            const { service_id, sub_status, failures, start, end } = this.editDowntime;
  
            this.downtime = {
                start,
                end,
                failures,
                serviceId: service_id,
                subStatus: sub_status
            };
        }      
    },
    methods: {
        isCreateDowntimeBtnEnabled: function () {
            const { id } = this.$route.params;
            const { serviceId, subStatus, failures, start, end } = this.downtime;

            return serviceId && subStatus && failures && start;
        },
        saveDowntime: async function () {
            const { id } = this.$route.params;
            const errors = checkFormErrors(this.downtime, id);
            
            // Check of invalid input.
            if (Object.keys(errors).length > 0) {
                this.errors = Object.assign({}, errors);
                return;
            }

            const { serviceId, subStatus, ...rest } = removeEmptyParams(this.downtime);

            const downtime = {
                ...rest,
                ...(!id && { 'service_id': serviceId }),
                'sub_status': subStatus,
            };

            this.isLoading=true;

            try {
                if (id) {
                    await Api.downtime_update({ id, data: downtime });
                } else {
                    await Api.downtime_create(downtime);
                }
                this.isLoading=false;
            } catch (error) {
                this.isLoading=false;
                throw new Error(error.message);
            }

            this.$router.push('/dashboard/downtimes');
        },
        handleFormChange: function (e) {
            const { name } = e.target;

            delete this.errors[name];
        }
    }
};
</script>