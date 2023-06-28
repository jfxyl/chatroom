<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import {message} from "ant-design-vue";

export default {
  data(){
    return{
    }
  },
  created() {
    var that = this
    that.$http.get('/v1/users')
      .then(function(data){
        if(data.data.errcode !== 0){
          message.error(data.data.msg);
          if(that.$route.path !== '/login' && that.$route.path !== '/Login'){
            that.$router.push('/login');
          }
        }else{
          that.$store.dispatch('LOGIN',data.data.data)
          that.$store.dispatch('GET_CHATS')
          that.$store.dispatch('GET_ROOMS')
          if(that.$route.path === '/login' || that.$route.path === '/Login'){
            that.$router.push('/');
          }
        }
      })
      .catch(function(err){
        console.log(err)
      })
  },
  methods:{

  },
  computed:{
    isLogin(){
      return this.$store.state.user.auth
    }
  },
  watch:{
    isLogin(){
      console.log(this.isLogin)
      if(this.isLogin){
        this.$store.dispatch('CONN_SOCKET')
      }else{
        this.$store.dispatch('CLOSE_SOCKET')
      }
    }
  }
}
</script>

<style>
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
body{
  margin: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
