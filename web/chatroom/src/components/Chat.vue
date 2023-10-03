<template>
  <div class="container">
    <div id="chat-content" v-show="currentChat.id">
      <div class="header">
        <div class="left">
          <div class="avatar">
            <img :src="currentChat.avatar" alt="">
          </div>
          <div class="info">
            <div class="name">{{currentChat.name}}</div>
            <div class="count">
              <img src="https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/static/users.png" alt="">
              <span>{{currentChat.users ? currentChat.users.length : 0}}</span>
            </div>
          </div>
        </div>
        <div class="right">
<!--          <div class="item" @click="adduser">-->
<!--            <a-icon type="usergroup-add" />-->
<!--          </div>-->
          <div class="item" @click="setting">
            <a-icon type="setting" />
          </div>
        </div>
      </div>
      <div id="chat-box" class="content"  @scroll="handleChatBoxRead" ref="chatmsgs">
        <ul v-infinite-scroll="loadMore">
          <li v-for="(msg,index) in currentMsgs" :id="'message-'+msg.id" :key="index" ref="msg">
            <div v-if="msg.msg_type === _const.TypeNotice" class="notice">{{msg.content}}</div>
            <div v-else :class="[msg.sender_id != userInfo.id?'other':'self','body']">
              <div class="avatar" @click="getUserInfo(msg.sender)">
                <img :src="msg.sender.avatar" alt="">
              </div>
              <div class="msg">
                <div class="msg-header">
                  <div class="sender">{{msg.sender.nickname}}</div>
                  <div class="sendtime">{{formatDate(msg.created_at)}}</div>
                </div>
                <div class="msg-body">
                  <div class="msg-content">
                    <div v-if="msg.msg_type === _const.TypeText" class="text">
                      <span><pre>{{msg.content}}</pre></span>
                      <div class="arrow"></div>
                    </div>
                    <div v-else-if="msg.msg_type === _const.TypeImg" class="img">
                      <img :src="msg.content" alt="">
                    </div>
                    <div v-else-if="msg.msg_type === _const.TypeFile" class="file">
                      <div class="file-type">
                        file
                      </div>
                      <div class="file-status">
                        <div class="file-name">{{msg.content}}</div>
                        <div class="file-size">250kb</div>
                      </div>
                    </div>
                  </div>
                </div>
                <a-popover title="消息接收人列表" trigger="click" placement="topRight">
                  <div class="msg-footer" v-if="msg.sender_id == userInfo.id" @click="getReadInfo(msg)">
                    <div class="unread" v-if="msg.unreaders.length > 0">
                      {{msg.unreaders.length}}人未读
                    </div>
                    <div class="allread" v-else >
                      全部已读
                    </div>
                  </div>
                  <template #content>
                    <div class="readinfo">
                      <div class="list readers">
                        <p class="notice"><span class="blue">{{readInfo.readers.length}}</span> 人已读</p>
                        <ul>
                          <li v-for="(reader,index) in readInfo.readers" :key="index">
                            <div class="avatar">
                              <img :src="reader.avatar" alt="">
                            </div>
                            <div class="name">{{reader.nickname}}</div>
                          </li>
                        </ul>
                      </div>
                      <div class="list unreaders">
                        <p class="notice"><span class="blue">{{readInfo.unreaders.length}}</span> 人未读</p>
                        <ul>
                          <li v-for="(reader,index) in readInfo.unreaders" :key="index">
                            <div class="avatar">
                              <img :src="reader.avatar" alt="">
                            </div>
                            <div class="name">{{reader.nickname}}</div>
                          </li>
                        </ul>
                      </div>
                    </div>
                  </template>
                </a-popover>
              </div>
            </div>
          </li>
        </ul>
      </div>
      <div class="input">
        <div class="tools">
          <ul class="toolsul">
            <li  class="toolsli" @click.stop="emojiVisible = !emojiVisible">
              <svg class="toolssvg emoji" aria-hidden="true">
                <use xlink:href="#icon-emoji"></use>
              </svg>
            </li>
            <li class="toolsli" @click="chooseFile">
              <svg class="toolssvg emoji" aria-hidden="true">
                <use xlink:href="#icon-file"></use>
              </svg>
            </li>
          </ul>
          <EmojiPanel v-if="emojiVisible" @chooseEmoji="chooseEmoji" class="emojis"></EmojiPanel>
        </div>
        <textarea class="textarea" ref="textarea" name="" id="" cols="30" rows="10" v-model="msg" @keydown.enter.prevent.exact="sendMsg" @keydown.shift.enter="addLine"></textarea>
        <!--<div class="textarea" contenteditable v-html="msg" @blur="bindMsg"  ></div>-->
      </div>
    </div>
    <div v-show="Object.keys(currentChat).length == 0">
    </div>
    <a-drawer
        title="设置"
        placement="right"
        :closable="true"
        :visible="settingVisible"
        width="420"
        :maskClosable="true"
        :mask="true"
        :maskStyle="{animation: 'none',opacity:0}"
        :bodyStyle="{padding:0}"
        @close="onSettingClose"
        :zIndex=1000
        :wrapStyle="{marginTop: '42px'}"
    >
      <div class="room-drawer">
        <div class="item header">
          <div class="avatar">
            <img :src="currentChat.avatar" alt="">
          </div>
          <div class="info">
            <div class="name">{{currentChat.name}}&nbsp;<a-icon type="copy" @click="copy(currentChat.name)"/></div>
          </div>
        </div>
        <div class="item">
          <div class="">
            成员 {{currentChat.users ? currentChat.users.length : 0}}人
          </div>
          <div v-if="currentChat && currentChat.users" class="members">
            <div class="member" v-for="(member,index) in currentChat.users.slice(0, 20)" :key="index" @click="getUserInfo(member)">
              <div class="avatar">
                <img :src="member.avatar" alt="">
              </div>
              <div class="nickname">
                {{member.nickname}}
              </div>
            </div>
          </div>
          <div v-if="currentChat && currentChat.users && currentChat.users.length > 20" class="more" @click="moreMembers">查看更多</div>
        </div>
        <div class="item">
          <div class="item-row">
            <span class="label">聊天室ID</span>
            <span class="content">{{currentChat.id}}&nbsp;<a-icon type="copy" @click="copy(currentChat.id)"/></span>
          </div>
        </div>
        <div class="item footer">
          <a-button type="danger" ghost @click="quitRoom(currentChat)">
            {{currentChat.owner == userInfo.id ? '解散' : '退出'}}聊天室
          </a-button>
        </div>
      </div>
    </a-drawer>
    <a-drawer
        title="成员"
        width="420"
        :closable="true"
        :visible="moreMembersVisible"
        :maskClosable="true"
        :mask="true"
        :maskStyle="{animation: 'none',opacity:0}"
        :bodyStyle="{padding:0,height: 'calc(100% - 97px)',overflow: 'auto'}"
        @close="onMoreMembersClose"
        :zIndex=1001
        :wrapStyle="{marginTop: '42px'}"
    >
<!--      <div slot="title" class="member-title">-->
<!--        <div>成员</div>-->
<!--        <div class="tools">-->
<!--          <div><a-icon type="search" /></div>-->
<!--        </div>-->
<!--      </div>-->
      <div class="member-drawer">
        <div v-if="currentChat && currentChat.users" class="members">
          <div class="member" v-for="(member,index) in currentChat.users" :key="index" @click="getUserInfo(member)">
            <div class="avatar">
              <img :src="member.avatar" alt="">
            </div>
            <div class="name">
              {{member.nickname}}
            </div>
          </div>
        </div>
      </div>
    </a-drawer>
    <a-modal
        v-model="userInfoVisible"
        title=""
        centered
        width="400"
        :footer="null"
        :mask=false
        :closable=false
        :body-style="{padding:0}"
        :zIndex=1002
    >
      <div class="card-content">
        <div class="card-title">
          <div class="avatar">
            <img :src="displayUserInfo.avatar" alt="">
          </div>
          <div class="info">
            <div class="nickname">{{displayUserInfo.nickname}}</div>
            <div class="id">{{displayUserInfo.id}}</div>
          </div>
        </div>
        <div class="card-col">
          <div class="col-title">个人信息</div>
          <div class="card-row">
            <span class="label">用户名</span>
            <span class="content">{{displayUserInfo.name}}</span>
          </div>
          <div class="card-row">
            <span class="label">昵称</span>
            <span class="content">{{displayUserInfo.nickname}}</span>
          </div>
          <div class="card-row">
            <span class="label">性别</span>
            <span class="content">{{_genderMap[displayUserInfo.gender]}}</span>
          </div>
          <div class="card-row">
            <span class="label">生日</span>
            <span class="content">{{displayUserInfo.birthday}}</span>
          </div>
          <div class="card-row">
            <span class="label">注册时间</span>
            <span class="content">{{displayUserInfo.created_at}}</span>
          </div>
        </div>
      </div>
    </a-modal>
<!--    <a-modal v-model="userInfoVisible">-->
<!--      <template v-slot:footer></template>-->
<!--      <p>这是一个没有关闭按钮的模态框。</p>-->
<!--    </a-modal>-->
  </div>
</template>

<script>
import EmojiPanel from './Emoji.vue'
// import ReadInfoPanel from './ReadInfo.vue'
import {message,Popover,Modal,infiniteScroll,Icon,Drawer,Button } from "ant-design-vue";
import {copy} from "@/utils/utils";
import moment from 'moment';

export default {
  name: "ChatPanel",
  directives: { infiniteScroll },
  components: {
    EmojiPanel,
    'a-modal':Modal,
    'a-popover':Popover,
    'a-icon':Icon,
    'a-drawer':Drawer,
    'a-button':Button,
  },
  data(){
    return {
      msg:"",
      emojiVisible:false,
      readInfoVisible: false,
      userInfoVisible: false,
      settingVisible: false,
      moreMembersVisible: false,
      readInfo:{
        readers:[],
        unreaders:[],
      },
      readInfoStyle:{
        position: 'absolute',
        top: '0px',
        right: '0px',
      },
      displayUserInfo:{},
    }
  },
  props:['chat'],
  created() {
    console.log("created")
  },
  mounted:function(){
    console.log("mounted")
    // this.scrollToBottom()
    this.globalClick(this.handleClickOut);
    this.$refs.chatmsgs.addEventListener('scroll', () => {
      console.log('this.$refs.chatmsgs.scrollTop',this.$refs.chatmsgs.scrollTop)
      if (this.$refs.chatmsgs.scrollTop === 0) {
          this.loadMore();
      }
    })
  },
  computed:{
    currentChat(){
      return this.$store.state.chat.currentChat
    },
    currentMsgs(){
      return this.$store.state.chat.currentMsgs
    },
    userInfo(){
      return this.$store.state.user.info
    },
  },
  methods:{
    async loadMore() {
      var hmsg, hid
      if (Array.isArray(this.$refs.msg)) {
        if (this.$refs.msg.length > 0) {
          hmsg = this.$refs.msg[0]
          hid = hmsg.id
        }
      } else {
        hmsg = this.$refs.msg
        hid = hmsg.id
      }
      var that = this
      await this.$store.dispatch('GET_MORE_MESSAGE').then(() => {
        if (hid && that.$refs.msg) {
          this.$nextTick(() => {
            hmsg = that.$refs.msg.find(item => item.id === hid);
            that.$refs.chatmsgs.scrollTop = hmsg.offsetTop - 102;
          })
        }
      });
    },
    copy(content){
      copy(content)
    },
    adduser(){

    },
    setting(){
      this.settingVisible = true
    },
    moreMembers(){
      this.moreMembersVisible = true
    },
    onSettingClose(){
      this.settingVisible = false
      console.log("onSettingClose")
    },
    onMoreMembersClose(){
      this.moreMembersVisible = false
      console.log("onSettingClose")
    },
    quitRoom(room) {
      let that = this
      Modal.confirm({
        title: `确定${this.currentChat.owner == this.userInfo.id ? '解散' : '退出'}聊天室么?`,
        okText: '确定',
        cancelText: '取消',
        onOk() {
          that.$http.post(`/v1/rooms/${room.id}/quit`)
              .then(function(data){
                if(data.data.errcode !== 0){
                  message.error(data.data.msg);
                }else{
                  that.settingVisible = false
                  that.$store.dispatch('SET_CURRENT_CHAT', {})
                  that.$store.dispatch('GET_ROOMS')
                  that.$store.dispatch('GET_CHATS')
                  message.success('操作成功');
                }
              })
              .catch(function (err){
                console.log(err)
              })
        },
        onCancel() {},
      });
    },
    formatDate(datetime) {
      return moment(datetime).format('YYYY年M月D日 HH:mm');
    },
    // 监听消息框的滚动事件
    handleChatBoxRead:function() {
      const chatBox = document.getElementById('chat-box');
      const visibleMessages = this.getVisibleMessages(chatBox);
      console.log('visibleMessages',visibleMessages)
      for (const message of visibleMessages) {
        if (message.readed === false) {
          this.markMessageAsRead(message);
          message.unread = false; // 清除未读标记
        }
      }
    },
    // 获取消息框内可视区域内的消息列表
    getVisibleMessages:function(chatBox) {
      const scrollTop = chatBox.scrollTop + 102;
      const visibleMessages = [];
      console.log('this.currentMsgs',this.currentMsgs)
      for (var message of this.currentMsgs) {
        var messageElement = document.getElementById(`message-${message.id}`);
        if(!messageElement) continue
        var middleOffsetTop = messageElement.offsetTop + messageElement.clientHeight / 2;
        console.log('messageElement',messageElement)
        console.log('middleOffsetTop',middleOffsetTop)
        if (middleOffsetTop >= scrollTop && middleOffsetTop <= scrollTop + chatBox.clientHeight) {
          visibleMessages.push(message);
        }
      }
      return visibleMessages;
    },
    markMessageAsRead:function(message){
      if (message.msg_type != this._const.TypeNotice &&  message.readed === false) {
        var content = {
          action:'read',
          body:{
            message_id:message.id
          }
        }
        this.$store.state.socket.socket.send(JSON.stringify(content))
        this.currentChat.unread_count = this.currentChat.unread_count > 0 ? this.currentChat.unread_count - 1 : 0;
        message.readed = true
      }
    },
    getUserInfo:function (user) {
      let that = this
      that.$http.get(`/v1/users/${user.id}`)
          .then(function(data){
            if(data.data.errcode !== 0){
              message.error(data.data.msg);
            }else{
              console.log(data.data.data)
              that.displayUserInfo = data.data.data
              that.userInfoVisible = true
            }
          })
    },
    getReadInfo:function(message){
      let that = this
      that.$http.get(`/v1/messages/${message.id}/readinfo`)
          .then(function(data){
            if(data.data.errcode !== 0){
              message.error(data.data.msg);
            }else{
              console.log(data.data.data)
              that.readInfo = data.data.data
              that.readInfoVisible = true
              // console.log('that.$refs.messages',that.$refs.messages)
              // const msgElement = that.$refs.messages[index];
              // const rect = msgElement.getBoundingClientRect();
              // that.readInfoStyle = {
              //   background:'#000000',
              //   position: 'absolute',
              //   top: rect.top + 'px',
              //   right: rect.right + 'px',
              // }
            }
          })
    },
    handleClickOut:function () {
      this.emojiVisible = false
      // this.settingVisible = false
    },
    sendMsg(){
      if(this.msg.length > 0){
        this.requestMsg(this._const.TypeText,this.msg)
      }
    },
    addLine(){
      this.insertText("\n")
    },
    chooseEmoji(emoji){
      this.insertText(emoji)
    },
    insertText(value){
      var textDom = this.$refs.textarea
      var startPos = textDom.selectionStart
      var endPos = textDom.selectionEnd
      this.msg = this.msg.substring(0,textDom.selectionStart) + value + this.msg.substring(textDom.selectionEnd,this.msg.length)
      startPos = startPos + value.length;
      endPos = endPos + value.length;
      textDom.blur()
      setTimeout(() => {
        textDom.setSelectionRange(startPos, endPos);
        textDom.focus();
      })
    },
    scrollToBottom(){
      console.log('scrollToBottom')
      var content = document.querySelector('#chat-box');
      console.log('content',content)
      setTimeout(() => {
        console.log(content.scrollHeight,content.scrollTop)
        content.scrollTop = content.scrollHeight
      })
    },
    chooseFile(){
      message.warning('暂不支持');
      // let _this = this
      // remote.dialog.showOpenDialog({
      //   properties: ['openFile']
      // },function(dir){
      //   if(dir !== void 0){
      //     _this._function.upload(dir[0])
      //         .then(function(result){
      //           if(result){
      //             _this.requestMsg(result.type,result.url)
      //           }else{
      //             console.log('文件上传失败')
      //           }
      //         })
      //   }
      // })
    },
    requestMsg(msg_type,content){
      var msg = {
        chat_type:2,  //群聊
        receiver_id:this.currentChat.id,
        msg_type:msg_type,
        content:content,
        reply_id:0,
      }
      console.log(msg,JSON.stringify(msg))
      let that = this
      that.$http.post('/v1/messages',msg)
        .then(function(data){
          if(data.data.errcode !== 0){
            message.error(data.data.msg);
          }else{
            that.currentMsgs.push(data.data.data)
            that.msg = ''
            that.scrollToBottom()
          }
        })
      // this.$store.state.socket.socket.send(JSON.stringify(message))
    }
  },
  watch:{
    currentChat(){
      this.$nextTick(() => {
        console.log('scrollToBottom')
        const chatBox = document.getElementById('chat-box');
        if(chatBox.scrollHeight <= chatBox.clientHeight){
          this.handleChatBoxRead()
        }else{
          this.scrollToBottom()
        }
      })
    },
    currentMsgs(msgs){
      this.$nextTick(() => {
        if (msgs.length === 0) return
        const chatBox = document.getElementById('chat-box');
        if (chatBox.scrollHeight <= chatBox.clientHeight) {
          this.handleChatBoxRead()
        } else {
          const scrollTop = chatBox.scrollTop + 102;
          //获取倒数第二条消息
          const index = msgs.length >= 2 ? msgs.length - 2 : 0;
          const messageElement = document.getElementById(`message-${msgs[index].id}`);
          console.log('messageElement')
          if (messageElement) {
            const middleOffsetTop = messageElement.offsetTop + messageElement.clientHeight / 2;
            console.log('middleOffsetTop', middleOffsetTop, 'scrollTop', scrollTop, 'chatBox.clientHeight', chatBox.clientHeight)
            if (middleOffsetTop >= scrollTop && middleOffsetTop <= scrollTop + chatBox.clientHeight) {
              this.scrollToBottom()
            }
          }
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.container{
  height: 100%;
  display: flex;
  flex-direction: column;
  flex:1;
  #chat-content{
    height: 100%;
    display: flex;
    flex-direction: column;
    flex:1;
    .header{
      height:55px;
      padding:0 20px;
      box-sizing: border-box;
      font-size:20px;
      border-bottom: 1px solid #ccc;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .left{
        display: flex;
        flex-direction: row;
        align-items: center;
        .avatar{
          width:44px;
          height: 44px;
          margin-right: 10px;
          img{
            width: 100%;
            height: 100%;
          }
        }
        .info{
          display: flex;
          flex-direction: column;
          line-height: 24px;
          align-items: baseline;
          .name{
            font-size: 16px;
          }
          .count{
            font-size: 12px;
            color: #999;
            img{
              height: 12px;
              vertical-align: baseline;
            }
          }
        }
      }
      .right{
        display: flex;
        .item{
          width: 24px;
          height: 24px;
          margin-left: 10px;
          i{
            width:100%;
            height: 100%;
          }
        }
      }
    }
    .content{
      flex:1;
      overflow-y: scroll;
      ul{
        list-style: none;
        li{
          min-height: 40px;
          .notice{
            text-align: center;
            color: #999;
            font-size: 12px;
            margin: 10px 0;
          }
          .body{
            overflow: auto;
            padding:10px;
            .avatar{
              background: #fff;
              margin:0 10px;
              width:40px;
              height:40px;
              display: inline-block;
              cursor: pointer;
              img{
                width:100%;
                height:100%;
                border-radius: 3px;
              }
            }
            .msg{
              .msg-header{
                color:#999;
                display: flex;
                margin-bottom: 3px;
                .sendtime{
                  font-size: 12px;
                }
                .sender{
                  font-size: 13px;
                }
              }
              .msg-body{
                display: inline-block;
                max-width: 340px;
                line-height: 20px;
                position: relative;
                display: flex;
                flex-direction: row;
                .msg-content{
                  span{
                    display: block;
                    padding:10px;
                    pre{
                      white-space: pre-wrap;
                      word-wrap: break-word;
                      margin:0;
                    }
                  }
                  img{
                    width:100%;
                    height:100%;
                    display: block;
                    border-radius: 3px;
                  }
                }
              }
            }
          }
          .other{
            .avatar{
              float:left;
            }
            .msg{
              float:left;
              .msg-header{
                .sender{
                  margin-right: 10px;
                }
              }
              .msg-body{
                float: left;
              }
            }
            .text{
              background: #f0f6ff;
              border-radius: 3px;
            }
            .file{
              display:flex;
              width:200px;
              height:60px;
              background-color:#fff;
              .file-type{
                flex: 0 0 60px;
              }
              .file-status{
                flex:1;
                display: flex;
                flex-direction: column;
                .file-name{
                  font-size:14px;
                  flex: 1;
                  width: 140px;
                  word-wrap: break-word;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  display: -webkit-box;
                  -webkit-box-orient: vertical;
                  -webkit-line-clamp: 2;
                  overflow: hidden;
                }
                .file-size{
                  height:30px;
                }
              }
            }
            .arrow {
              position:absolute;
              top: 5px;
              left:-15px; /* 圆角的位置需要细心调试哦 */
              width:0;
              height:0;
              font-size:0;
              border:solid 8px;
              border-color:transparent #f0f6ff transparent transparent;
            }
          }
          .self{
            .avatar{
              float:right;
            }
            .msg{
              float:right;
              .msg-header{
                flex-direction: row-reverse;
                .sender{
                  margin-left: 10px;
                }
              }
              .msg-body{
                float: right;
              }
              .msg-footer{
                margin-top: 2px;
                clear: both;
                float: right;
                font-size: 12px;
                cursor: pointer;
                .unread{
                  color: #006aff;
                }
                .allread{
                  color: #999;
                }
              }
            }
            .text{
              background: #9EEA6A;
              border-radius: 3px;
            }
            .file{
              display:flex;
              width:200px;
              height:60px;
              background-color:#fff;
              .file-type{
                flex: 0 0 60px;
                text-align: center;
                line-height: 60px;
              }
              .file-status{
                flex:1;
                display: flex;
                flex-direction: column;
                .file-name{
                  font-size:14px;
                  flex: 1;
                  width: 140px;
                  word-wrap: break-word;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  display: -webkit-box;
                  -webkit-box-orient: vertical;
                  -webkit-line-clamp: 2;
                  overflow: hidden;
                }
                .file-size{
                  height:20px;
                }
              }
            }
            .arrow {
              position:absolute;
              top: 5px;
              right:-15px; /* 圆角的位置需要细心调试哦 */
              width:0;
              height:0;
              font-size:0;
              border:solid 8px;
              border-color:transparent transparent transparent #9EEA6A;
            }
          }

        }
      }
    }
    .input{
      border-top:1px solid #ccc;
      min-height:100px;
      .tools{
        position: relative;
        height:38px;
        .toolsul{
          list-style: none;
          height:24px;
          padding:7px 10px;
          .toolsli{
            display: block;
            float:left;
            margin:0 5px;
            .toolssvg{
              display: block;
              width: 26px;
              height: 26px;
              margin: 0 auto;
              fill: currentColor;
              overflow: hidden;
            }
          }
        }
        .emojis{
          position: absolute;
          bottom:38px;
          left:0px;
        }
      }
      .textarea{
        outline: none;
        width:100%;
        height:150px;
        border:0;
        background:#F5F5F5;
        padding:10px;
        box-sizing: border-box;
        font-size:18px;
      }
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
}
.readinfo{
  display: flex;
  flex-direction: row;
  .list{
    width:180px;
    height:300px;
    overflow:auto;
    overflow-x:hidden;
    .notice{
      padding: 0 10px;
      margin: 0;
      .blue{
        color: #006aff;
      }
    }
    ul{
      list-style: none;
      li {
        min-height: 40px;
        display: flex;
        flex-direction: row;
        padding: 5px 10px;
        .avatar{
          width: 30px;
          height: 30px;
          line-height: 30px;
          img{
            width: 100%;
            height: 100%;
          }
        }
        .name{
          margin-left: 10px;
          line-height: 30px;
        }
      }
    }
  }
}
.room-drawer{
  width: 100%;
  height: 100%;
  .item{
    border-bottom: 8px solid #f7f8f9;
    padding:12px 24px;
    align-items: center;
    .members{
      display: flex;
      flex-wrap: wrap;
      .member{
        display: flex;
        flex-direction: column;
        width: 34.5px;
        margin-right: 2px;
        margin-bottom: 6px;
        .avatar{
          width:34.5px;
          height:34.5px;
          img{
            width:100%;
            height:100%;
          }
        }
        .nickname{
          word-break: break-all;
          word-wrap: break-word;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          width: 100%;
          text-align: center;
          font-size: 12px;
          margin-top: 4px;
        }
      }
    }
    .more{
      text-align: center;
      color:#006aff;
      cursor: pointer;
    }
    .item-row{
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      .label{
        font-size: 13px;
        font-weight: 500;
        color: #131f41;
      }
    }
  }
  .header{
    display: flex;
    flex-direction: row;
    align-items: center;
    .avatar{
      width:40px;
      height: 40px;
      margin-right: 10px;
      img{
        width: 100%;
        height: 100%;
      }
    }
    .info{
      display: flex;
      flex-direction: column;
      line-height: 24px;
      align-items: baseline;
      .name{
        font-size: 16px;
      }
    }
  }
  .footer{
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
}
//.member-title{
//  display: flex;
//  justify-content: space-between;
//  padding: 0 35px 0 0;
//  align-items: center;
//}
.member-drawer{
  .members{
    overflow:auto;
    overflow-x:hidden;
    .member{
      display: flex;
      flex-direction: row;
      align-items: center;
      padding:12px;
      .avatar{
        width:40px;
        height:40px;
        margin-right: 5px;
        img{
          width:100%;
          height:100%;
        }
      }
      .name{
        /*align-items: baseline;*/
        /*word-break: break-all;*/
      //word-wrap: break-word;
      //overflow: hidden;
      //text-overflow: ellipsis;
      //white-space: nowrap;
      //width: 100%;
      //font-size: 12px;
      //margin-top: 4px;
      }
    }
  }
}
</style>
