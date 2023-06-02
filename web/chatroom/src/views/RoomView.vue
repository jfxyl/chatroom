<template>
  <div class="room-container">
    <RoomsPanel></RoomsPanel>
    <RoomPanel></RoomPanel>
  </div>
</template>

<script>
import RoomsPanel from '@/components/Rooms.vue'
import RoomPanel from '@/components/Room.vue'
// import { Icon  } from 'ant-design-vue';
export default {
  name: 'RoomView',
  components: {
    RoomsPanel,
    RoomPanel
  },
  computed: {},
  mounted:function() {
    this.$store.dispatch('SET_CURRENT_MENU', 'room')
  },
  methods: {},
  data() {
    return {}
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.room-container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  height: 100%;
  flex:1;
}
</style>
