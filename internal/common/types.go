package common

import "chatroom/app/models"

type WsRequest struct {
	Action string `json:"action"`
	Body   any    `json:"body"`
}

type ReadBody struct {
	UserId    uint64 `json:"user_id"`
	MessageID uint64 `json:"message_id"`
}

type Chat struct {
	ID          uint64         `json:"id"`
	Name        string         `json:"name"`
	Alias       string         `json:"alias"`
	Avatar      string         `json:"avatar"`
	Owner       uint64         `json:"owner"`
	CreatedAt   string         `json:"created_at"`
	UnreadCount uint64         `json:"unread_count"`
	Users       []*models.User `json:"users"`
}

type ReadBodyBatch struct {
	ReadBodyList []*ReadBody
}
