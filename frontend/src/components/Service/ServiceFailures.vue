<template>
    <div class="list-group mt-3 mb-4">

        <div v-for="(failure, index) in failures" :key="index" class="mb-2 list-group-item list-group-item-action flex-column align-items-start">
            <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{failure.issue}}</h5>
                <small>{{toLocal(failure.created_at)}}</small>
            </div>
            <p class="mb-1">{{failure.issue}}</p>
        </div>

        <nav aria-label="Page navigation example">
            <ul class="pagination">
                <li class="page-item">
                    <a class="page-link" href="#" aria-label="Previous">
                        <span aria-hidden="true">&laquo;</span>
                        <span class="sr-only">Previous</span>
                    </a>
                </li>
                <li class="page-item"><a class="page-link" href="#">1</a></li>
                <li class="page-item"><a class="page-link" href="#">2</a></li>
                <li class="page-item"><a class="page-link" href="#">3</a></li>
                <li class="page-item">
                    <a class="page-link" href="#" aria-label="Next">
                        <span aria-hidden="true">&raquo;</span>
                        <span class="sr-only">Next</span>
                    </a>
                </li>
            </ul>
        </nav>
    </div>
</template>

<script>
import ServiceChart from "./ServiceChart";
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
            limit: 15,
            offset: 0,
        }
    },
    async mounted () {
      this.failures = await Api.service_failures(this.service.id, this.now(), this.now(), this.limit, this.offset)
    },
    methods: {
        smallText(s) {
            if (s.online) {
                return `Online, last checked ${this.ago(s.last_success)}`
            } else {
                return `Offline, last error: ${s.last_failure.issue} ${this.ago(s.last_failure.created_at)}`
            }
          },
          ago(t1) {
            const tm = this.parseTime(t1)
            return this.duration(this.$moment().utc(), tm)
          }
      }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
