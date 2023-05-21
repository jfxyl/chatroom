package models

import "chatroom/app/models/base"

type MessageRead struct {
	base.BaseIDModel

	MessageID uint64 `gorm:"column:message_id;index;not null;"`
	Message   Message
	UserID    string `gorm:"column:user_id;type:varchar(20);index;not null;"`
	User      User

	base.BaseCreatedAtModel
}
