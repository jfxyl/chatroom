import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
Vue.http = Vue.prototype.$http = axios
import 'ant-design-vue/dist/antd.css';

import _const from './config/config.js'
Vue.prototype._const = _const
import utils from './utils/utils.js'
Vue.prototype.utils = utils

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App),
  created(){
    //请求拦截器
    this.$http.interceptors.request.use(config => {
      // console.log('请求拦截器')
      // 获取token
      // let token = this.$store.state.UserInfo.token?this.$store.state.UserInfo.token:localStorage.getItem('token')
      let token = localStorage.getItem('token')
      if(token){
        config.headers.Authorization = token
      }
      return config
    }, error => {
      return Promise.reject(error)
    })
    // 自定义的 axios 响应拦截器
    this.$http.interceptors.response.use((response) => {
      // console.log('响应拦截器')
      if(response.data.errcode == 1){
        if(this.$route.path !== '/login'){
          this.$router.push('/login');
        }
      }else if(response.headers.authorization){
        // 如果 header 中存在 token，那么触发 refreshToken 方法，替换本地的 token
        // this.$store.dispatch('REFRESH_TOKEN', response.headers.authorization)
      }
      return response
    }, (error) => {

      return Promise.reject(error)
    })
  }
}).$mount('#app')
