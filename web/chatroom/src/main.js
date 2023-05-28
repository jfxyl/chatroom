import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
Vue.http = Vue.prototype.$http = axios
import 'ant-design-vue/dist/antd.css';

import _const from './config/config.js'
Vue.prototype._const = _const

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
