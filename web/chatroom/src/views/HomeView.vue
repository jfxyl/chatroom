<template>
  <div class="container">
    <div class="top">
      <div class="toolbar">
        <div class="search"></div>
        <a-dropdown placement="bottomCenter">
          <a-icon type="plus" class="add" />
          <a-menu slot="overlay">
            <a-menu-item @click="createRoomModalShow">
              创建聊天室
            </a-menu-item>
            <a-menu-item @click="joinRoomModalShow">
              加入聊天室
            </a-menu-item>
          </a-menu>
        </a-dropdown>
      </div>
    </div>
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
    <a-modal
        v-model="createRoomModalIsShow"
        title="创建聊天室"
        centered
        @ok="createRoom"
    >
      <a-form  :form="createRoomform">
        <a-form-item label="房间名">
          <a-input
              placeholder="请输入房间名"
              allow-clear
              v-decorator="[
              'name',
              { rules: [{ required: true, message: '请输入房间名' }] },
            ]"
          ></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>

</template>

<script>
import Vue from 'vue'
import {Menu, Dropdown, Icon, Modal, Form, Input, message} from 'ant-design-vue';
import {formErrorPrompt} from "@/utils/utils";
Vue.use(Modal)
export default {
  name: 'HomeView',
  components:{
    'a-menu':Menu,
    'a-menu-item':Menu.Item,
    'a-dropdown':Dropdown,
    'a-icon':Icon,
    // 'a-modal':Modal,
    'a-form':Form,
    'a-form-item':Form.Item,
    'a-input':Input,
    // 'a-button':Button,
  },
  computed: {
    currentMenu(){
      return this.$store.state.menu.currentMenu
    },
  },
  methods: {
    setMenu(menu){
      this.$store.dispatch('SET_CURRENT_MENU',menu)
    },
    createRoomModalShow(){
       this.createRoomModalIsShow = true
    },
    joinRoomModalShow(){
      this.joinRoomModalIsShow = true
    },
    createRoom(e){
      e.preventDefault();
      this.createRoomform.validateFields((err, form) => {
        if (!err) {
          let that = this
          that.$http.post('/v1/rooms',form)
            .then(function(data){
              if(data.data.errcode !== 0){
                formErrorPrompt(that.createRoomform,form,data.data.msg)
              }else{
                that.$store.dispatch('GET_ROOMS')
                message.success('创建成功');
                that.createRoomModalIsShow = false
              }
            })
            .catch(function (err){
              console.log(err)
            })
        }
      });
    }
  },
  data(){
    return{
      createRoomModalIsShow:false,
      joinRoomModalIsShow:false,
      createRoomform:this.$form.createForm(this, { name: 'coordinated' }),
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
    justify-content: center;
    background-color: #438be5;
    -webkit-app-region: drag;
    .toolbar{
      display: flex;
      -webkit-app-region: no-drag;
      .search{
        width: 30vw;
        padding: 2px 0;
        display: flex;
        justify-content: center;
        align-items: center;
        border-radius: 6px;
        background-color: #5eacf1;
        color: #d2e3f8;
        cursor: pointer;
      }
      .add{
        font-size: 20px;
        color: #fff;
        margin-left: 10px;
      }
    }
  }
  .center{
    display: flex;
    flex-direction: row;
    flex: 1;
    overflow: hidden;
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
