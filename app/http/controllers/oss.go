package controllers

import (
	"chatroom/internal/common"
	"chatroom/internal/oss"
	"github.com/gin-gonic/gin"
)

type OssController struct {
}

func (o *OssController) Signature(c *gin.Context) {
	var (
		err         error
		uploadUrl   string
		url         string
		dirtype     string
		contentType string
	)
	dirtype = c.Query("dirtype")
	contentType = c.Query("content-type")
	if uploadUrl, err = oss.Signature(dirtype, contentType); err != nil {
		common.RespFail(c, common.StatusInternal, err.Error())
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
