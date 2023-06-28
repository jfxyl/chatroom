package controllers

import (
	"chatroom/app/http/services"
	"chatroom/internal/common"
	"github.com/gin-gonic/gin"
)

type OssController struct {
	OssService *services.OssService
}

func (o *OssController) Signature(c *gin.Context) {
	var (
		err         error
		uploadUrl   string
		url         string
		dirtype     string
		contentType string
		codeErr     *common.CodeErr
	)
	dirtype = c.Query("dirtype")
	contentType = c.Query("content-type")
	if uploadUrl, codeErr = o.OssService.Signature(dirtype, contentType); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	//将signURL ?后面的参数去掉
	if url, err = common.GetLinkWithoutParams(uploadUrl); err != nil {
		common.RespFail(c, common.StatusInternal, err.Error())
		return
	}
	common.RespOk(c, gin.H{
		"uploadUrl": uploadUrl,
		"url":       url,
	})
}
