const Index = () => import(/* webpackChunkName: "index" */ '@/pages/Index')
const Dashboard = () => import(/* webpackChunkName: "dashboard" */ '@/pages/Dashboard')
const DashboardIndex = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/DashboardIndex')
const DashboardUsers = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/DashboardUsers')
const DashboardServices = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/DashboardServices')
const DashboardMessages = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/DashboardMessages')
const EditService = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/EditService')
const Logs = () => import(/* webpackChunkName: "dashboard" */ '@/pages/Logs')
const Help = () => import(/* webpackChunkName: "dashboard" */ '@/pages/Help')
const Settings = () => import(/* webpackChunkName: "dashboard" */ '@/pages/Settings')
const Login = () => import(/* webpackChunkName: "index" */ '@/pages/Login')
const Service = () => import(/* webpackChunkName: "index" */ '@/pages/Service')
const Setup = () => import(/* webpackChunkName: "index" */ '@/forms/Setup')
const Incidents = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Incidents')
const Checkins = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Checkins')
const Failures = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Failures')
const NotFound = () => import(/* webpackChunkName: "index" */ '@/pages/NotFound')
const Importer = () => import(/* webpackChunkName: "index" */ '@/components/Dashboard/Importer')

import VueRouter from "vue-router";
import Api from "./API";
import store from "./store"

const Loading = {
  template: '<div class="jumbotron">LOADING</div>'
}

const routes = [
  {
    path: '/setup',
    name: 'Setup',
    component: Setup,
    meta: {
      title: 'Statping Setup',
    }
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
      requiresAuth: true,
      title: 'Statping - Dashboard',
    },
    beforeEnter: async (to, from, next) => {
      if (to.matched.some(record => record.meta.requiresAuth)) {
        if (to.path !== '/login') {
          if(store.getters.loggedIn) {
            next()
            return
          }
          const token = $cookies.get('statping_auth')
          if (!token) {
            next('/login')
            return
          }
          try {
            const jwt = await Api.check_token(token)
            store.commit('setAdmin', jwt.admin)
            if (jwt.admin) {
              store.commit('setLoggedIn', true)
              store.commit('setUser', true)
            } else {
              store.commit('setLoggedIn', false)
              next('/login')
              return
            }
          } catch (e) {
            console.error(e)
            next('/login')
            return
          }
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
        requiresAuth: true,
        title: 'Statping - Dashboard',
      }
    },{
      path: 'users',
      component: DashboardUsers,
      loading: Loading,
        meta: {
            requiresAuth: true,
          title: 'Statping - Users',
        }
    },{
      path: 'services',
      component: DashboardServices,
        meta: {
            requiresAuth: true,
          title: 'Statping - Services',
        }
    },{
      path: 'create_service',
      component: EditService,
        meta: {
            requiresAuth: true,
          title: 'Statping - Create Service',
        }
    },{
      path: 'edit_service/:id',
      component: EditService,
      meta: {
        requiresAuth: true,
        title: 'Statping - Edit Service',
      }
    },{
      path: 'service/:id/incidents',
      component: Incidents,
      meta: {
        requiresAuth: true,
        title: 'Statping - Incidents',
      }
    },{
      path: 'service/:id/checkins',
      component: Checkins,
      meta: {
        requiresAuth: true,
        title: 'Statping - Checkins',
      }
    },{
      path: 'service/:id/failures',
      component: Failures,
      meta: {
        requiresAuth: true,
        title: 'Statping - Service Failures',
      }
    },{
      path: 'messages',
      component: DashboardMessages,
        meta: {
            requiresAuth: true,
          title: 'Statping - Messages',
        }
    },{
      path: 'settings',
      component: Settings,
        meta: {
            requiresAuth: true,
          title: 'Statping - Settings',
        }
    },{
      path: 'logs',
      component: Logs,
        meta: {
            requiresAuth: true,
          title: 'Statping - Logs',
        }
    },{
      path: 'help',
      component: Help,
        meta: {
            requiresAuth: true,
          title: 'Statping - Help',
        }
    },{
      path: 'import',
      component: Importer,
      meta: {
        requiresAuth: true,
        title: 'Statping - Import',
      }
    }]
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      title: 'Statping - Login',
    }
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

router.beforeEach((to, from, next) => {
  const nearestWithTitle = to.matched.slice().reverse().find(r => r.meta && r.meta.title);
  const nearestWithMeta = to.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
  const previousNearestWithMeta = from.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);
  if (nearestWithTitle) document.title = nearestWithTitle.meta.title;
  Array.from(document.querySelectorAll('[data-vue-router-controlled]')).map(el => el.parentNode.removeChild(el));
  if (!nearestWithMeta) return next();
  nearestWithMeta.meta.metaTags.map(tagDef => {
    const tag = document.createElement('meta');
    Object.keys(tagDef).forEach(key => {
      tag.setAttribute(key, tagDef[key]);
    });
    tag.setAttribute('data-vue-router-controlled', '');
    return tag;
  })
    .forEach(tag => document.head.appendChild(tag));
  next();
});

export default router
