import Vue from 'vue'
import Vuex from 'vuex'
import userModule from './modules/User'
import chatModule from './modules/Chat'
import roomModule from './modules/Room'
import menuModule from './modules/Menu'
import socketModule from './modules/Socket'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user: userModule,
    chat: chatModule,
    room: roomModule,
    menu: menuModule,
    socket: socketModule,
  }
})
