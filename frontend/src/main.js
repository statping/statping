import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Index from "./components/Pages/Index";

require("./assets/css/bootstrap.min.css")
require("./assets/css/base.css")

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    }
  ]
})


Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
