<template>
  <a-popover
      placement="bottom"
      trigger="click"
      @visible-change="handlePopoverVisibleChange"
  >
    <template #content>
      <div>
        <p>已读用户：</p>
        <ul>
          <li v-for="user in messageUsers.readUsers" :key="user.id">
            {{ user.name }}
          </li>
        </ul>
        <p>未读用户：</p>
        <ul>
          <li v-for="user in messageUsers.unreadUsers" :key="user.id">
            {{ user.name }}
          </li>
        </ul>
      </div>
    </template>
  </a-popover>
</template>

<script>
export default {
  name: "ReadInfoPanel",
  data(){
    return{

    }
  },
  methods:{

  }
}
</script>

<style scoped lang="scss">

</style>
