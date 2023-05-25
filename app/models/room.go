package models

import (
	"chatroom/app/models/base"
)

type Room struct {
	base.BaseIDModel

	Name     string `gorm:"column:name;type:char(20);not null;"`
	Owner    uint64 `gorm:"column:owner;not null;"`
	IsPublic bool   `gorm:"column:is_public;not null;default:false;"`
	Notice   string `gorm:"column:notice;type:varchar(2048);not null;"`

	Users []*User `gorm:"many2many:user_rooms;"`
	base.BaseTimeModel
}

func (m *Room) Transform() map[string]any {
	return map[string]any{
		"id":         m.ID,
		"name":       m.Name,
		"is_public":  m.IsPublic,
		"notice":     m.Notice,
		"created_at": m.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
