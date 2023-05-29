package models

import (
	"chatroom/app/models/base"
	"chatroom/internal/config"
	"github.com/golang-jwt/jwt/v5"
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
	Nickname string     `gorm:"column:nickname;type:varchar(20);not null;" json:"nickname"`
	Gender   int8       `gorm:"column:gender;not null;default:0;" json:"gender"`
	Avatar   string     `gorm:"column:avatar;type:varchar(50);not null;" json:"avatar"`
	Birthday *time.Time `gorm:"column:birthday;" json:"birthday"`
	Password string     `gorm:"column:password;type:varchar(100);not null;" json:"-"`

	Rooms []*Room `gorm:"many2many:user_rooms;"`
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

func (m *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

func (m *User) Transform() map[string]any {
	var birthday = ""
	if m.Birthday != nil {
		birthday = m.Birthday.Format("2006-01-02")
	}
	return map[string]any{
		"id":         m.ID,
		"name":       m.Name,
		"nickname":   m.Nickname,
		"gender":     m.Gender,
		"avatar":     m.Avatar,
		"birthday":   birthday,
		"created_at": m.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (m *User) LoginTransform(token string) any {
	var (
		userinfo map[string]any
		jwtinfo  map[string]any
	)
	userinfo = m.Transform()
	jwtinfo = map[string]any{
		"token":      token,
		"expired_at": time.Now().Add(time.Duration(config.G_Config.Jwt.Expiration) * time.Second).Unix(),
	}
	return map[string]any{
		"user": userinfo,
		"jwt":  jwtinfo,
	}
}

func (m *User) GenerateJWT() (string, error) {
	var (
		err         error
		token       *jwt.Token
		tokenString string
	)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    config.G_Config.Name,
		"exp":    time.Now().Add(time.Duration(config.G_Config.Jwt.Expiration) * time.Second).Unix(),
		"iat":    time.Now().Unix(),
		"userID": m.ID,
		"user":   m.Transform(),
	})
	tokenString, err = token.SignedString([]byte(config.G_Config.Jwt.Secret))
	return tokenString, err
}
