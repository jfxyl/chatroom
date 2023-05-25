package auth

import (
	"chatroom/app/models"
	"chatroom/internal/global"
	"github.com/gin-gonic/gin"
)

type Info struct {
	UserID int
	User   *models.User
}

func User(c *gin.Context) *models.User {
	cc, exists := c.Get("authInfo")
	if exists {
		authInfo := cc.(*Info)
		if authInfo.User == nil {
			global.DB.Limit(1).Find(&authInfo.User, authInfo.UserID)
		}
		return authInfo.User
	}
	return nil
}
