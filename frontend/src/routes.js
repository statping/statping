import Index from "./components/Pages/Index";
import VueRouter from 'vue-router'

const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    }
  ]
})

export default router
