package models

import (
	"chatroom/app/models/base"
	"time"
)

const GenderUnknown = 0
const GenderMale = 1
const GenderFemale = 2

type User struct {
	base.BaseIDModel

	Name     string     `gorm:"column:name;type:varchar(20);not null;"`
	Gender   int8       `gorm:"column:gender;not null;default:0;"`
	Avatar   string     `gorm:"column:avatar;type:varchar(50);not null;"`
	Birthday *time.Time `gorm:"column:birthday;"`
	Password string     `gorm:"column:password;type:varchar(100);not null;"`

	base.BaseTimeModel
}

//func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
//
//}
