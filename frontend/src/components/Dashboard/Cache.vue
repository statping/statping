<template>
    <div>
        <table class="table">
            <thead>
            <tr>
                <th scope="col">URL</th>
                <th scope="col">Size</th>
                <th scope="col">Expiration</th>
            </tr>
            </thead>
            <tbody>

            <tr v-for="(cache, index) in cache">
                <td>{{cache.url}}</td>
                <td>{{cache.size}}</td>
                <td>{{ago(cache.expiration)}}</td>
            </tr>

            </tbody>
        </table>
        <button @click.prevent="clearCache" class="btn btn-danger btn-block">Clear Cache</button>
    </div>
</template>

<script>
import Api from "../../API";

export default {
      name: 'Cache',
      data() {
          return {
              cache: [],
          }
      },
      async mounted() {
          this.cache = await Api.cache()
      },
      methods: {
          async clearCache() {
              await Api.clearCache()
              this.cache = []
          }
      }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
