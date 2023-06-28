package models

import "chatroom/app/models/base"

type MessageRead struct {
	base.BaseIDModel

	MessageID uint64 `gorm:"column:message_id;uniqueIndex:index_message_user;not null;"`
	Message   Message
	UserID    uint64 `gorm:"column:user_id;uniqueIndex:index_message_user;not null;"`
	User      User
	Read      bool `gorm:"column:read;not null;default:false;"`

	base.BaseCreatedAtModel
}
