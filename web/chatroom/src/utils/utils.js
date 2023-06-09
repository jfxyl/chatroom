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


export function getContentTypeFromBase64(base64Data) {
    const base64Header = /^data:([A-Za-z-+/]+);base64,/;

    const matches = base64Data.match(base64Header);
    if (matches && matches.length > 1) {
        return matches[1];
    }

    return null;
}


export function b64toBlob(b64Data, contentType='', sliceSize=512) {
    b64Data = b64Data.split(',').pop();
    const byteCharacters = atob(b64Data);
    const byteArrays = [];
    for (let offset = 0; offset < byteCharacters.length; offset += sliceSize) {
        const slice = byteCharacters.slice(offset, offset + sliceSize);

        const byteNumbers = new Array(slice.length);
        for (let i = 0; i < slice.length; i++) {
            byteNumbers[i] = slice.charCodeAt(i);
        }

        const byteArray = new Uint8Array(byteNumbers);

        byteArrays.push(byteArray);
    }

    const blob = new Blob(byteArrays, {type: contentType});
    return blob;
}