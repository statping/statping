import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Index from "./pages/Index";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import Settings from "./pages/Settings";
import Services from "./pages/Services";
import Service from "./pages/Service";

require("./assets/css/bootstrap.min.css")
require("./assets/css/base.css")

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
    alias: ['/dashboard/settings', '/dashboard/services', '/dashboard/messages', '/dashboard/groups', '/dashboard/users', '/dashboard/logs', '/dashboard/help']
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  { path: '/logout', redirect: '/' },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings
  },
  {
    path: '/services',
    name: 'Services',
    component: Services
  },
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

Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
