<template>
  <div class="chats">
    <a-dropdown class="chat" v-for="(chat,name,index) in chats" :key="index" :trigger="['contextmenu']" >
      <div @click="current(chat)" :class="{active:currentChat.id === chat.id}">
        <div class="avatar">
          <img class="img" :src="chat.avatar" alt="">
        </div>
        <div class="info" >
          <div class="top">{{chat.name}}</div>
          <!--          <div class="bottom">1000人</div>-->
        </div>
      </div>
      <a-menu slot="overlay">
        <a-menu-item key="1" @click="toggleTop(chat)">置顶</a-menu-item>
        <a-menu-item key="2" @click="toggleDisturb(chat)">消息免打扰</a-menu-item>
        <a-menu-divider />
<!--        <a-menu-item key="3" @click="">退出聊天室</a-menu-item>-->
      </a-menu>
    </a-dropdown>
  </div>
</template>

<script>
import { Menu,Dropdown  } from 'ant-design-vue';

export default {
  name: 'ChatsPanel',
  components:{
    'a-menu':Menu,
    'a-menu-item':Menu.Item,
    'a-menu-divider':Menu.Divider,
    'a-dropdown':Dropdown,
  },
  data(){
    return{

    }
  },
  computed: {
    chats() {
      return this.$store.state.chat.chats
    },
    currentChat() {
      return this.$store.state.chat.currentChat
    },
  },
  methods:{
    current(chat){
      this.$store.dispatch('SET_CURRENT_CHAT',chat)
    },
    toggleTop(chat){
      console.log(chat)
    },
    toggleDisturb(chat){
      console.log(chat)
    }
  }
}
</script>

<style scoped>
.chats{
  width:350px;
  padding: 10px;
  height:100%;
  max-height: 100%;
  border-right: 1px solid #ccc;
  overflow:auto;
  overflow-x:hidden;
  cursor: default;
  &::-webkit-scrollbar {
    width: 6px; /* 设置滚动条宽度 */
  }
  &::-webkit-scrollbar-thumb {
    background-color: rgba(0, 0, 0, 0.2); /* 设置滚动条颜色 */
    border-radius: 4px;
  }

  &::-webkit-scrollbar-thumb:hover {
    background-color: rgba(0, 0, 0, 0.4); /* 设置滚动条悬停时的颜色 */
  }
  .active{
    background-color: #f3f8fe;
  }
  .chat{
    display: flex;
    align-items: center;
    padding: 11px 16px;
    border-radius: 8px;
    .avatar{
      width: 36px;
      height: 36px;
      line-height: 36px;
      font-size: 18px;
      min-width: 36px;
      .img{
        width:100%;
        height:100%;
        border-radius: 6px;
      }
    }
    .info{
      display: flex;
      flex: auto;
      flex-direction: column;
      margin-left:10px;

      .top{
        display: flex;
        flex-direction: row;
        .name{
          font-size: 14px;
          text-align: left;
          flex: 1;
        }
        .time{
          text-align: left;
          font-size: 12px;
          color: #999;
        }
      }
      .bottom{
        font-size: 12px;
        color: #999;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 1;
        overflow: hidden;
        word-break: break-all;
        max-width: 90%;
        text-align: left;
      }

    }

  }
  .current{
    background-color: #f3f8fe;
    border-radius: 8px;
  }
}
</style>
