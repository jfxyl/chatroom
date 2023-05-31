import {message} from "ant-design-vue";

export function formErrorPrompt(formObj,formVal,errs){
    if(typeof errs === 'string'){
        message.error(errs);
    }else{
        var errors = {}
        for(var field in errs){
            errors[field] = {
                value: formVal[field],
                errors: [new Error(errs[field])]
            }
        }
        formObj.setFields(errors);
    }
}

// function isLoggedIn(){
//     return localStorage.getItem('token') && localStorage.getItem('expired_at') > Date.now();
// }


export async function getUserInfo(_this){
    try {
        const response = await _this._vm.$http.get('/v1/users');
        if(response.data.errcode !== 0){
            message.error(response.data.msg);
        }else{
            console.log('_this',_this)
            _this.dispatch('LOGIN',response.data.data)
        }
        // 处理响应数据
        console.log(response.data);
    } catch (error) {
        // 处理错误
        console.error(error);
    }
}