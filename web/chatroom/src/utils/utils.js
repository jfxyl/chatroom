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