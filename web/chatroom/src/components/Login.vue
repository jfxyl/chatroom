<template>

  <div class="login-form">
    <div class="form_title">ChatRoom</div>
    <a-form  :form="form" layout="vertical"  @submit="handleSubmit">
      <a-form-item label="用户名">
        <a-input
            placeholder="请输入用户名"
            allow-clear
            v-decorator="[
              'name',
              { rules: [{ required: true, message: '请输入用户名' }] },
            ]"
        ></a-input>
      </a-form-item>
      <a-form-item label="密码">
        <a-input
            placeholder="请输入密码"
            type="password"
            allow-clear
            v-decorator="[
              'password',
              { rules: [{ required: true, message: '请输入密码' }] },
            ]"
        ></a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit" block>登录</a-button>
      </a-form-item>
    </a-form>
    <div class="access_bottom">
<!--      <a-button type="link">-->
<!--        忘记密码-->
<!--      </a-button>-->
      <span></span>
      <a-button type="link" @click="$emit('jump', 'register')">
        立即注册
      </a-button>
    </div>
  </div>
</template>

<script>
import { Form,  Input, Button,message  } from 'ant-design-vue';
import {formErrorPrompt} from '../utils/utils.js';

export default {
  name: 'LoginPanel',
  components:{
    'a-form': Form,
    'a-form-item': Form.Item,
    'a-input': Input,
    'a-button': Button,
  },
  data(){
    return{
      form: this.$form.createForm(this, { name: 'coordinated' }),
    }
  },
  methods:{
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, form) => {
        if (!err) {
          let that = this
          that.$http.post('/v1/users/login',form)
          .then(function(data){
            if(data.data.errcode !== 0){
              formErrorPrompt(that.form,form,data.data.msg)
            }else{
              localStorage.setItem('token','Bearer '+data.data.data.jwt.token)
              localStorage.setItem('expired_at',data.data.data.jwt.expired_at * 1000)
              message.success('登录成功');
            }
          })
        }
      });
      // this.$router.push('/chatroom');
    },
  }
}
</script>

<style scoped>
.login-form{
  border-radius: 10px;
  box-shadow: 0 0 30px rgba(0,0,0,.1);
  padding: 47px 44px;
  position: relative;
  width: 325px;
  .form_title{
    font-size: 16px;
    font-weight: 500;
    padding-top: 0;
    padding-bottom: 24px;
  }
  .access_bottom{
    font-size: 12px;
    margin-top: 4px;
    display: flex;
    justify-content: space-between;
  }
}
</style>
