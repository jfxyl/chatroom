<template>
  <div v-if="currentRoom.id" class="container">
    <div class="avatar">
      <img class="img" :src="currentRoom.avatar" alt="">
    </div>
    <div class="name">{{currentRoom.name}}</div>
    <button class="button" @click="toChat">进入聊天室</button>
  </div>
</template>

<script>
export default {
  name: 'RoomPanel',
  components:{
  },
  data(){
    return{

    }
  },
  computed: {
    currentRoom() {
      return this.$store.state.room.currentRoom
    },
    chats() {
      return this.$store.state.chat.chats
    },
  },
  methods:{
    toChat(){
      //循环chats,找到currentRoom.id==chat.roomId的chat,将其赋值给currentChat
      for(let i=0;i<this.chats.length;i++){
        if(this.chats[i].id==this.currentRoom.id){
          this.$store.dispatch('SET_CURRENT_CHAT',this.chats[i])
          break
        }
      }
      this.$router.push({path:'/chat'})
    },
  }
}
</script>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%; /* 设置容器高度，使其铺满整个视口 */
  flex:1;
  .avatar {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    background-color: #ccc;
    .img{
      width: 100%;
      height: 100%;
    }
  }

  .name {
    margin-top: 20px;
    font-size: 16px;
    font-weight: bold;
  }

  .button {
    margin-top: 20px;
    padding: 10px 20px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
}


</style>
