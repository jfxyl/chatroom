package services

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"chatroom/internal/mgr"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewMessageService() *MessageService {
	return &MessageService{
		readChan: make(chan *common.ReadBody, 1000),
	}
}

type MessageService struct {
	readChan       chan *common.ReadBody
	AutoCommitChan chan *common.ReadBodyBatch

	msgReadMapChan map[uint64]chan *common.ReadBody
}

func (s *MessageService) PushMessage(message *models.Message) (err error) {
	var (
		msg []byte
	)
	//发送同步消息
	msg, err = json.Marshal(message)
	err = s.Push(common.MessageTopic, msg)
	return
}

func (s *MessageService) PushRead(readbody *common.ReadBody) (err error) {
	var (
		msg []byte
	)
	//发送同步消息
	msg, err = json.Marshal(readbody)
	err = s.Push(common.ReadTopic, msg)
	return
}

func (s *MessageService) Push(topic string, body []byte) (err error) {
	//发送同步消息
	err = mgr.G_MessageMgr.Push(topic, body)
	return
}

func (s *MessageService) Read(readBody common.ReadBody, userID uint64) (err error) {
	db.G_DB.Model(models.MessageRead{}).Where(map[string]any{
		"message_id": readBody.MessageID,
		"user_id":    userID,
	}).Updates(map[string]any{"read": true})
	return
}

func (s *MessageService) Send(c *gin.Context, form requests.MessageForm) (*models.Message, *common.CodeErr) {
	var (
		err     error
		message models.Message
	)
	//查询当前聊天室的所有用户
	var users []*models.UserRoom
	if db.G_DB.Model(&models.UserRoom{}).Where("room_id = ?", form.ReceiverID).Not("user_id = ?", auth.User(c).ID).Find(&users).Error != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	var reads []*models.MessageRead
	for _, user := range users {
		reads = append(reads, &models.MessageRead{UserID: user.UserID})
	}
	message = models.Message{
		ChatType:   models.ChatType(form.ChatType),
		MsgType:    models.MsgType(form.MsgType),
		SenderID:   auth.User(c).ID,
		ReceiverID: form.ReceiverID,
		Content:    form.Content,
		ReplyID:    form.ReplyID,
		ReaderInfo: reads,
	}
	if db.G_DB.Create(&message).Error != nil {
		fmt.Println("db.G_DB.Create(&message).Error", db.G_DB.Create(&message).Error)
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if err = s.PushMessage(&message); err != nil {
		fmt.Println("err", err)
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	message.Sender = auth.User(c)
	return &message, nil
}

func (s *MessageService) List(c *gin.Context, id uint64, minMsgID uint64) ([]*models.Message, *common.CodeErr) {
	var (
		err  error
		room models.Room
	)
	if err = db.G_DB.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		fmt.Println(minMsgID)
		if minMsgID > 0 {
			db.Where("id < ?", minMsgID)
		}
		return db.Order("id desc").Limit(20)
	}).Preload("Messages.Sender").Preload("Messages.ReaderInfo").Limit(1).Find(&room, id).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if room.ID == 0 {
		return nil, common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	//倒序
	for i, j := 0, len(room.Messages)-1; i < j; i, j = i+1, j-1 {
		room.Messages[i], room.Messages[j] = room.Messages[j], room.Messages[i]
	}
	return room.Messages, nil
}

func (s *MessageService) ReadInfo(c *gin.Context, id uint64) (*models.Message, *common.CodeErr) {
	var (
		err     error
		message models.Message
	)
	if err = db.G_DB.Preload("ReaderInfo.User").Limit(1).Find(&message, id).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if message.ID == 0 {
		return nil, common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	return &message, nil
}
