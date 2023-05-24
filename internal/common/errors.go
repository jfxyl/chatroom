package common

import "errors"

var (
	ERR_UNAUTHORIZED = errors.New("未认证")

	ERR_FORBIDDEN = errors.New("无权限")

	ERR_INTERNAL_SERVER = errors.New("内部错误")

	ERR_NOT_FOUND = errors.New("资源不存在")
)
