// const state = {
//     chats : {},
//     currentChatId : 1,
//     currentChat : {}
// }
//
// const mutations = {
//     SET_CHATS (state,chats) {
//         state.chats = chats
//     },
//     SET_CURRENT_CHAT_ID(state,chatId) {
//         state.currentChatId = chatId
//     },
//     SET_CURRENT_CHAT(state,chat) {
//         state.currentChat = chat
//     },
// }
//
// const actions = {
//     GET_MSGS(context,chat){
//
//     },
//     GET_CHATS(context){
//         let _this = this
//         _this._vm.$http.get(_this._vm._const.API_URL+'/chats')
//             .then(function (data) {
//                 if(data.data.code == 1000){
//                     let chatObj = _this._vm._function.arr2obj(data.data.data,'unique_id')
//                     chatObj = _this._vm._function.chatObjSort(chatObj)
//                     context.commit('SET_CHATS',chatObj)
//                 }else{
//                     _this._vm.$layer.msg(data.data.msg,{
//                         time:2
//                     });
//                 }
//             })
//             .catch(function (error) {
//                 console.log(error);
//             });
//     },
//     SET_CHATS(context,chats){
//         context.commit('SET_CHATS',chats)
//     },
//     SET_CURRENT_CHAT_ID(context,chatId) {
//         context.commit('SET_CURRENT_CHAT_ID',chatId)
//     },
//     SET_CURRENT_CHAT(context,chat) {
//         context.commit('SET_CURRENT_CHAT',chat)
//         this.dispatch('GET_MSGS',chat)
//     }
// }
//
// export default {
//     state,
//     mutations,
//     actions
// }
