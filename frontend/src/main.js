import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'
import VueObserveVisibility from 'vue-observe-visibility'
import VueClipboard from 'vue-clipboard2'
import VueCookies from 'vue-cookies'
import VueI18n from 'vue-i18n'
import * as Sentry from "@sentry/browser";
import * as Integrations from "@sentry/integrations";
import router from './routes'
import "./mixin"
import "./icons"
import store from './store'
import language from './languages'

const errorReporter = "https://bed4d75404924cb3a799e370733a1b64@sentry.statping.com/3"

const App = () => import(/* webpackChunkName: "index" */ '@/App.vue')

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

Sentry.init({
  dsn: errorReporter,
  integrations: [new Integrations.Vue({Vue, attachProps: true, logErrors: true})],
});

Vue.config.productionTip = process.env.NODE_ENV !== 'production'
Vue.config.devtools = process.env.NODE_ENV !== 'production'
Vue.config.performance = process.env.NODE_ENV !== 'production'

new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
}).$mount('#app')
