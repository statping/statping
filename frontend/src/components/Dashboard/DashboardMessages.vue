<template>
    <div>
    <div class="col-12">
        <h1 class="text-black-50">Messages</h1>
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">Title</th>
                <th scope="col" class="d-none d-md-table-cell">Service</th>
                <th scope="col" class="d-none d-md-table-cell">Begins</th>
                <th scope="col"></th>
            </tr>
            </thead>
            <tbody>

            <tr v-for="(message, index) in $store.getters.messages" v-bind:key="index">
                <td>{{message.title}}</td>
                <td class="d-none d-md-table-cell">
                    <router-link to="/service/${service(message.service).id}">{{service(message.service)}}</router-link>
                </td>
                <td class="d-none d-md-table-cell">{{message.start_on}}</td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="message/1" class="btn btn-outline-secondary"><i class="fas fa-exclamation-triangle"></i> Edit</a>
                        <a @click="deleteMessage(message)" href="#" class="btn btn-danger"><font-awesome-icon icon="times" /></a>
                    </div>
                </td>
            </tr>

            </tbody>
        </table>
    </div>


    <div class="col-12">
        <h1 class="text-black-50 mt-5">Create Message</h1>
        <div class="card">
            <div class="card-body">
                <FormMessage/>
            </div>
        </div>
    </div>
    </div>
</template>

<script>
  import Api from "../API"
  import FormMessage from "../../forms/Message";

  export default {
  name: 'DashboardMessages',
    components: {FormMessage},
    data () {
    return {

    }
  },
  created() {

  },
  methods: {
    service (id) {
        return this.$store.getters.serviceById(id).name || ""
    },
    async deleteMessage(m) {
      let c = confirm(`Are you sure you want to delete message '${m.title}'?`)
      if (c) {
        await Api.message_delete(m.id)
        const messages = await Api.messages()
        this.$store.commit('setMessages', messages)
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
