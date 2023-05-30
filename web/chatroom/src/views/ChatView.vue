<template>
  <div class="container">
    <ChatsPanel></ChatsPanel>
  </div>
</template>

<script>
import ChatsPanel from '@/components/Chats.vue'
// import { Icon  } from 'ant-design-vue';
export default {
  name: 'ChatView',
  components:{
    ChatsPanel
  },
  computed: {

  },
  methods: {

  },
  data(){
    return{
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

}
</style>
