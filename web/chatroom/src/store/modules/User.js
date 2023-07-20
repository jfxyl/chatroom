import {getUserInfo} from '../../utils/utils.js';

const state = {
    info:{},
    auth:false,
    token:null,
    expired_at:null,
}

const mutations = {
    LOGIN (state){
        state.auth = true
    },
    LOGOUT (state){
        state.auth = false
        localStorage.removeItem('userInfo')
        localStorage.removeItem('token')
        localStorage.removeItem('expired_at')
    },
    SET_USER_INFO (state,userInfo) {
        state.info = {
            id : userInfo.id,
            name : userInfo.name,
            nickname : userInfo.nickname,
            avatar : userInfo.avatar,
            gender : userInfo.gender,
            birthday : userInfo.birthday,
            created_at : userInfo.created_at,
        }
        localStorage.setItem('userInfo',JSON.stringify(userInfo))
    },
    REFRESH_TOKEN (state,tokenInfo){
        var token = 'Bearer ' + tokenInfo.token
        var expired_at = tokenInfo.expired_at * 1000
        state.token = token
        state.expired_at = expired_at
        localStorage.setItem('token',token)
        localStorage.setItem('expired_at',expired_at)
        console.log(localStorage.getItem('token'))
    },
}

const actions = {
    LOGIN (context,userInfo){
        context.commit('LOGIN')
        this.dispatch('SET_USER_INFO',userInfo)
    },
    LOGOUT (context){
        context.commit('LOGOUT')
    },
    SET_USER_INFO (context,userInfo) {
        // do something async
        context.commit('SET_USER_INFO',userInfo)
    },
    REFRESH_TOKEN (context,tokenInfo) {
        context.commit('REFRESH_TOKEN',tokenInfo)
    },
    GET_USER (){
        getUserInfo(this)
    }
}

export default {
    state,
    mutations,
    actions
}
