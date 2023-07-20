package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func RespOk(c *gin.Context, data any) {
	Resp(c, http.StatusOK, StatusOK, data, "success")
}

func RespFail(c *gin.Context, errcode errCode, msg any) {
	Resp(c, http.StatusOK, errcode, nil, msg)
}

func RespAbort(c *gin.Context, errcode errCode, msg any) {
	c.Abort()
	Resp(c, http.StatusOK, errcode, nil, msg)
}

//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//	"message": "Invalid token",
//})
func Resp(c *gin.Context, status int, errcode errCode, data any, msg any) {
	switch v := msg.(type) {
	case error:
		msg = v.Error()
	}
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

//随机生成指定长度的字符串
func RandString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
		//result = append(result, bytes[1])
	}
	return string(result)
}

func RandInt(l int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(12)
}

//func RandString(){
//
//}

func GenerateRandomFilename() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s", RandString(16))
}

func GetLinkWithoutParams(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	u.RawQuery = "" // 清除查询参数部分
	u.Fragment = "" // 清除片段部分

	return u.String(), nil
}
