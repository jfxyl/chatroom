package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type MessageForm struct {
	ChatType   uint8  `form:"chat_type" json:"chat_type" valid:"chat_type"`
	ReceiverID uint64 `form:"receiver_id" json:"receiver_id" valid:"receiver_id"`
	MsgType    uint8  `form:"msg_type" json:"msg_type" valid:"msg_type"`
	Content    string `form:"content" json:"content" valid:"content"`
	ReplyID    uint64 `form:"reply_id" json:"reply_id" valid:"reply_id"`
}

func ValidateMessageForm(data MessageForm) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"chat_type":   []string{"required", "in:1,2"},
		"receiver_id": []string{"required"},
		"msg_type":    []string{"required", "in:1,2,3,4,5,6,7,8,9"},
		"content":     []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:消息类型错误",
			"in:消息类型错误",
		},
		"receiver_id": []string{
			"required:消息接收人错误",
		},
		"msg_type": []string{
			"in:消息类型错误",
		},
		"content": []string{
			"required:消息内容不能为空",
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
