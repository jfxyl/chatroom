const state = {
    socket:null,
    timer: null,
}

const mutations = {
    SET_SOCKET (state,socket) {
        state.socket = socket
    },
    SET_TIMER(state, timer) {
        state.timer = timer;
    },
    CLEAR_TIMER(state) {
        clearTimeout(state.timer);
    },
}

const actions = {
    CONN_SOCKET(context,that){
        that = that || this
        const socket = new WebSocket('ws://127.0.0.1:8081/v1/ws?Authorization=' + localStorage.getItem('token'));
        // WebSocket 连接建立成功的回调函数
        socket.onopen = function() {
            console.log('WebSocket 连接已建立');
            context.commit("CLEAR_TIMER");
        };
        // WebSocket 接收到消息的回调函数
        socket.onmessage = function(event) {
            const message = JSON.parse(event.data);
            console.log('收到消息:', message);
            console.log('that',that)
            // 在这里处理接收到的消息

            //将其他人发的消息增加未读数量
            for(let i=0;i<that.state.chat.chats.length;i++){
                if(message.sender_id != that.state.user.info.id && message.receiver_id == that.state.chat.chats[i].id && message.msg_type != that._vm._const.TypeNotice){
                    that.state.chat.chats[i].unread_count = (that.state.chat.chats[i].unread_count || 0) + 1;
                    break
                }
            }

            if(message.receiver_id == that.state.chat.currentChat.id){
                if(message.sender_id != that.state.user.info.id){
                    that.state.chat.currentMsgs.push(message)
                    console.log('message.msg_type == that._vm._const.TypeNotice',message.msg_type == that._vm._const.TypeNotice)
                    console.log('message.operate == that._vm._const.OperateJoinRoom',message.operate, that._vm._const.OperateJoinRoom)
                    if(message.msg_type == that._vm._const.TypeNotice){
                        if(message.operate == that._vm._const.OperateJoinRoom){
                            var exists = that.state.chat.currentChat.users.some(obj => obj.id == message.sender_id);
                            if(!exists){
                                that.state.chat.currentChat.users.push(message.sender)
                            }
                        }else if(message.operate == that._vm._const.OperateQuitRoom){
                            that.state.chat.currentChat.users = that.state.chat.currentChat.users.filter(obj => obj.id != message.sender_id);
                        }
                    }
                }else{
                    for(let i=0;i<that.state.chat.currentMsgs.length;i++){
                        if(that.state.chat.currentMsgs[i].id == message.id){
                            that.state.chat.currentMsgs[i].readers = message.readers
                            that.state.chat.currentMsgs[i].unreaders = message.unreaders
                            break
                        }
                    }
                }
            }
        };
        // WebSocket 连接关闭的回调函数
        socket.onclose = function() {
            console.log('WebSocket 连接已关闭');
            actions.RECONNECT_SOCKET(context,that);
        };
        // WebSocket 连接出错的回调函数
        socket.onerror = function(error) {
            console.error('WebSocket 连接出错:', error);
            // actions.RECONNECT_SOCKET(context);
        };
        // 将 WebSocket 实例保存到 Vue 实例或全局对象中，以便在其他组件中访问
        context.commit('SET_SOCKET',socket)
    },
    CLOSE_SOCKET(context){
        if(this.$store.state.socket){
            this.$store.state.socket.close();
            context.commit('SET_SOCKET',null)
        }
    },
    RECONNECT_SOCKET(context,that) {
        console.log('RECONNECT_SOCKET')
        const timer = setTimeout(function() {
            console.log("Reconnecting WebSocket...");
            actions.CONN_SOCKET(context,that);
        }, 2000);
        context.commit("SET_TIMER", timer);
    },
}

export default {
    state,
    mutations,
    actions
}
