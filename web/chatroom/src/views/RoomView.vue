<template>
  <div className="container">
    <RoomsPanel></RoomsPanel>
  </div>
</template>

<script>
import RoomsPanel from '@/components/Rooms.vue'
// import { Icon  } from 'ant-design-vue';
export default {
  name: 'RoomView',
  components: {
    RoomsPanel
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
.container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100vh;

}
</style>
