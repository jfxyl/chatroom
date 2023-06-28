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
        :afterClose="createRoomAfterCloseHandle"
    >
      <a-form layout="horizontal" :form="createRoomform">
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
        <a-form-item label="头像" hidden>
          <a-input
              v-decorator="[
              'avatar',
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="头像">
          <a-upload
              name="avatar"
              list-type="picture-card"
              class="avatar-uploader"
              :show-upload-list="false"
              :before-upload="beforeUpload"
          >
            <img v-if="previewData" :src="previewData" alt="avatar" />
            <div v-else>
              <a-icon :type="loading ? 'loading' : 'plus'" />
              <div class="ant-upload-text">
                Upload
              </div>
            </div>
          </a-upload>
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal
        v-model="joinRoomModalIsShow"
        title="加入聊天室"
        centered
        width="400"
        @ok="joinRoom"
        :afterClose="joinRoomAfterCloseHandle"
        :footer="null"
    >
      <a-form layout="inline" :form="searchRoomform" @submit="searchRoom">
        <a-form-item >
          <a-input
              v-decorator="[
                'roomid',
                { rules: [{ required: true, message: '请输入房间ID!' }] },
              ]"
              placeholder="房间ID"
          >
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" >
            查找
          </a-button>
        </a-form-item>
      </a-form>
      <div v-if="searchRoomInfo" >
        <div class="roominfo">
          <div class="left">
            <div class="row roomname">
              {{searchRoomInfo.name}}
            </div>
            <div class="row roomtime">
              {{this.$moment(searchRoomInfo.created_at).format('YYYY-MM-DD')}}
            </div>
          </div>
          <div class="right">
            <img class="avatar" :src="searchRoomInfo.avatar" alt="">
          </div>
        </div>
        <a-button type="primary" block @click="joinRoom">加入聊天室</a-button>
      </div>
    </a-modal>
  </div>
</template>

<script>
import Vue from 'vue'
import {Menu, Dropdown, Icon, Modal, Form, Input,Button,Upload, message} from 'ant-design-vue';
import {formErrorPrompt,getContentTypeFromBase64,b64toBlob} from "@/utils/utils";

Vue.use(Modal)
export default {
  name: 'HomeView',
  components:{
    'a-menu':Menu,
    'a-menu-item':Menu.Item,
    'a-dropdown':Dropdown,
    'a-icon':Icon,
    'a-form':Form,
    'a-form-item':Form.Item,
    'a-input':Input,
    'a-upload':Upload,
    'a-button':Button,
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
    beforeUpload(file) {
      const filetypes = [
        "image/jpeg","image/png","image/gif","image/webp","image/bmp"
      ];
      const isAllow = filetypes.includes(file.type);
      if (!isAllow) {
        message.error('不支持的图片格式!');
        return false;
      }
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
        message.error('图片大小不能超过2MB!');
        return false;
      }
      const reader = new FileReader();
      reader.onload = (e) => {
        const url = e.target.result;
        this.showPreview(url);
      };
      reader.readAsDataURL(file);
      return false;
    },
    showPreview(url) {
      this.previewData = url;
    },
    async uploadRoomAvatar(){
      let that = this
      //先尝试上传聊天室图片
      if(that.previewData !== ""){
        const contentType = getContentTypeFromBase64(that.previewData)
        if(contentType){
          const response = await that.$http.get(`/v1/oss/signature?dirtype=room_avatar&content-type=${contentType}`);
          if(response.data.errcode !== 0) {
            message.error(response.data.msg)
          }else{
            const ossUploadUrl = response.data.data.uploadUrl
            var config = {
              method: 'put',
              url: ossUploadUrl,
              headers: {
                'Content-Type': contentType
              },
              data : b64toBlob(that.previewData)
            };
            await that.$http.request(config)
                .then(function () {
                  that.createRoomform.setFieldsValue({
                    avatar: response.data.data.url,
                  });
                })
                .catch(function (error) {
                  console.log(error);
                });
          }
        }
      }
    },
    createRoom(e){
      e.preventDefault();
      this.createRoomform.validateFields(async(err) => {
        if (!err) {
          let that = this
          await that.uploadRoomAvatar()
          that.createRoomform.validateFields(async(err, form) => {
            if (!err) {
              that.$http.post('/v1/rooms',form)
                .then(function(data){
                  if(data.data.errcode !== 0){
                    formErrorPrompt(that.createRoomform,form,data.data.msg)
                  }else{
                    that.$store.dispatch('GET_ROOMS')
                    that.$store.dispatch('GET_CHATS')
                    message.success('创建成功');
                    that.createRoomModalIsShow = false
                  }
                })
                .catch(function (err){
                  console.log(err)
                })
            }
          })
        }
      });
    },
    joinRoom(){
      let that = this
      that.$http.post(`/v1/rooms/${that.searchRoomInfo.id}/join`)
          .then(function(data){
            if(data.data.errcode !== 0){
              message.error(data.data.msg);
            }else{
              that.$store.dispatch('GET_ROOMS')
              that.$store.dispatch('GET_CHATS')
              message.success('操作成功');
              that.joinRoomModalIsShow = false
            }
          })
          .catch(function (err){
            console.log(err)
          })
    },
    searchRoom(e){
      e.preventDefault();
      this.searchRoomform.validateFields((err, form) => {
        if (!err) {
          let that = this
          that.$http.get(`/v1/rooms/${form.roomid}`)
              .then(function(data){
                if(data.data.errcode !== 0){
                  formErrorPrompt(that.createRoomform,form,data.data.msg)
                }else{
                  that.searchRoomInfo = data.data.data
                }
              })
              .catch(function (err){
                console.log(err)
              })
        }
      });
    },
    createRoomAfterCloseHandle(){
      this.createRoomform.resetFields(["name"])
    },
    joinRoomAfterCloseHandle(){
      this.searchRoomform.resetFields(["roomid"])
      this.searchRoomInfo = null
    }
  },
  data(){
    return{
      createRoomModalIsShow:false,
      joinRoomModalIsShow:false,
      createRoomform:this.$form.createForm(this, { name: 'coordinated' }),
      searchRoomform:this.$form.createForm(this, { name: 'coordinated' }),
      searchRoomInfo:null,
      loading: false,
      previewData: '',
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
.roominfo{
  display: flex;
  flex-direction: row;
  background-color: #f0f5fc;
  padding: 12px 24px;
  margin: 10px 0;
  .left{
    flex: 1;
    .row{
      height: 30px;
      line-height: 30px;
    }
    .roomname{
      font-size: 20px;
    }
  }
  .right{
    width:60px;
    height: 60px;
    border-radius: 6px;
    .avatar{
      width:100%;
      height:100%;
    }
  }
}
.avatar-uploader {
  width: 128px;
  height: 128px;
}

.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card img {
  width: 100%;
  height: 100%;
}

</style>
