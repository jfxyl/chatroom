package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type RoomForm struct {
	Name     string `form:"name" json:"name" valid:"name"`
	Avatar   string `form:"avatar" json:"avatar" valid:"avatar"`
	IsPublic uint8  `form:"is_public" json:"is_public" valid:"is_public"`
}

func ValidateCreateRoomForm(data RoomForm) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":      []string{"required", "min_cn:3", "max_cn:10", "softdel_not_exists:rooms,name"},
		"is_public": []string{"in:0,1"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:房间名为必填项",
			"min_cn:房间名为字母、数字组成的3-10位字符",
			"max_cn:房间名为字母、数字组成的3-10位字符",
			"softdel_not_exists:房间名已存在",
		},
		"is_public": []string{
			"in:房间类型选择错误",
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
	return govalidator.New(opts).ValidateStruct()
}
