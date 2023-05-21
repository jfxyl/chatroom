package models

import (
	"chatroom/app/models/base"
)

type chatType uint8

const singleChat chatType = 1
const groupChat chatType = 2

type messageType uint8

const typeNotice messageType = 1
const typeText messageType = 2
const typeImg messageType = 3
const typeFile messageType = 4
const typeVoice messageType = 5
const typeVideo messageType = 6
const typeLocation messageType = 7
const typeChatRecords messageType = 8
const typeShare messageType = 9

type Message struct {
	base.BaseIDModel

	ChatType     chatType    `gorm:"column:chat_type;not null;"`
	SenderID     uint64      `gorm:"column:sender_id;not null;index"`
	ReceiverID   uint64      `gorm:"column:receiver_id;not null;index"`
	Type         messageType `gorm:"column:type;not null;"`
	Content      string      `gorm:"column:content;type:varchar(2048);not null;"`
	ReplyID      uint64      `gorm:"column:reply_id;index"`
	ReplyMessage *Message    `gorm:"-"`
	Revoked      uint8       `gorm:"column:revoked;default:0;comment:是否撤回"`

	base.BaseCreatedAtModel
}
