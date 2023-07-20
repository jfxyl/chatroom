package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/http/services"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewMessageController() *MessageController {
	return &MessageController{
		MessageService: services.NewMessageService(),
	}
}

type MessageController struct {
	MessageService *services.MessageService
}

func (o *MessageController) Send(c *gin.Context) {
	var (
		err     error
		errs    map[string]string
		form    requests.MessageForm
		codeErr *common.CodeErr
		message *models.Message
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateMessageForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}

	if message, codeErr = o.MessageService.Send(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, message.Transform(auth.User(c).ID))
}

func (o *MessageController) List(c *gin.Context) {
	var (
		err         error
		idStr       string
		id          uint64
		minMsgIDStr string
		minMsgID    uint64
		codeErr     *common.CodeErr
		message     *models.Message
		messages    []*models.Message

		transformMessages = make([]map[string]any, 0)
	)
	idStr = c.Param("id")

	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	minMsgIDStr = c.Query("min_msg_id")
	if minMsgID, err = strconv.ParseUint(minMsgIDStr, 10, 64); err != nil {
		minMsgID = 0
	}
	if messages, codeErr = o.MessageService.List(c, id, minMsgID); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	for _, message = range messages {
		transformMessages = append(transformMessages, message.Transform(auth.User(c).ID))
	}
	common.RespOk(c, transformMessages)
}

func (o *MessageController) ReadInfo(c *gin.Context) {
	var (
		err     error
		idStr   string
		id      uint64
		codeErr *common.CodeErr
		message *models.Message
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	if message, codeErr = o.MessageService.ReadInfo(c, id); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, message.TransformReaders())
}
