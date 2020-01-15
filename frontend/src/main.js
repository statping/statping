import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
const router = require("routes")

require("./assets/css/bootstrap.min.css")
require("./assets/css/base.css")

Vue.use(VueRouter)

Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
