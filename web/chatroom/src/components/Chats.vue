<template>
  <div class="chats">
    <div class="chat" v-for="(chat,name,index) in chats" :key="index">
      <div class="avatar">
        <img class="img" :src="chat.avatar" alt="">
      </div>
      <div class="info" :data-time="chat.lasttime">
        <div class="top">
          <div class="name">{{chat.name}}</div>
          <div class="time">{{chat.lasttime}}</div>
        </div>
        <div class="bottom">{{chat.lastmag.context}}</div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ChatsPanel',
  components:{

  },
  data(){
    return{
      chats:[
          {name:"aaa",avatar:"",lastmag:{type:"text",context:"haha11111111111hahahhaha11111111111hahahhaha11111111111hahahhaha11111111111hahah"},lasttime:"2023-01-01",unread:99},
          {name:"aaa",avatar:"https://web.rentsoft.cn/static/media/login_bg.e42640a5.png",lastmag:{type:"text",context:"hhhhhhh"},lasttime:"2023-01-01",unread:88}
      ]
    }
  },
  methods:{

  }
}
</script>

<style scoped>
.chats{
  width:350px;
  padding: 10px;
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
