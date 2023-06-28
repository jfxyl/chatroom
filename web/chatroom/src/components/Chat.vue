<template>
  <div class="container">
    <div id="chat-content" v-show="Object.keys(currentChat).length > 0">
      <div class="header">
        {{currentChat.name}}
      </div>
      <div id="chat-box" class="content"  @scroll="handleChatBoxScroll">
        <ul>
          <li v-for="(msg,index) in currentMsgs" :id="'message-'+msg.id" :key="index">
            <div :class="[msg.sender_id != userInfo.id?'other':'self','body']">
              <div class="avatar">
                <img :src="msg.sender.avatar" alt="">
              </div>
              <div class="msg">
                <div class="msg-header">
                  <div class="sender">{{msg.sender.nickname}}</div>
                  <div class="sendtime">{{formatDate(msg.created_at)}}</div>
                </div>
                <div class="msg-body">
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
                <a-popover title="消息接收人列表" trigger="click" placement="topRight">
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
                  <div class="msg-footer" v-if="msg.sender_id == userInfo.id && msg.unreaders.length > 0" @click="getReadInfo(msg)">
                    {{msg.unreaders.length}}人未读
                  </div>
                </a-popover>
              </div>
            </div>
          </li>
        </ul>
      </div>
      <div class="input">
        <div class="tools">
          <ul class="toolsul">
            <li  class="toolsli" @click.stop="emojiIsShow = !emojiIsShow">
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
          <EmojiPanel v-if="emojiIsShow" @chooseEmoji="chooseEmoji" class="emojis"></EmojiPanel>
        </div>
        <textarea class="textarea" ref="textarea" name="" id="" cols="30" rows="10" v-model="msg" @keydown.enter.prevent.exact="sendMsg" @keydown.shift.enter="addLine"></textarea>
        <!--<div class="textarea" contenteditable v-html="msg" @blur="bindMsg"  ></div>-->
      </div>
    </div>
    <div v-show="Object.keys(currentChat).length == 0">
    </div>
<!--    <ReadInfoPanel v-if="readInfoVisible"></ReadInfoPanel>-->
  </div>
</template>

<script>
import EmojiPanel from './Emoji.vue'
// import ReadInfoPanel from './ReadInfo.vue'
import {message,Popover} from "ant-design-vue";
import moment from 'moment';

export default {
  name: "ChatPanel",
  components: {
    EmojiPanel,
    // 'a-modal':Modal,
    'a-popover':Popover,
  },
  data(){
    return {
      msg:"",
      emojiIsShow:false,
      readInfoVisible: false,
      readInfo:{
        readers:[],
        unreaders:[],
      },
      readInfoStyle:{
        position: 'absolute',
        top: '0px',
        right: '0px',
      }
    }
  },
  props:['chat'],
  created() {
    console.log("created")
  },
  mounted:function(){
    console.log("mounted")
    this.scrollToBottom()
    this.globalClick(this.handleClickOut);
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
    handlePopoverVisibleChange(){

    },
    formatDate(datetime) {
      return moment(datetime).format('YYYY年M月D日 HH:mm');
    },
    // 监听消息框的滚动事件
    handleChatBoxScroll:function() {
      const chatBox = document.getElementById('chat-box');
      const visibleMessages = this.getVisibleMessages(chatBox);
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
      for (const message of this.currentMsgs) {
        const messageElement = document.getElementById(`message-${message.id}`);
        if (messageElement.offsetTop >= scrollTop && messageElement.offsetTop <= scrollTop + chatBox.clientHeight) {
          visibleMessages.push(message);
        }
      }
      return visibleMessages;
    },
    markMessageAsRead:function(message){
      if (message.readed === false) {
        var content = {
          action:'read',
          body:{
            message_id:message.id
          }
        }
        this.$store.state.socket.socket.send(JSON.stringify(content))
        message.readed = true
      }
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
      this.emojiIsShow = false
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
      var content = document.querySelector('.content');
      setTimeout(() => {
        content.scrollTop = content.scrollHeight
      })
    },
    chooseFile(){
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
    currentMsgs(){
      this.scrollToBottom()
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
      height:60px;
      padding:25px 0 0 25px;
      box-sizing: border-box;
      font-size:20px;
      border-bottom: 1px solid #ccc;
    }
    .content{
      flex:1;
      overflow-y: scroll;
      ul{
        list-style: none;
        li{
          min-height: 40px;
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
                clear: both;
                float: right;
                color: #006aff;
                font-size: 12px;
                cursor: pointer;
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
  .readers{
  }
  .unreaders{
  }
}
</style>
