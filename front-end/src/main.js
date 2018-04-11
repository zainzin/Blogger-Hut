// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VeeValidate from 'vee-validate'
import VueNoty from 'vuejs-noty'
import Vues from 'vuex'
import App from './App'
import router from './router'

Vue.config.productionTip = false

Vue.use(VeeValidate)
Vue.use(VueNoty, {
  progressBar: true,
  timeout: 5000,
  layout: 'topRight',
  theme: 'sunset'
})
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
