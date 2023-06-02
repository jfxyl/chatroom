<template>
  <div class="container">
    <div class="top"></div>
    <div class="center">
      <div class="menu">
        <div class="menu-item avatar-item">
          <img  class="avatar" src="https://web.rentsoft.cn/static/media/login_bg.e42640a5.png" alt="">
        </div>
        <div class="menu-item" @click="setMenu('chat')">
          <router-link to="/chat">
            <svg class="icon" aria-hidden="true">
              <use :xlink:href="currentMenu == 'chat'?'#icon-chat-active':'#icon-chat'"></use>
            </svg>
          </router-link>
        </div>
        <div class="menu-item" @click="setMenu('room')">
          <router-link to="/room">
            <svg class="icon" aria-hidden="true">
              <use :xlink:href="currentMenu == 'room'?'#icon-friends_active':'#icon-friends'"></use>
            </svg>
          </router-link>
        </div>
      </div>
      <div class="right-container">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomeView',
  components:{
    // 'a-icon':Icon
  },
  computed: {
    currentMenu(){
      return this.$store.state.menu.currentMenu
    },
  },
  methods: {
    setMenu(menu){
      this.$store.dispatch('SET_CURRENT_MENU',menu)
    }
  },
  data(){
    return{
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
  .top{
    display: flex;
    align-items: center;
    max-height: 42px;
    min-height: 42px;
    justify-content: space-between;
    background-color: #438be5;
    -webkit-app-region: drag;
  }
  .center{
    display: flex;
    flex-direction: row;
    flex: auto;
    .menu{
      background-color: #f4f4f4!important;
      display: flex;
      flex-direction: column;
      flex: 0 0 60px;
      max-width: 60px;
      min-width: 60px;
      width: 60px;
      padding-top: 24px;
      .avatar-item{
        margin-bottom: 20px;
        padding: 10px;

      }
      .menu-item{
        height:60px;
        width: 60px;
        .avatar{
          height:100%;
          width: 100%;
          border-radius: 5px;
          border: 1px solid #ccc;
        }
        .icon {
          display:block;
          width: 35px;
          height: 35px;
          margin:0 auto;
          fill: currentColor;
          overflow: hidden;
        }
      }
    }
    .right-container{
      display: flex;
      flex-direction: row;
      flex:1;
    }
  }
}
</style>
