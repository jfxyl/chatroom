package services

import (
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"fmt"
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
	var result []struct {
		ReceiverID uint64
		Count      uint64
	}
	var resultMap = make(map[uint64]uint64)
	if len(user.UserRooms) > 0 {
		db.G_DB.Model(&models.Message{}).
			Joins("JOIN message_reads ON messages.id = message_reads.message_id").
			Where("message_reads.user_id = ?", auth.User(c).ID).
			Where("message_reads.read = 0").
			Group("receiver_id").
			Select("receiver_id, COUNT(1) count").
			Scan(&result)
		for _, r := range result {
			resultMap[r.ReceiverID] = r.Count
		}
		fmt.Println(result)
	}
	for _, v := range user.UserRooms {
		chats = append(chats, &common.Chat{
			ID:          v.Room.ID,
			Name:        v.Room.Name,
			Alias:       v.Alias,
			Avatar:      v.Room.Avatar,
			CreatedAt:   v.Room.CreatedAt.Format("2006-01-02 15:04:05"),
			Users:       v.Room.Users,
			UnreadCount: resultMap[v.Room.ID],
		})
	}
	return chats, nil
}
