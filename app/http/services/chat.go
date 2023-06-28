package services

import (
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"github.com/gin-gonic/gin"
)

func NewChatService() *ChatService {
	return &ChatService{}
}

type ChatService struct {
}

func (s *ChatService) List(c *gin.Context) ([]*common.Chat, *common.CodeErr) {
	var (
		err   error
		chats []*common.Chat
		user  models.User
	)
	if err = db.G_DB.Preload("UserRooms.Room.Users").Find(&user, auth.User(c).ID).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	for _, v := range user.UserRooms {
		chats = append(chats, &common.Chat{
			ID:        v.Room.ID,
			Name:      v.Room.Name,
			Alias:     v.Alias,
			Avatar:    v.Room.Avatar,
			CreatedAt: v.Room.CreatedAt.Format("2006-01-02 15:04:05"),
			Users:     v.Room.Users,
		})
	}
	return chats, nil
}
