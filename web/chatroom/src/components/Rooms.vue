<template>
  <div class="rooms">
    <a-dropdown class="room" v-for="(room,name,index) in rooms" :key="index" :trigger="['contextmenu']" >
      <div @click="current(room)" :class="{active:currentRoom.id === room.id}">
        <div class="avatar">
          <img class="img" :src="room.avatar" alt="">
        </div>
        <div class="info" >
          <div class="top">{{room.name}}</div>
<!--          <div class="bottom">1000人</div>-->
        </div>
      </div>
      <a-menu slot="overlay">
        <a-menu-item key="1" @click="toChat(room)">发消息</a-menu-item>
        <a-menu-item key="2" @click="setAlias(room)">备注设置</a-menu-item>
        <a-menu-divider />
        <a-menu-item key="3" @click="quitRoom(room)">退出聊天室</a-menu-item>
      </a-menu>
    </a-dropdown>
  </div>
</template>

<script>
import { Menu,Dropdown  } from 'ant-design-vue';
export default {
  name: 'RoomsPanel',
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
    rooms() {
      return this.$store.state.room.rooms
    },
    currentRoom() {
      return this.$store.state.room.currentRoom
    },
  },
  methods:{
    current(room){
      this.$store.dispatch('SET_CURRENT_ROOM',room)
    },
    toChat(room){
      console.log(room)
    },
    setAlias(room){
      console.log(room)
    },
    quitRoom(room){
      console.log(room)
    }
  }
}
</script>

<style scoped>
.rooms{
  width:350px;
  padding: 10px;
  height:100%;
  max-height: 100%;
  border-right: 1px solid #ccc;
  overflow:auto;
  overflow-x:hidden;
  .active{
    background-color: #f3f8fe;
  }
  .room{
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
        font-size: 16px;
        text-align: left;
        flex: 1;
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
