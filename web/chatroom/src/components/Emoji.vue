<template>
  <div id="emoji" @click.stop>
    <ul class="emojiul">
      <li class="emojili" v-for="(emoji,index) in emojis" :key="index" @click.stop="chooseEmoji(emoji)">
        {{emoji}}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "EmojiPanel",
  data(){
    return{
      emojis:["😄","😃","😀","😊","😉","😍","😘","😚","😗","😙","😜","😝","😛","😳","😁","😔","😌","😒","😞","😣","😢","😂","😭","😪","😥","😰","😅","😓","😩","😫","😨","😱","😠","😡","😤","😖","😆","😋","😷","😎","😴","😵","😲","😟","😦","😧","👿","😮","😬","😐","😕","😯","😏","😑","👲","👳","👮","👷","💂","👶","👦","👧","👨","👩","👴","👵","👱","👼","👸","😺","😸","😻","😽","😼","🙀","😿","😹","😾","👹","👺","🙈","🙉","🙊","💀","👽","💩","🔥","✨","🌟","💫","💥","💢","💦","💧","💤","💨","👂","👀","👃","👅","👄","👍","👎","👌","👊","✊","👋","✋","👐","👆","👇","👉","👈","🙌","🙏","👏","💪","🚶","🏃","💃","👫","👪","💏","💑","👯","🙆","🙅","💁","🙋","💆","💇","💅","👰","🙎","🙍","🙇","🎩","👑","👒","👟","👞","👡","👠","👢","👕","👔","👚","👗","🎽","👖","👘","👙","💼","👜","👝","👛","👓","🎀","🌂","💄","💛","💙","💜","💚","💔","💗","💓","💕","💖","💞","💘","💌","💋","💍","💎","👤","💬","👣"]
    }
  },
  methods:{
    chooseEmoji(emoji){
      this.$emit('chooseEmoji',emoji)
    }
  }
}
</script>

<style scoped>
#emoji{
  cursor: pointer;
  background:#fff;
  padding:10px;
  max-width:340px;
  border-radius: 3px;
  box-shadow:0px 0px 5px #333333;
  max-height:300px;
  overflow-y: scroll;
  .emojiul{
    list-style: none;
    height:100%;
    overflow: hidden;
    margin: 0;
    .emojili{
      display: block;
      height:25px;
      width:25px;
      padding:5px;
      float:left;
    }
  }
}

</style>
