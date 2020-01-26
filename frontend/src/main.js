import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'

import App from '@/App.vue'
import store from './store'

import {library} from '@fortawesome/fontawesome-svg-core'
import {fas} from '@fortawesome/fontawesome-free-solid';
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'
import router from './routes'
import "./mixin"

library.add(fas)

Vue.component('apexchart', VueApexCharts)
Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.use(VueRouter);
Vue.use(require('vue-moment'));

Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
