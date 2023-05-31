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
          if(that.$route.path !== '/login' || that.$route.path !== '/Login'){
            that.$router.push('/login');
          }
        }else{
          that.$store.dispatch('LOGIN',data.data.data)
          that.$store.dispatch('GET_ROOMS')
          if(that.$route.path === '/login' || that.$route.path === '/Login'){
            that.$router.push('/');
          }
        }
      })
  },
  methods:{

  },
  computed:{

  },
  watch:{

  }
}
</script>

<style>
*{
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
