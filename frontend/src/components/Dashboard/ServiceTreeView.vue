<template>
  <div>
    <div class="row">
      <div class="col-sm-6">
        <form>
          <div class="form-row">
            <div class="form-group col-sm-6 mb-md-0">
              <FlatPickr
                id="dateTime"
                ref="dateTimeRef"
                v-model="dateTime"
                type="text"
                name="dateTime"
                class="form-control form-control-plaintext"
                :config="{
                  altFormat: 'J M, Y, h:iK',
                  altInput: true,
                  enableTime: true,
                  dateFormat: 'Z',
                  maxDate: new Date().toJSON(),
                }"
                placeholder="Select Start Date"
              />
            </div>
            

            <div class="form-group col-sm-6 mb-md-0">
              <div role="group">
                <button
                  type="submit"
                  class="btn btn-primary mr-1"
                  :disabled="dateTime === '' || isLoading"
                  @click.prevent="handleFilterSearch"
                >
                  {{ $t('search') }}
                </button>
                <button
                  type="button"
                  class="btn btn-outline-secondary"
                  @click.prevent="handleClearFilter"
                >
                  Reset
                </button>
              </div>
            </div>
          </div>
        </form>
      </div>

      <div class="col-sm-6">
        <ul class="d-flex justify-content-end align-items-center">
          <li class="d-flex">
            <div class="mr-1 text-shade-success">
              <FontAwesomeIcon
                icon="circle"
                class="border border-secondary rounded-circle"
              />
            </div>
            <div>Up</div>
          </li>
          <li class="d-flex ml-3">
            <div class="mr-1 text-shade-warning">
              <FontAwesomeIcon
                icon="circle"
                class="border border-secondary rounded-circle"
              />
            </div>
            <div>Degraded</div>
          </li>
          <li class="d-flex ml-3">
            <div class="mr-1 text-shade-danger">
              <FontAwesomeIcon
                icon="circle"
                class="border border-secondary rounded-circle"
              />
            </div>
            <div>Down</div>
          </li>
        </ul>
      </div>
    </div>

    <div class="mt-3">
      <div v-if="isLoading">
        <div class="col-12 text-center">
          <FontAwesomeIcon
            icon="circle-notch"
            size="3x"
            spin
          />
        </div>
        <div class="col-12 text-center mt-3 mb-3">
          <span class="text-muted">
            Loading Services
          </span>
        </div>
      </div>
    
      <ul
        v-else
        class="parent-list-group pl-0 mb-0 overflow-auto"
      >
        <TreeItem
          v-for="service in treeData"
          :key="service.id"
          :item="service"
        />
      </ul>
    </div>
  </div>
</template>

<script>
import TreeItem from '../Elements/TreeItem.vue';
import FlatPickr from 'vue-flatpickr-component';
import 'flatpickr/dist/flatpickr.css';

const getRootNodes = (data) => {
    if (!data || data.lenght === 0) {
        return;
    }

    const rootNode = data.reduce((acc, service) => {
        const isChild = data.find((item) => {
            if (item.sub_services_detail) {
                return Object.keys(item.sub_services_detail).includes(String(service.id));
            }

            return false;
        });

        if (!isChild) {
            acc.push(service);
        }

        return acc;
    }, []);

    return rootNode;
};

const getTreeData = (parentServices, serviceStatus) => {
    const treeData = [];

    for (let i=0; i<parentServices.length; i++) {
        if (!parentServices[i].sub_services_detail) {
            treeData.push({ parent: parentServices[i], children: [] });
        } else {
            const subServices = Object.keys(parentServices[i].sub_services_detail).reduce((acc, key) => {
                const service = serviceStatus.find((item) => item.id == key);

                if (service) {
                    acc.push({ ...service, ...parentServices[i].sub_services_detail[key] });
                }

                return acc;
            }, []);

            const children = getTreeData(subServices, serviceStatus);

            treeData.push({ parent: parentServices[i], children });
        }
    }

    return treeData;
};

export default {
    name: 'ServiceTreeView',
    components: {
        TreeItem,
        FlatPickr
    },
    data: function () {
        return {
            isLoading: false,
            dateTime: '',
            treeData: [],
        };
    },
    computed: {
        serviceStatus () {
            return this.$store.state.serviceStatus;
        }
    },
    created: async function () {
        await this.getServiceStatus(this.dateTime);

        this.treeInitialize();
    },
    methods: {
        treeInitialize: function () {
            const rootNode = getRootNodes(this.serviceStatus);
            const treeData = getTreeData(rootNode, this.serviceStatus);
            this.treeData = treeData;
        },
        getServiceStatus: async function (dateTime) {
            let sec = null;

            this.isLoading = true;
            if (!dateTime) {
                sec = '';
            } else {
                sec = this.convertDateObjToSec(dateTime);
            }

            await this.$store.dispatch({ type: 'getServiceStatus', payload: sec });

            this.treeInitialize();
            this.isLoading = false;
        },
        handleFilterSearch: function () {
            // Remove focus and close date time picker on enter click.
            this.$refs.dateTimeRef.$el.nextSibling.blur();
            this.$refs.dateTimeRef.fp.close();

            this.getServiceStatus(this.dateTime);
        },
        handleClearFilter: function () {
            this.dateTime = new Date().toJSON();
        }
    }
};
</script>

<style scoped>
ul, li {
 margin: 0;
 padding: 0;
 list-style: none;
 display: block;
}

.text-shade-danger {
 color: #f7cecc;
 border: red
}

.text-shade-warning {
 color: #FFE6CC;
}

.text-shade-success {
 color: #D6E7D4;
}

.parent-list-group li {
 position: relative;
}
</style>