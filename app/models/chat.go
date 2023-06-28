package models

//
//import (
//	"chatroom/app/models/base"
//	"gorm.io/gorm"
//)
//
//type Chat struct {
//	base.BaseIDModel
//
//	ChatType string `gorm:"column:chat_type;not null;"`
//	UserID   uint64 `gorm:"column:user_id;uniqueIndex:index_user_room;not null;"`
//	User     *User
//
//	Avatar   string `gorm:"column:avatar;type:varchar(100);not null;" json:"avatar"`
//	Owner    uint64 `gorm:"column:owner;not null;"`
//	IsPublic bool   `gorm:"column:is_public;not null;default:false;"`
//	Notice   string `gorm:"column:notice;type:varchar(2048);not null;"`
//
//	Users []*User `gorm:"many2many:user_rooms;"`
//	base.BaseTimeModel
//}
//
//func (m *Chat) BeforeCreate(tx *gorm.DB) (err error) {
//	if m.Avatar == "" {
//		m.Avatar = defaultAvatar
//	}
//	return
//}
//
//func (m *Chat) Transform() map[string]any {
//	avatar := m.Avatar
//	if avatar == "" {
//		avatar = defaultAvatar
//	}
//	return map[string]any{
//		"id":         m.ID,
//		"name":       m.Name,
//		"avatar":     avatar,
//		"is_public":  m.IsPublic,
//		"notice":     m.Notice,
//		"created_at": m.CreatedAt.Format("2006-01-02 15:04:05"),
//		"updated_at": m.UpdatedAt.Format("2006-01-02 15:04:05"),
//	}
//}
