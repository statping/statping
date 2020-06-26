const Index = () => import('@/pages/Index')
const Dashboard = () => import('@/pages/Dashboard')
const DashboardIndex = () => import('@/components/Dashboard/DashboardIndex')
const DashboardUsers = () => import('@/components/Dashboard/DashboardUsers')
const DashboardServices = () => import('@/components/Dashboard/DashboardServices')
const DashboardMessages = () => import('@/components/Dashboard/DashboardMessages')
const EditService = () => import('@/components/Dashboard/EditService')
const Logs = () => import('@/pages/Logs')
const Settings = () => import('@/pages/Settings')
const Login = () => import('@/pages/Login')
const Service = () => import('@/pages/Service')
const Setup = () => import('@/forms/Setup')
const Incidents = () => import('@/components/Dashboard/Incidents')
const Checkins = () => import('@/components/Dashboard/Checkins')
const Failures = () => import('@/components/Dashboard/Failures')
const NotFound = () => import('@/pages/NotFound')

import VueRouter from "vue-router";
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
    beforeEnter: async (to, from, next) => {
      if (to.matched.some(record => record.meta.requiresAuth)) {
        let tk = await Api.token()
        if (to.path !== '/login' && !tk) {
          next('/login')
          return
        }
        next()
      } else {
        next()
      }
    },
    children: [{
      path: '',
      component: DashboardIndex,
      meta: {
        requiresAuth: true
      }
    },{
      path: 'users',
      component: DashboardUsers,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'services',
      component: DashboardServices,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'create_service',
      component: EditService,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'edit_service/:id',
      component: EditService,
      meta: {
        requiresAuth: true
      }
    },{
      path: 'service/:id/incidents',
      component: Incidents,
      meta: {
        requiresAuth: true
      }
    },{
      path: 'service/:id/checkins',
      component: Checkins,
      meta: {
        requiresAuth: true
      }
    },{
      path: 'service/:id/failures',
      component: Failures,
      meta: {
        requiresAuth: true
      }
    },{
      path: 'messages',
      component: DashboardMessages,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'settings',
      component: Settings,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'logs',
      component: Logs,
        meta: {
            requiresAuth: true
        }
    },{
      path: 'help',
      component: Logs,
        meta: {
            requiresAuth: true
        }
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
  },
  {
    path: '*',
    component: NotFound,
    name: 'NotFound',
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

let CheckAuth = (to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    let item = this.$cookies.get("statping_auth")
    window.console.log(item)
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
