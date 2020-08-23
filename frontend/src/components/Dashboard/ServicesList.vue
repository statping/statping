<template>
    <div>
        <div v-if="servicesList.length === 0">
            <div class="alert alert-dark d-block mt-3 mb-0">
                You currently don't have any services!
            </div>
        </div>
    <table v-else class="table">
        <thead>
        <tr>
            <th scope="col">Name</th>
            <th scope="col" class="d-none d-md-table-cell">Status</th>
            <th scope="col" class="d-none d-md-table-cell">Visibility</th>
            <th scope="col" class="d-none d-md-table-cell">{{ $t('group') }}</th>
            <th scope="col" class="d-none d-md-table-cell" style="width: 130px">
              Failures
              <div class="btn-group float-right" role="group">
                <a @click="list_timeframe='3h'" type="button" class="small" :class="{'text-success': list_timeframe==='3h', 'text-muted': list_timeframe!=='3h'}">3h</a>
                <a @click="list_timeframe='12h'" type="button" class="small" :class="{'text-success': list_timeframe==='12h', 'text-muted': list_timeframe!=='12h'}">12h</a>
                <a @click="list_timeframe='24h'" type="button" class="small" :class="{'text-success': list_timeframe==='24h', 'text-muted': list_timeframe!=='24h'}">24h</a>
                <a @click="list_timeframe='7d'" type="button" class="small" :class="{'text-success': list_timeframe==='7d', 'text-muted': list_timeframe!=='7d'}">7d</a>
              </div>
            </th>
            <th scope="col"></th>
        </tr>
        </thead>
        <draggable id="services_list" tag="tbody" v-model="servicesList" handle=".drag_icon">
            <tr v-for="(service, index) in servicesList" :key="service.id">
                <td>
                    <span v-if="$store.state.admin" class="drag_icon d-none d-md-inline">
                        <font-awesome-icon icon="bars" class="mr-3"/>
                    </span> {{service.name}}
                </td>
              <td class="d-none d-md-table-cell">
                    <span class="badge text-uppercase" :class="{'badge-success': service.online, 'badge-danger': !service.online}">
                        {{service.online ? $t('online') : $t('offline')}}
                    </span>
              </td>
                <td class="d-none d-md-table-cell">
                    <span class="badge text-uppercase" :class="{'badge-primary': service.public, 'badge-secondary': !service.public}">
                        {{service.public ? $t('public') : $t('private')}}
                    </span>
                </td>
                <td class="d-none d-md-table-cell">
                    <div v-if="service.group_id !== 0">
                        <span class="badge badge-secondary">{{serviceGroup(service)}}</span>
                    </div>
                </td>
              <td class="d-none d-md-table-cell">
                <ServiceSparkList :service="service" :timeframe="list_timeframe"/>
              </td>
                <td class="text-right">
                    <div class="btn-group">
                        <button :disabled="loading" v-if="$store.state.admin" @click.prevent="goto({path: `/dashboard/edit_service/${service.id}`, params: {service: service} })" class="btn btn-sm btn-outline-secondary">
                            <font-awesome-icon icon="edit" />
                        </button>
                        <button :disabled="loading" @click.prevent="goto({path: serviceLink(service), params: {service: service} })" class="btn btn-sm btn-outline-secondary">
                            <font-awesome-icon icon="chart-area" />
                        </button>
                        <button :disabled="loading" v-if="$store.state.admin" @click.prevent="deleteService(service)" class="btn btn-sm btn-danger">
                            <font-awesome-icon v-if="!loading" icon="times" />
                            <font-awesome-icon v-if="loading" icon="circle-notch" spin/>
                        </button>
                    </div>
                </td>
            </tr>
        </draggable>
    </table>
    </div>
</template>

<script>
import Api from "../../API";
import ServiceSparkList from "@/components/Service/ServiceSparkList";
const draggable = () => import(/* webpackChunkName: "dashboard" */ 'vuedraggable')
const ToggleSwitch = () => import(/* webpackChunkName: "dashboard" */ '../../forms/ToggleSwitch');

export default {
      name: 'ServicesList',
    components: {
      ServiceSparkList,
        ToggleSwitch,
          draggable
    },
      data() {
        return {
          loading: false,
          list_timeframe: "12h",
          chartOpts: {
            chart: {
              type: 'bar',
              height: 50,
              sparkline: {
                enabled: true
              },
            },
            xaxis: {
              type: 'numeric',
            },
            showPoint: false,
            fullWidth:true,
            chartPadding: {top: 0,right: 0,bottom: 0,left: 0},
            stroke: {
              curve: 'straight'
            },
            fill: {
              opacity: 0.8,
            },
            yaxis: {
              min: 0
            },
            plotOptions: {
              bar: {
                colors: {
                  ranges: [{
                    from: 0,
                    to: 1,
                    color: '#39c10a'
                  }, {
                    from: 2,
                    to: 90,
                    color: '#e01a1a'
                  }]
                },
              },
            },
            tooltip: {
              theme: false,
              enabled: false,
            },
            title: {
              enabled: false,
            },
            subtitle: {
              enabled: false,
            }
          }
        }
      },
    computed: {
        servicesList: {
            get () {
                return this.$store.getters.servicesInOrder
            },
            set (value) {
                this.updateOrder(value)
            }
        }
    },
      methods: {
        goto(to) {
          this.$router.push(to)
        },
          async updateOrder(value) {
              let data = [];
              value.forEach((s, k) => {
                  data.push({ service: s.id, order: k + 1 })
              });
              await Api.services_reorder(data)
              await this.update()
          },
          async deleteService(s) {
              let c = confirm(`Are you sure you want to delete '${s.name}'?`)
              if (c) {
                this.loading = true
                  await Api.service_delete(s.id)
                  await this.update()
                this.loading = false
              }
          },
          serviceGroup(s) {
              let group = this.$store.getters.groupById(s.group_id)
              if (group) {
                  return group.name
              }
              return ""
          },
          async update() {
              const services = await Api.services()
              this.$store.commit('setServices', services)
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
