import {message} from "ant-design-vue";

const state = {
    chats : [],
    currentChat:{},
    currentMsgs:[],
}

const mutations = {
    SET_CHATS (state,chats) {
        state.chats = chats
    },
    SET_CURRENT_CHAT (state,chat) {
        state.currentChat = chat
    },
    SET_CURRENT_MSGS (state,msgs) {
        state.currentMsgs = msgs
    },
}

const actions = {
    GET_CHATS(context){
        let _this = this
        _this._vm.$http.get('/v1/chats')
            .then(function(data){
                if(data.data.errcode !== 0){
                    message.error(data.data.msg);
                }else{
                    context.commit('SET_CHATS',data.data.data)
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    // GET_CHAT_MESSAGE(context,chat){
    //     let _this = this
    //     _this._vm.$http.get(   `/v1/messages/${chat.id}`)
    //         .then(function(data){
    //             if(data.data.errcode !== 0){
    //                 message.error(data.data.msg);
    //             }else{
    //                 chat.msgs = data.data.data
    //             }
    //             console.log("获取到msg")
    //         })
    //         .catch(function (error) {
    //             console.log(error);
    //         });
    // },
    SET_CHATS(context,chats){
        context.commit('SET_CHATS',chats)
    },
    // SET_CURRENT_CHAT(context,chat){
    //     this.dispatch('GET_CHAT_MESSAGE',chat)
    //     context.commit('SET_CURRENT_CHAT',chat)
    //     console.log("设置聊天")
    // },
    async GET_CHAT_MESSAGE(context, chat) {
        try {
            const response = await this._vm.$http.get(`/v1/rooms/${chat.id}/messages`);
            if (response.data.errcode !== 0) {
                message.error(response.data.msg);
            } else {
                context.commit('SET_CURRENT_MSGS', response.data.data);
            }
            console.log("获取到msg");
        } catch (error) {
            console.log(error);
            throw error; // 抛出错误
        }
    },
    async SET_CURRENT_CHAT(context, chat) {
        try {
            await this.dispatch('GET_CHAT_MESSAGE', chat);
            context.commit('SET_CURRENT_CHAT', chat);
            console.log("设置聊天");
        } catch (error) {
            console.log(error);
            throw error; // 抛出错误
        }
    },
}

export default {
    state,
    mutations,
    actions,
}
