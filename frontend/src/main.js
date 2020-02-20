import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'

import App from '@/App.vue'
import store from './store'

import router from './routes'
import "./mixin"
import "./icons"

Vue.component('apexchart', VueApexCharts)

Vue.use(VueRouter);

Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
