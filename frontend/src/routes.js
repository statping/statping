import Index from "./pages/Index";
import VueRouter from 'vue-router'
import Dashboard from "./pages/Dashboard";
import Settings from "./pages/Settings";

const router = new VueRouter({
  routes: [
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
    }
  ]
})

export default router
