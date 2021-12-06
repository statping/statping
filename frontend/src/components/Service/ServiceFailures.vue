<template>
    <div class="list-group mt-3 mb-4">

        <div v-for="(failure, index) in failures" :key="index" class="mb-2 list-group-item list-group-item-action flex-column align-items-start">
            <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{failure.issue}}</h5>
                <small>{{niceDate(failure.created_at)}}</small>
            </div>
            <p class="mb-1">{{failure.issue}}</p>
        </div>

        <nav v-if="total > 4" class="mt-3">
            <ul class="pagination justify-content-center">
                <li class="page-item" :class="{'disabled': page===1}">
                    <a @click.prevent="gotoPage(page-1)" :disabled="page===1" class="page-link" href="#" aria-label="Previous">
                        <span aria-hidden="true">&laquo;</span>
                        <span class="sr-only">Previous</span>
                    </a>
                </li>
                <li v-for="n in maxPages" class="page-item" :class="{'active': page === n}">
                    <a @click.prevent="gotoPage(n)" class="page-link" href="#">{{n}}</a>
                </li>
                <li class="page-item" :class="{'disabled': page===Math.floor(total / limit)}">
                    <a @click.prevent="gotoPage(page+1)" :disabled="page===Math.floor(total / limit)" class="page-link" href="#" aria-label="Next">
                        <span aria-hidden="true">&raquo;</span>
                        <span class="sr-only">Next</span>
                    </a>
                </li>
            </ul>
        </nav>
    </div>
</template>

<script>
const ServiceChart = () => import(/* webpackChunkName: "service" */ "./ServiceChart");
import Api from "../../API";

export default {
  name: 'ServiceFailures',
  components: {ServiceChart},
  props: {
    service: {
      type: Object,
      required: true
    },
  },
    data () {
        return {
            failures: [],
            limit: 4,
            offset: 0,
            total: this.service.stats.failures,
            page: 1
        }
    },
    computed: {
      pages() {
          return Math.floor(this.total / this.limit)
      },
        maxPages() {
          const p = Math.floor(this.total / this.limit)
          if (p > 16) {
              return 16
            } else {
              return p
          }
        }
    },
    async mounted () {
        await this.gotoPage(1)
    },
    methods: {
        async gotoPage(page) {
            this.page = page;

            this.offset = (page-1) * this.limit;

            window.console.log('page', this.page, this.limit, this.offset);

            this.failures = await Api.service_failures(this.service.id, 0, 9999999999, this.limit, this.offset)
        },
        smallText(s) {
            if (s.online) {
                return `Online, last checked ${this.ago(s.last_success)}`
            } else {
                return `Offline, last error: ${s.last_failure.issue} ${this.ago(s.last_failure.created_at)}`
            }
          }
      }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
