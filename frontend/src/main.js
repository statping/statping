import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'
import VueObserveVisibility from 'vue-observe-visibility'
import VueClipboard from 'vue-clipboard2'
import VueCookies from 'vue-cookies'
import VueI18n from 'vue-i18n'
import router from './routes'
import "./mixin"
import "./icons"
const App = () => import('@/App.vue')
import store from './store'
import language from './languages'

Vue.component('apexchart', VueApexCharts)

Vue.use(VueClipboard);
Vue.use(VueRouter);
Vue.use(VueObserveVisibility);
Vue.use(VueCookies);
Vue.use(VueI18n);

const i18n = new VueI18n({
  fallbackLocale: "en",
  messages: language
});

Vue.$cookies.config('3d')

Vue.config.productionTip = false
new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
}).$mount('#app')
