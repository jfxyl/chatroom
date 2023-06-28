package common

import "chatroom/app/models"

type WsRequest struct {
	Action string `json:"action"`
	Body   any    `json:"body"`
}

type ReadBody struct {
	MessageID uint64 `json:"message_id"`
}

type Chat struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	Alias     string         `json:"alias"`
	Avatar    string         `json:"avatar"`
	CreatedAt string         `json:"created_at"`
	Users     []*models.User `json:"users"`
}
