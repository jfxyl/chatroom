package models

import (
	"chatroom/app/models/base"
)

type Room struct {
	base.BaseIDModel

	Name   string `gorm:"column:name;type:char(20);not null;"`
	Owner  string `gorm:"column:owner;type:char(20);not null;"`
	Notice string `gorm:"column:notice;type:varchar(2048);not null;"`

	base.BaseTimeModel
}
