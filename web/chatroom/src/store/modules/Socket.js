const state = {
    socket:null,
}

const mutations = {
    SET_SOCKET (state,socket) {
        state.socket = socket
    },
}

const actions = {
    CONN_SOCKET(context){
        let that = this
        const socket = new WebSocket('ws://127.0.0.1:8081/v1/ws?Authorization=' + localStorage.getItem('token'));
        // WebSocket 连接建立成功的回调函数
        socket.onopen = function() {
            console.log('WebSocket 连接已建立');
            // 可以在这里发送认证信息或其他初始化操作
        };

        // WebSocket 接收到消息的回调函数
        socket.onmessage = function(event) {
            const message = JSON.parse(event.data);
            console.log('收到消息:', message);
            // 在这里处理接收到的消息
            console.log(that)
            console.log(that.$route)
            console.log(that.$store)
            console.log(message.receiver_id == that.state.chat.currentChat.id)
            console.log(message.sender_id , that.state.user.info.id)
            if(message.receiver_id == that.state.chat.currentChat.id && message.sender_id != that.state.user.info.id){
                that.state.chat.currentMsgs.push(message)
            }else{
                for(let i=0;i<that.state.chat.chats.length;i++){
                    if(that.state.chat.chats[i].id == message.receiver_id){
                        that.state.chat.chats[i].unread_count += 1
                        break
                    }
                }
            }
        };

        // WebSocket 连接关闭的回调函数
        socket.onclose = function() {
            console.log('WebSocket 连接已关闭');
        };

        // WebSocket 连接出错的回调函数
        socket.onerror = function(error) {
            console.error('WebSocket 连接出错:', error);
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
}

export default {
    state,
    mutations,
    actions
}
