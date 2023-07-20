<template>
  <div class="container">
    <div class="top">
      <div class="toolbar">
        <div class="search">疾风夕颜</div>
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
        <a-popover v-model="userSettingVisible" trigger="click" placement="rightTop">
          <template slot="content">
            <div class="tools">
              <div class="item" @click="userInfoModalShow"><span>我的信息</span><a-icon type="right" /></div>
              <div class="item" @click="updatePasswordModalShow"><span>修改密码</span><a-icon type="right" /></div>
              <div class="item" @click="logout"><span>退出登录</span><a-icon type="right" /></div>
            </div>
          </template>
          <template slot="title">
            <div class="userinfo">
              <div class="avatar">
                <img :src="userInfo.avatar" alt="">
              </div>
              <div>
                {{userInfo.name}}
              </div>
            </div>
          </template>
          <div class="menu-item avatar-item">
            <img :src="userInfo.avatar" alt="">
          </div>
        </a-popover>
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
        v-model="createRoomModaVisible"
        title="创建聊天室"
        centered
        @ok="createRoom"
        :afterClose="createRoomAfterCloseHandle"
    >
      <a-form layout="horizontal" :form="createRoomForm">
        <a-form-item label="房间名" :label-col="{span: 4}" :wrapper-col="{span: 14}">
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
        <a-form-item label="头像" :label-col="{span: 4}" :wrapper-col="{span: 14}">
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
        v-model="joinRoomModalVisible"
        title="加入聊天室"
        centered
        width="400"
        @ok="joinRoom"
        :afterClose="joinRoomAfterCloseHandle"
        :footer="null"
    >
      <a-form layout="inline" :form="searchRoomForm" @submit="searchRoom">
        <a-form-item>
          <a-input
              v-decorator="[
                'content',
                {
                  rules: [{ required: true, message: '请输入聊天室信息！' }],
                },
              ]"
              placeholder="聊天室信息"
              style="width: 100%"
          >
            <a-select
                slot="addonBefore"
                v-decorator="['field', { initialValue: 'id' }]"
                style="width: 120px"
            >
              <a-select-option value="id">
                聊天室ID
              </a-select-option>
              <a-select-option value="name">
                聊天室名称
              </a-select-option>
            </a-select>
          </a-input>
        </a-form-item>
<!--        <a-form-item >-->
<!--          <a-input-->
<!--              v-decorator="[-->
<!--                'roomid',-->
<!--                { rules: [{ required: true, message: '请输入聊天室ID!' }] },-->
<!--              ]"-->
<!--              placeholder="聊天室ID"-->
<!--          >-->
<!--          </a-input>-->
<!--        </a-form-item>-->
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
    <a-modal
        v-model="userInfoVisible"
        title=""
        centered
        width="400"
        :footer="null"
        :mask=false
        :closable=false
        :body-style="{padding:0}"
    >
      <div class="card-content">
        <div class="card-title">
          <div class="avatar">
            <img :src="userInfo.avatar" alt="">
          </div>
          <div class="info">
            <div class="nickname">{{userInfo.nickname}}</div>
            <div class="id">{{userInfo.id}}</div>
          </div>
        </div>
        <div class="card-col">
          <div class="col-title">个人信息</div>
          <div class="card-row">
            <span class="label">用户名</span>
            <span class="content">{{userInfo.name}}</span>
          </div>
          <div class="card-row">
            <span class="label">昵称</span>
            <span class="content">{{userInfo.nickname}}</span>
          </div>
          <div class="card-row">
            <span class="label">性别</span>
            <span class="content">{{_genderMap[userInfo.gender]}}</span>
          </div>
          <div class="card-row">
            <span class="label">生日</span>
            <span class="content">{{userInfo.birthday}}</span>
          </div>
          <div class="card-row">
            <span class="label">注册时间</span>
            <span class="content">{{userInfo.created_at}}</span>
          </div>
        </div>
        <div class="footer" >
          <a-button type="primary" @click="updateUserInfoModalShow()" block>
            编辑资料
          </a-button>
        </div>
      </div>
    </a-modal>
    <a-modal
        v-model="updateUserInfoVisible"
        title="编辑资料"
        centered
        @ok="updateUserInfo"
        :afterClose="updateUserInfoAfterCloseHandle"
    >
      <a-form layout="horizontal" :form="updateUserInfoForm" @submit="updateUserInfo">
        <a-form-item label="用户名" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-input
              placeholder="请输入用户名"
              allow-clear
              v-decorator="[
              'name',
              { rules: [{ required: true, message: '请输入用户名' }] },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="昵称" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-input
              placeholder="请输入昵称"
              allow-clear
              v-decorator="[
              'nickname',
              { rules: [{ required: true, message: '请输入昵称' }] },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="头像" hidden :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-input
              v-decorator="[
              'avatar',
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="头像" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-upload
              name="avatar"
              list-type="picture-card"
              class="avatar-uploader"
              :show-upload-list="false"
              :before-upload="beforeUpload"
          >
            <img v-if="previewData || updateUserInfoForm.getFieldValue('avatar')" :src="previewData || updateUserInfoForm.getFieldValue('avatar')" alt="avatar" />
            <div v-else>
              <a-icon :type="loading ? 'loading' : 'plus'" />
              <div class="ant-upload-text">
                Upload
              </div>
            </div>
          </a-upload>
        </a-form-item>
        <a-form-item label="性别" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-select
              v-decorator="[
                'gender',
                { rules: [{ required: true, message: '请选择性别' }] },
              ]"
              placeholder="性别"
          >
            <a-select-option :value="1">
              男
            </a-select-option>
            <a-select-option :value="2">
              女
            </a-select-option>
            <a-select-option :value="0">
              -
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="生日" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-date-picker
              style="width: 100%"
              :valueFormat="'YYYY-MM-DD'"
              v-decorator="[
                'birthday',
                { rules: [{ required: true, message: '请选择生日' }] },
              ]"
              placeholder="生日"
          />
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal
        v-model="updatePasswordVisible"
        title="修改密码"
        centered
        @ok="updatePassword"
        :afterClose="updatePasswordAfterCloseHandle"
    >
      <a-form :form="updatePasswordForm" layout="horizontal">
        <a-form-item label="密码" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-input
              class="ant-input-sm"
              placeholder="请输入密码"
              type="password"
              allow-clear
              v-decorator="[
              'password',
              { rules: [{ required: true, message: '请输入密码' }] },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="确认密码" :label-col="{span: 4}" :wrapper-col="{span: 14}">
          <a-input
              class="ant-input-sm"
              placeholder="请输入密码"
              type="password"
              allow-clear
              v-decorator="[
              'confirm_password',
              { rules: [
                  { required: true, message: '请输入确认密码' },
                  {validator: compareToFirstPassword1,},
              ] },
            ]"
          ></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import Vue from 'vue'
import {Menu, Dropdown, Icon, Modal, Form, Input,Button,Upload,Select, message,Popover,DatePicker} from 'ant-design-vue';

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
    'a-select':Select,
    'a-select-option':Select.Option,
    'a-popover':Popover,
    'a-date-picker':DatePicker,
  },
  computed: {
    currentMenu(){
      return this.$store.state.menu.currentMenu
    },
    userInfo(){
      return this.$store.state.user.info
    },
  },
  methods: {
    setMenu(menu){
      this.$store.dispatch('SET_CURRENT_MENU',menu)
    },
    createRoomModalShow(){
       this.createRoomModaVisible = true
    },
    joinRoomModalShow(){
      this.joinRoomModalVisible = true
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
    async uploadAvatar(param,callback){
      let dirtype = param.dirtype
      let that = this
      //先尝试上传聊天室图片
      if(that.previewData !== ""){
        const contentType = getContentTypeFromBase64(that.previewData)
        if(contentType){
          const response = await that.$http.get(`/v1/oss/signature?dirtype=${dirtype}&content-type=${contentType}`);
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
                .then(callback(response))
                .catch(function (error) {
                  console.log(error);
                });
          }
        }
      }
    },
    createRoom(e){
      e.preventDefault();
      this.createRoomForm.validateFields(async(err) => {
        if (!err) {
          let that = this
          await that.uploadAvatar({dirtype:'room_avatar'},function (response) {
            that.createRoomForm.setFieldsValue({
              avatar: response.data.data.url,
            });
          })
          that.createRoomForm.validateFields(async(err, form) => {
            if (!err) {
              that.$http.post('/v1/rooms',form)
                .then(function(data){
                  if(data.data.errcode !== 0){
                    formErrorPrompt(that.createRoomForm,form,data.data.msg)
                  }else{
                    that.$store.dispatch('GET_ROOMS')
                    that.$store.dispatch('GET_CHATS')
                    message.success('创建成功');
                    that.createRoomModaVisible = false
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
              that.joinRoomModalVisible = false
            }
          })
          .catch(function (err){
            console.log(err)
          })
    },
    searchRoom(e){
      e.preventDefault();
      this.searchRoomForm.validateFields((err, form) => {
        if (!err) {
          let that = this
          that.$http.get(`/v1/rooms/find?field=${form.field}&content=${form.content}`)
              .then(function(data){
                if(data.data.errcode !== 0){
                  formErrorPrompt(that.searchRoomForm,form,data.data.msg)
                  that.searchRoomInfo = null
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
      this.createRoomForm.resetFields(["name","avatar"])
      this.previewData = ''
    },
    joinRoomAfterCloseHandle(){
      this.searchRoomForm.resetFields(["field","content"])
      this.searchRoomInfo = null
    },
    updateUserInfoAfterCloseHandle(){
      this.updateUserInfoForm.resetFields(["name","nickname","avatar","gender","birthday"])
      this.previewData = ''
    },
    updatePasswordAfterCloseHandle(){
      this.updatePasswordForm.resetFields(["password","confirm_password"])
    },
    userInfoModalShow(){
      this.getUserInfo()
      this.userSettingVisible = false
    },
    getUserInfo:function () {
      let that = this
      that.$http.get(`/v1/users/${this.userInfo.id}`)
          .then(function(data){
            if(data.data.errcode !== 0){
              message.error(data.data.msg);
            }else{
              that.$store.dispatch('SET_USER_INFO',data.data.data)
              that.userInfoVisible = true
            }
          })
    },
    updateUserInfoModalShow(){
      this.getUserInfo()
      this.updateUserInfoVisible = true
      this.Data = this.userInfo.avatar
      this.$nextTick(() => {
        this.updateUserInfoForm.setFieldsValue({
          name:this.userInfo.name,
          nickname:this.userInfo.nickname,
          avatar:this.userInfo.avatar,
          gender:this.userInfo.gender,
          birthday:this.userInfo.birthday,
        });
      });
    },
    updateUserInfo(e){
      e.preventDefault();
      this.updateUserInfoForm.validateFields(async(err) => {
        console.log('updateUserInfo')
        if (!err) {
          let that = this
          await that.uploadAvatar({dirtype:'user_avatar'},function (response) {
            that.updateUserInfoForm.setFieldsValue({
              avatar: response.data.data.url,
            });
          })
          that.updateUserInfoForm.validateFields(async(err,form) => {
            if (!err) {
              that.$http.put('/v1/users',form)
                  .then(function(data){
                    if(data.data.errcode !== 0){
                      formErrorPrompt(that.updateUserInfoForm,form,data.data.msg)
                    }else{
                      that.$store.dispatch('GET_USER')
                      message.success('修改成功');
                      that.updateUserInfoVisible = false
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
    updatePassword(e){
      e.preventDefault();
      this.updatePasswordForm.validateFields(async(err,form) => {
        if (!err) {
          let that = this
          that.$http.put('/v1/users/password',form)
              .then(function(data){
                if(data.data.errcode !== 0){
                  formErrorPrompt(that.updatePasswordForm,form,data.data.msg)
                }else{
                  message.success('修改成功');
                  that.updatePasswordVisible = false
                }
              })
              .catch(function (err){
                console.log(err)
              })
        }
      });
    },
    updatePasswordModalShow(){
      this.updatePasswordVisible = true
    },
    compareToFirstPassword1(rule, value, callback) {
      // console.log('compareToFirstPassword1')
      const form = this.updatePasswordForm;
      if (value && value !== form.getFieldValue('password')) {
        callback('两次密码输入不一致!');
      } else {
        callback();
      }
    },
    logout(){
      this.$store.dispatch('LOGOUT')
      this.$router.push('/login')
    }
  },
  data(){
    return{
      createRoomModaVisible:false,
      joinRoomModalVisible:false,
      userInfoVisible:false,
      userSettingVisible:false,
      updateUserInfoVisible:false,
      updatePasswordVisible:false,
      createRoomForm:this.$form.createForm(this, { name: 'coordinated' }),
      searchRoomForm:this.$form.createForm(this, { name: 'coordinated' }),
      updateUserInfoForm:this.$form.createForm(this, { name: 'coordinated' }),
      updatePasswordForm:this.$form.createForm(this, { name: 'coordinated' }),
      searchRoomInfo:null,
      loading: false,
      previewData: '',
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
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
      align-items: center;
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
        img{
          width: 100%;
          height: 100%;
        }
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
.card-content{
  display: flex;
  flex-direction: column;
  width: 332px;
  height: 484px;
  background: url(https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/static/card-background.png);
  background-repeat: no-repeat;
  padding: 90px 24px 0;
  background-size: 332px 115px;
  .card-title{
    display: flex;
    flex-direction: row;
    .avatar{
      width: 48px;
      height: 48px;
      img{
        width: 100%;
        height: 100%;
      }
    }
    .info{
      margin-left: 10px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      line-height: 24px;
      .nickname{
        font-size: 14px;
        color: #fff;
      }
      .id{
        font-size: 12px;
        color: #515e70;
      }
    }
  }
  .card-col{
    padding:10px 0;
    border-bottom: 1px solid #ccc;
    .col-title{
      font-size: 14px;
    }
    .card-row{
      height:28px;
      line-height: 28px;
      font-size: 12px;
      .label{
        display: inline-block;
        width: 72px;
      }
      .content{
        color: #131f41;
      }
    }

  }
  .footer{
    display: flex;
    flex-direction: row;
    justify-content: center;
    height: 60px;
    align-items: center;
  }
}

.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card img {
  width: 100%;
  height: 100%;
}

.userinfo{
  display: flex;
  flex-direction: row;
  align-items: center;
  margin:5px 0px;
  .avatar{
    width:40px;
    height:40px;
    margin-right: 10px;
    img{
      width:100%;
      height:100%;
    }
  }
}

.tools{
  cursor: pointer;
  .item{
    display: flex;
    height: 40px;
    align-items: center;
    justify-content: space-between;
    padding: 0 5px;
  }
  .item:hover {
    background-color: #f3f8fe;
  }
}
</style>
