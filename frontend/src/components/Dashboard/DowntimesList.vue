<template>
  <div>
    <div
      v-if="downtimes.length === 0"
      class="alert alert-dark d-block mt-3 mb-0"
    >
      You currently don't have any downtimes for this services!
    </div>
    
    <table
      v-else
      class="table"
    >
      <thead>
        <tr>
          <th scope="col">
            {{ $t('service') }}
          </th>
          <th
            scope="col"
            class="d-none d-md-table-cell"
          >
            {{ $t('start_time') }}
          </th>
          <th
            scope="col"
            class="d-none d-md-table-cell"
          >
            {{ $t('end_time') }}
          </th>
          <th
            scope="col"
            class="d-none d-md-table-cell"
          >
            {{ $t('status') }}
          </th>
          <th
            scope="col"
            class="d-none d-md-table-cell"
          >
            {{ $t('failures') }}
          </th>
          <th
            scope="col"
            class="d-none d-md-table-cell"
          >
            {{ $t('actions') }}
          </th>
        </tr>

        <tr
          v-for="downtime in downtimes"
          :key="downtime.id"
        >
          <td>
            <span :class="{'text-danger': !downtime.service}">
              {{ (downtime.service && downtime.service.name) || 'Deleted service' }}
            </span>
          </td>
          <td class="d-none d-md-table-cell">
            <span
              class=""
            >
              {{ niceDateWithYear(downtime.start) }}
            </span>
          </td>
          <td class="d-none d-md-table-cell">
            <span>
              {{ downtime.end ? niceDateWithYear(downtime.end) : 'Ongoing' }}
            </span>
          </td>
          <td class="d-none d-md-table-cell">
            <span
              class="badge text-uppercase"
              :class="[downtime.sub_status === 'down' ? 'badge-danger' : 'badge-warning']"
            >
              {{ downtime.sub_status }}
            </span>
          </td>
          <td class="d-none d-md-table-cell">
            <span
              class=""
            >
              {{ downtime.failures }}
            </span>
          </td>
          <td class="text-right">
            <div v-if="downtime.service" class="btn-group">
              <button
                v-if="$store.state.admin"
                :disabled="isLoading"
                class="btn btn-sm btn-outline-secondary"
                @click.prevent="goto(`/dashboard/edit_downtime/${downtime.id}`)"
              >
                <FontAwesomeIcon icon="edit" />
              </button>
              <button
                v-if="$store.state.admin"
                :disabled="downtimeDeleteId === downtime.id && isLoading"
                class="btn btn-sm btn-danger"
                @click.prevent="handleDowntimeDelete(downtime)"
              >
                <FontAwesomeIcon
                  v-if="!isLoading"
                  icon="trash"
                />
                <FontAwesomeIcon
                  v-if="downtimeDeleteId === downtime.id && isLoading"
                  icon="circle-notch"
                  spin
                />
              </button>
            </div>
          </td>
        </tr>
      </thead>
    </table>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Api from '../../API';

export default {
    name: 'DashboardDowntimeList',
    props: {
        getDowntimes: {
            type: Function,
            default: function () {}
        }
    },
    data: function () {
        return {
            isLoading: false,
            downtimeDeleteId: null,
        };
    },
    computed: {
        ...mapState([ 'downtimes' ]),
    },
    methods: {
        goto: function (to) {
            this.$router.push(to);
        },
        setIsLoading ({ id, isLoading }) {
            this.downtimeDeleteId = id;
            this.isLoading = isLoading;
        },
        delete: async function (id) {
            this.setIsLoading({ id, isLoading: true });
            await Api.downtime_delete(id);
            this.setIsLoading({ id: null, isLoading: false });

            this.getDowntimes();
        },
        handleDowntimeDelete: async function (downtime) {
            const modal = {
                visible: true,
                title: 'Delete Downtime',
                body: `Are you sure you want to delete the downtime for service ${downtime.service.name}?`,
                btnColor: 'btn-danger',
                btnText: 'Delete Downtime',
                func: () => this.delete(downtime.id),
            };
            this.$store.commit('setModal', modal);
        }
    }
};
</script>