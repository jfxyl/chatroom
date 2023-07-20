package models

import (
	"chatroom/app/models/base"
)

type ChatType uint8

const SingleChat ChatType = 1
const GroupChat ChatType = 2

type MsgType uint8

const TypeNotice MsgType = 1
const TypeText MsgType = 2
const TypeImg MsgType = 3
const TypeFile MsgType = 4
const TypeVoice MsgType = 5
const TypeVideo MsgType = 6
const TypeLocation MsgType = 7
const TypeChatRecords MsgType = 8
const TypeShare MsgType = 9

const OperateCreateRoom = "create_room"
const OperateJoinRoom = "join_room"
const OperateQuitRoom = "quit_room"

type Message struct {
	base.BaseIDModel

	ChatType     ChatType `gorm:"column:chat_type;not null;" json:"chat_type"`
	SenderID     uint64   `gorm:"column:sender_id;not null;index" json:"sender_id"`
	ReceiverID   uint64   `gorm:"column:receiver_id;not null;index" json:"receiver_id"`
	MsgType      MsgType  `gorm:"column:msg_type;not null;" json:"msg_type"`
	Content      string   `gorm:"column:content;type:varchar(2048);not null;" json:"content"`
	ReplyID      uint64   `gorm:"column:reply_id;index" json:"reply_id"`
	ReplyMessage *Message `gorm:"-"`
	Revoked      uint8    `gorm:"column:revoked;default:0;comment:是否撤回" json:"revoked"`
	Operate      string   `gorm:"column:operate;type:varchar(30);" json:"operate"`

	Sender     *User          `gorm:"foreignKey:sender_id"`
	ReaderInfo []*MessageRead `gorm:"foreignKey:message_id"`
	Readers    []*User        `gorm:"many2many:message_reads;"`

	base.BaseCreatedAtModel
}

func (m *Message) Transform(userID uint64) map[string]any {
	var sender map[string]any
	if m.Sender != nil {
		sender = m.Sender.Transform()
	}
	var readers = []uint64{}
	var unreaders = []uint64{}
	var readed = false
	for _, reader := range m.ReaderInfo {
		if reader.UserID == userID && reader.Read {
			readed = true
		}
		if reader.Read {
			readers = append(readers, reader.UserID)
		} else {
			unreaders = append(unreaders, reader.UserID)
		}
	}
	if m.SenderID == userID {
		readed = true
	}
	return map[string]any{
		"id":          m.ID,
		"chat_type":   m.ChatType,
		"msg_type":    m.MsgType,
		"sender_id":   m.SenderID,
		"receiver_id": m.ReceiverID,
		"sender":      sender,
		"readers":     readers,
		"unreaders":   unreaders,
		"readed":      readed,
		"content":     m.Content,
		"reply_id":    m.ReplyID,
		"revoked":     m.Revoked,
		"operate":     m.Operate,
		"created_at":  m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (m *Message) TransformReaders() map[string]any {
	var readers, unreaders = make([]map[string]any, 0), make([]map[string]any, 0)

	for _, reader := range m.ReaderInfo {
		if reader.Read {
			readers = append(readers, reader.User.Transform())
		} else {
			unreaders = append(unreaders, reader.User.Transform())
		}
	}
	return map[string]any{
		"readers":   readers,
		"unreaders": unreaders,
	}
}
