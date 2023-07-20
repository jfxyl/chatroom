import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
Vue.http = Vue.prototype.$http = axios
import 'ant-design-vue/dist/antd.css';

import moment from 'moment'
Vue.prototype.$moment = moment;


import _config from './config/config.js'
Vue.prototype._config = _config
import _const from './constant/constant.js'
Vue.prototype._const = _const
import utils from './utils/utils.js'
Vue.prototype.utils = utils

Vue.prototype._genderMap = {
  0: '未知',
  1: '男',
  2: '女'
}

Vue.config.productionTip = false

Vue.prototype.globalClick = function (callback) { //页面全局点击
  document.addEventListener('click',callback);
}

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
