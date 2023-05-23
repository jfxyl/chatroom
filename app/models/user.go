package models

import (
	"chatroom/app/models/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

const GenderUnknown = 0
const GenderMale = 1
const GenderFemale = 2

type User struct {
	base.BaseIDModel

	Name     string     `gorm:"column:name;type:varchar(20);not null;" json:"name"`
	Gender   int8       `gorm:"column:gender;not null;default:0;" json:"gender"`
	Avatar   string     `gorm:"column:avatar;type:varchar(50);not null;" json:"avatar"`
	Birthday *time.Time `gorm:"column:birthday;" json:"birthday"`
	Password string     `gorm:"column:password;type:varchar(100);not null;" json:"-"`

	base.BaseTimeModel
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	if len(m.Password) != 60 {
		var pwd []byte
		if pwd, err = bcrypt.GenerateFromPassword([]byte(m.Password), 14); err != nil {
			return
		}
		m.Password = string(pwd)
	}
	return
}

func (m *User) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (m *User) Transform() any {
	return map[string]any{
		"id":         m.ID,
		"name":       m.Name,
		"gender":     m.Gender,
		"avatar":     m.Avatar,
		"birthday":   m.Birthday.Format("2006-01-02"),
		"created_at": m.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
