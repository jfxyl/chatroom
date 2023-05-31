import {message} from "ant-design-vue";

const state = {
    rooms : [],
}

const mutations = {
    SET_ROOMS (state,rooms) {
        state.rooms = rooms
    },
}

const actions = {
    GET_ROOMS(context){
        let _this = this
        _this._vm.$http.get('/v1/rooms')
            .then(function(data){
                if(data.data.errcode !== 0){
                    message.error(data.data.msg);
                }else{
                    context.commit('SET_ROOMS',data.data.data)
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    SET_ROOMS(context,rooms){
        context.commit('SET_ROOMS',rooms)
    },
}

export default {
    state,
    mutations,
    actions
}
