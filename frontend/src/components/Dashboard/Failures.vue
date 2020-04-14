<template>
    <div class="col-12">
        <h2>{{service.name}} Failures
        <button v-if="failures.length>0" @click="deleteFailures" class="btn btn-outline-danger float-right">Delete All</button></h2>
    <div class="list-group mt-3 mb-4">

        <div class="alert alert-info" v-if="failures.length===0">
            You don't have any failures for {{service.name}}. Way to go!
        </div>

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
            <div class="text-center">
             <span class="text-black-50">{{total}} Total</span>
            </div>
        </nav>
    </div>
    </div>
</template>

<script>
import Api from "../../API";

export default {
    name: 'Failures',
    data() {
        return {
            service: {},
          failures: [],
          limit: 10,
          offset: 0,
          total: 0,
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
      async created() {
        this.service = await Api.service(this.$route.params.id)
        this.total = this.service.stats.failures
        await this.gotoPage(1)
      },
    methods: {
      async deleteFailures() {
        const c = confirm('Are you sure you want to delete all failures?')
        if (c) {
          await Api.service_failures_delete(this.service)
          this.service = await Api.service(this.service.id)
          this.total = 0
          await this.load()
        }
      },
      async gotoPage(page) {
        this.page = page;
        this.offset = (page-1) * this.limit;
        await this.load()
      },
      async load() {
        this.failures = await Api.service_failures(this.service.id, 0, 9999999999, this.limit, this.offset)
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .sm {
        font-size: 8pt;
    }
</style>
