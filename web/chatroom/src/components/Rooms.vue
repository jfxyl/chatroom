<template>
  <div class="rooms">
    <div class="room" v-for="(room,name,index) in rooms" :key="index">
      <div class="avatar">
        <img class="img" :src="room.avatar" alt="">
      </div>
      <div class="info" >
        <div class="top">
          <div class="name">{{room.name}}</div>
        </div>
        <div class="bottom">1000人</div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'RoomsPanel',
  components:{

  },
  data(){
    return{

    }
  },
  computed: {
    rooms() {
      return this.$store.state.room.rooms
    },
  },
  methods:{

  }
}
</script>

<style scoped>
.rooms{
  width:350px;
  padding: 10px;
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
