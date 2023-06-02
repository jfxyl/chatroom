import {message} from "ant-design-vue";

const state = {
    rooms : [],
    currentRoom:{},
}

const mutations = {
    SET_ROOMS (state,rooms) {
        state.rooms = rooms
    },
    SET_CURRENT_ROOM (state,room) {
        state.currentRoom = room
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
    SET_CURRENT_ROOM(context,room){
        context.commit('SET_CURRENT_ROOM',room)
    },
}

export default {
    state,
    mutations,
    actions
}
