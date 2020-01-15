import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Index from "./pages/Index";
import Dashboard from "./pages/Dashboard";
import Settings from "./pages/Settings";
import Service from "./pages/Service";
import Services from "./pages/Services";

require("./assets/css/bootstrap.min.css")
require("./assets/css/base.css")

const routes = [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: Dashboard
    },
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
    component: Service
  }
];

const router = new VueRouter
({
  mode: 'history',
  routes
})

Vue.use(VueRouter);

Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
