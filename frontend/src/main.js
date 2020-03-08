import Vue from 'vue'
import VueRouter from 'vue-router'
import VueApexCharts from 'vue-apexcharts'
import VueObserveVisibility from 'vue-observe-visibility'

import App from '@/App.vue'
import store from './store'
import * as Sentry from '@sentry/browser';
import * as Integrations from '@sentry/integrations';
const errorReporter = "https://bed4d75404924cb3a799e370733a1b64@sentry.statping.com/3"
import router from './routes'
import "./mixin"
import "./icons"

Vue.component('apexchart', VueApexCharts)

Vue.use(VueRouter);
Vue.use(VueObserveVisibility);

Sentry.init({
  dsn: errorReporter,
  integrations: [new Integrations.Vue({Vue, attachProps: true})],
});


Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
