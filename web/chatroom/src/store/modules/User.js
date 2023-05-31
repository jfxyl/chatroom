import {getUserInfo} from '../../utils/utils.js';

const state = {
    id:null,
    info:null,
    auth:false,
    token:null,
    expired_at:null,
}

const mutations = {
    LOGIN (state){
        state.auth = true
    },
    SET_USER_INFO (state,userInfo) {
        state.info = {
            id : userInfo.id,
            name : userInfo.name,
            nickname : userInfo.nickname,
            avatar : userInfo.avatar,
            gender : userInfo.gender,
            birthday : userInfo.birthday,
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
