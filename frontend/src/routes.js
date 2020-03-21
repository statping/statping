import Index from "./pages/Index";
import Dashboard from "./pages/Dashboard";
import DashboardIndex from "./components/Dashboard/DashboardIndex";
import DashboardUsers from "./components/Dashboard/DashboardUsers";
import DashboardServices from "./components/Dashboard/DashboardServices";
import EditService from "./components/Dashboard/EditService";
import DashboardMessages from "./components/Dashboard/DashboardMessages";
import Logs from './pages/Logs';
import Settings from "./pages/Settings";
import Login from "./pages/Login";
import Service from "./pages/Service";
import VueRouter from "vue-router";
import Setup from "./forms/Setup";

import Api from "./API";

const routes = [
  {
    path: '/setup',
    name: 'Setup',
    component: Setup
  },
  {
    path: '/',
    name: 'Index',
    component: Index,
  },
  {
    path: '/dashboard',
    component: Dashboard,
    meta: {
      requiresAuth: true
    },
    beforeEnter: CheckAuth,
    children: [{
      path: '',
      component: DashboardIndex,
      meta: {
        requiresAuth: true
      }
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
      component: Logs
    },{
      path: 'help',
      component: Logs
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
    scrollBehavior (to, from, savedPosition) {
      if (savedPosition) {
        return savedPosition
      } else {
        return { x: 0, y: 0 }
      }
    },
    routes
})

function CheckAuth(to, from, next) {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    let item = localStorage.getItem("statping_user")
    if (to.path !== '/login' && !item) {
      next('/login')
      return
    }
    const auth = JSON.parse(item)
    if (!auth.token) {
      next('/login')
      return
    }
    next()
  } else {
    next()
  }
}

export default router
