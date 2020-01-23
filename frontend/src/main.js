import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'

import App from '@/App.vue'
import store from '@/store'

import {library} from '@fortawesome/fontawesome-svg-core'
import {fas} from '@fortawesome/fontawesome-free-solid';
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'
import DashboardIndex from "./components/Dashboard/DashboardIndex";
import DashboardUsers from "./components/Dashboard/DashboardUsers";
import DashboardServices from "./components/Dashboard/DashboardServices";
import DashboardMessages from "./components/Dashboard/DashboardMessages";
import Settings from "./pages/Settings";
import EditService from "./components/Dashboard/EditService";
import Dashboard from "./pages/Dashboard";
import Index from "./pages/Index";
import Login from "./pages/Login";
import Service from "./pages/Service";

library.add(fas)

Vue.component('apexchart', VueApexCharts)
Vue.component('font-awesome-icon', FontAwesomeIcon)

require("@/assets/css/bootstrap.min.css")
require("@/assets/css/base.css")


// require("./assets/js/bootstrap.min")
// require("./assets/js/flatpickr")
// require("./assets/js/inputTags.min")
// require("./assets/js/rangePlugin")
// require("./assets/js/sortable.min")

const routes = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    children: [{
      path: '',
      component: DashboardIndex
    },{
      path: 'users',
      component: DashboardUsers
    },{
      path: 'services',
      component: DashboardServices
    },{
      path: 'create_service',
      component: EditService
    },{
      path: 'edit_service/:id',
      component: EditService
    },{
      path: 'messages',
      component: DashboardMessages
    },{
      path: 'settings',
      component: Settings
      },{
      path: 'logs',
      component: DashboardUsers
    },{
      path: 'help',
      component: DashboardUsers
    }]
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  { path: '/logout', redirect: '/' },
  {
    path: '/service/:id',
    name: 'Service',
    component: Service,
    props: true
  }
];

const router = new VueRouter({
  mode: 'history',
  routes
})

Vue.use(VueRouter);
Vue.use(require('vue-moment'));

Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
