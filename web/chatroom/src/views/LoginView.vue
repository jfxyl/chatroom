<template>
  <div class="container">
    <div class="top"></div>
    <div class="center">
      <div class="center-container">
        <div class="left-container">
          <div class="title">在线化办公</div>
          <span class="sub_title">多人协作，打造高效办公方式</span>
          <img class="intro-image" src="https://web.rentsoft.cn/static/media/login_bg.e42640a5.png">
        </div>
        <LoginPanel v-if="action == 'login'" @jump="action = $event"></LoginPanel>
        <RegisterPanel v-else-if="action == 'register'" @jump="action = $event"></RegisterPanel>
      </div>
    </div>
    <div class="bottom"></div>
  </div>
</template>

<script>
import LoginPanel from '@/components/Login.vue'
import RegisterPanel from '@/components/Register.vue'

export default {
  name: 'LoginView',
  components:{
    LoginPanel,
    RegisterPanel,
  },
  computed: {

  },
  methods: {

  },
  data(){
    return{
      action:"login",
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
  height: 100vh;
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
    justify-content: center;
    align-items: center;
    .center-container{
      display: flex;
      min-height: 420px;
      .left-container{
        text-align: center;
        width: 354px;
        margin-right: 72px;
        .title{
          font-size: 22px;
          font-weight: 500;
          padding-top: 43px;
        }
        .sub_title{
          font-size: 12px;
          color: #777;
        }
        .intro-image{
          width: 100%;
        }
      }
    }
  }
  .bottom{
    height: 30px;
    width: 100%;
    background-color: #438be5;
  }
}
</style>
