package models

import (
	"chatroom/app/models/base"
)

type UserRoom struct {
	base.BaseIDModel

	UserID uint64 `gorm:"column:user_id;uniqueIndex:index_user_room;not null;"`
	User   *User

	RoomID uint64 `gorm:"column:room_id;uniqueIndex:index_user_room;not null;"`
	Room   *Room

	Alias string `gorm:"column:alias;type:varchar(100);" json:"alias"`

	base.BaseCreatedAtModel
}
