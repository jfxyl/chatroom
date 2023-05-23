package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type RegisterForm struct {
	Name            string `form:"name" json:"name" valid:"name"`
	Password        string `form:"password" json:"password" valid:"password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" valid:"confirm_password"`
	Gender          int8   `form:"gender" json:"gender" valid:"gender"`
	Birthday        string `form:"birthday" json:"birthday" valid:"birthday"`
	Avatar          string `form:"avatar" json:"avatar" valid:"avatar"`
}

type LoginForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=2,max=10"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=10"`
}

func ValidateUserForm(data RegisterForm) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":             []string{"required", "min:3", "max:10"},
		"password":         []string{"required", "min:6", "max:10", "alpha_num"},
		"confirm_password": []string{"required"},
		"gender":           []string{"in:0,1,2"},
		"birthday":         []string{"required", "date:yyyy-mm-dd"},
		"avatar":           []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"min:用户名长度为3-10位字符",
			"max:用户名长度为3-10位字符",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码为字母、数字组成的6-10位字符",
			"max:密码为字母、数字组成的6-10位字符",
			"alpha_num:密码为字母、数字组成的6-10位字符",
		},
		"gender": []string{
			"in:性别选择错误",
		},
		"birthday": []string{
			"required:出生日期为必填项",
			"date:出生日期格式不正确",
		},
		"avatar": []string{
			"required:用户头像为必填项",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 4. 开始验证
	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.ConfirmPassword {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}
