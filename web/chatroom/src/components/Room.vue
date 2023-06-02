<template>
  <div v-if="currentRoom.id" class="container">
    <div class="avatar">
      <img :src="currentRoom.avatar" alt="">
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
  },
  methods:{
    toChat(){
      console.log(this.currentRoom)
    },
  }
}
</script>

<style scoped>
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
