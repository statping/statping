<template>
  <div>
    <div
      v-if="downtimes.length === 0"
      class="alert alert-dark d-block mt-3 mb-0"
    >
      You currently don't have any services!
    </div>
    
    <table
      v-else
      class="table"
    >
      <thead>
        <tr>
          <th scope="col">
            {{ $t('name') }}
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
            <span>
              <!-- {{ serviceById(downtime.service_id).name }} -->
              {{ downtime.service_id }}
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
            <span
              class=""
            >
              {{ niceDateWithYear(downtime.end) }}
            </span>
          </td>
          <td class="d-none d-md-table-cell">
            <span
              class="badge text-uppercase"
              :class="[downtime.sub_status === 'down' ? 'badge-danger' : 'badge-warning' ]"
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
            <div class="btn-group">
              <button
                v-if="$store.state.admin"
                :disabled="isLoading"
                class="btn btn-sm btn-outline-secondary"
                @click.prevent="goto({path: `/dashboard/edit_service/${service.id}`, params: {service: service} })"
              >
                <FontAwesomeIcon icon="edit" />
              </button>
              <button
                v-if="$store.state.admin"
                :disabled="isLoading"
                class="btn btn-sm btn-danger"
                @click.prevent="deleteService(service)"
              >
                <FontAwesomeIcon
                  v-if="!isLoading"
                  icon="times"
                />
                <FontAwesomeIcon
                  v-if="isLoading"
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
import { mapGetters, mapState } from 'vuex';

export default {
    data: function () {
        return {
            isLoading: false,
        };
    },
    computed: {
        ...mapState([ 'downtimes' ]),
        ...mapGetters([ 'serviceById' ])
    }
};
</script>