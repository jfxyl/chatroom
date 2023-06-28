package common

import "errors"

type errCode int

type CodeErr struct {
	Code errCode
	Err  error
}

func NewCodeErr(code errCode, err error) *CodeErr {
	return &CodeErr{
		Code: code,
		Err:  err,
	}
}

const (
	//此处不使用iota，是为了方便后续扩展
	StatusOK              errCode = 0
	StatusUnauthorized    errCode = 1
	StatusInvalidArgument errCode = 2
	StatusForbidden       errCode = 3
	StatusNotFound        errCode = 4
	StatusInternal        errCode = 5
	StatusAlreadyExists   errCode = 6
)

var (
	ERR_UNAUTHORIZED = errors.New("未认证")

	ERR_FORBIDDEN = errors.New("无权限")

	ERR_INTERNAL_SERVER = errors.New("内部错误")

	ERR_NOT_FOUND = errors.New("资源不存在")
)
