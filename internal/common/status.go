package common

type ErrCode int

const (
	//此处不使用iota，是为了方便后续扩展
	StatusOK              ErrCode = 0
	StatusUnauthorized    ErrCode = 1
	StatusInvalidArgument ErrCode = 2
	StatusForbidden       ErrCode = 3
	StatusNotFound        ErrCode = 4
	StatusInternal        ErrCode = 5
	StatusAlreadyExists   ErrCode = 6
)
