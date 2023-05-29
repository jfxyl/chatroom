<template>
  <div class="register-form">
    <a-button type="link" class="back" @click="$emit('jump', 'login')">
      返回
    </a-button>
    <div class="form_title">ChatRoom</div>
    <a-form :form="form" layout="vertical" @submit="handleSubmit">
      <a-form-item label="用户名">
        <a-input
            class="ant-input-sm"
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
            class="ant-input-sm"
            placeholder="请输入密码"
            type="password"
            allow-clear
            v-decorator="[
              'password',
              { rules: [{ required: true, message: '请输入密码' }] },
            ]"
        ></a-input>
      </a-form-item>
      <a-form-item label="确认密码">
        <a-input
            class="ant-input-sm"
            placeholder="请输入密码"
            type="password"
            allow-clear
            v-decorator="[
              'confirm_password',
              { rules: [
                  { required: true, message: '请输入确认密码' },
                  {validator: compareToFirstPassword,},
              ] },
            ]"
        ></a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit" block >注册</a-button>
      </a-form-item>
    </a-form>
  </div>

</template>

<script>
import {Form, Input, Button, message} from 'ant-design-vue';
import {formErrorPrompt} from '../utils/utils.js';
export default {
  name: 'RegisterPanel',
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
    compareToFirstPassword(rule, value, callback) {
      const form = this.form;
      if (value && value !== form.getFieldValue('password')) {
        callback('两次密码输入不一致!');
      } else {
        callback();
      }
    },
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, form) => {
        if (!err) {
          let that = this
          console.log('Received values of form: ', form);
          that.$http.post('/v1/users/register',form)
              .then(function(data){
                if(data.data.errcode !== 0){
                  formErrorPrompt(that.form,form,data.data.msg)
                }else{
                  message.success('注册成功');
                }
              })
        }
      });
      // this.$router.push('/chatroom');
    },
  },
}
</script>

<style scoped>
.register-form{
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
  .ant-form-item{
    //margin-bottom: 15px;
  }
  .back{
    position: absolute;
    top: 14px;
    left: 44px;
    font-size: 12px;
    padding: 0;
    color: rgb(119, 119, 119);
  }
}
</style>
