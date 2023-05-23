package common

import (
	"chatroom/app/http/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespOk(c *gin.Context, data any) {
	Resp(c, http.StatusOK, 0, data, "success")
}

func RespFail(c *gin.Context, status int, msg any) {
	Resp(c, status, 1, nil, msg)
}

func RespErrs(c *gin.Context, status int, errs map[string]string) {
	Resp(c, status, 1, nil, "")
}

func Resp(c *gin.Context, status int, errcode int, data any, msg any) {
	c.JSON(status, gin.H{
		"errcode": errcode,
		"data":    data,
		"msg":     msg,
	})
}

// 简化错误返回
//func SimplifyError(errsMapping map[string][]string) error {
//	for _, errs := range errsMapping {
//		if len(errs) > 0 {
//			return errors.New(errs[0])
//		}
//	}
//	return nil
//}

func SimplifyError(errsMapping map[string][]string) map[string]string {
	var errMapping = make(map[string]string)
	for field, errs := range errsMapping {
		if len(errs) > 0 {
			errMapping[field] = errs[0]
		}
	}
	if len(errMapping) > 0 {
		return errMapping
	}
	return nil
}

func JsonReqValidate(c *gin.Context, form requests.RegisterForm) any {
	var (
		err  error
		errs map[string]string
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		return err
	}
	fmt.Println("form", form)
	if errs = SimplifyError(requests.ValidateUserForm(form)); errs != nil && len(errs) > 0 {
		return errs
	}
	return nil
}
