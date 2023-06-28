package controllers

import (
	"chatroom/app/http/services"
	"chatroom/internal/common"
	"github.com/gin-gonic/gin"
)

func NewChatController() *ChatController {
	return &ChatController{
		ChatService: services.NewChatService(),
	}
}

type ChatController struct {
	ChatService *services.ChatService
}

func (o *ChatController) List(c *gin.Context) {
	var (
		chats   []*common.Chat
		codeErr *common.CodeErr
	)
	if chats, codeErr = o.ChatService.List(c); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, chats)
}
