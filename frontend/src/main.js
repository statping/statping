import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'
import VueObserveVisibility from 'vue-observe-visibility'
import VueClipboard from 'vue-clipboard2'
import VueCookies from 'vue-cookies'

import App from '@/App.vue'
import store from './store'

Vue.component('apexchart', VueApexCharts)

Vue.use(VueClipboard);
Vue.use(VueRouter);
Vue.use(VueObserveVisibility);
Vue.use(VueCookies);
Vue.$cookies.config('3d')

import router from './routes'
import "./mixin"
import "./icons"


Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
