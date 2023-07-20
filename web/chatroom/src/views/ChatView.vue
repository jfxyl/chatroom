<template>
  <div class="container">
    <ChatsPanel></ChatsPanel>
    <ChatPanel></ChatPanel>
  </div>
</template>

<script>
import ChatsPanel from '@/components/Chats.vue'
import ChatPanel from '@/components/Chat.vue'
// import { Icon  } from 'ant-design-vue';
export default {
  name: 'ChatView',
  components:{
    ChatsPanel,
    ChatPanel,
  },
  computed: {

  },
  mounted:function() {
    this.$store.dispatch('SET_CURRENT_MENU', 'chat')
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
<style scoped lang="scss">
.container{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  height: 100%;
  flex:1;
}
</style>
