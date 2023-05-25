package middleware

import (
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization string
			token         *jwt.Token
			claims        jwt.MapClaims
			ok            bool
		)
		authorization = c.GetHeader("Authorization")
		authorization = strings.TrimPrefix(authorization, "Bearer ")
		if authorization == "" {
			common.RespAbort(c, http.StatusUnauthorized, common.ERR_UNAUTHORIZED)
			return
		}
		// 解析 JWT
		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.G_Config.Jwt.Secret), nil
		})
		if err != nil {
			common.RespAbort(c, http.StatusUnauthorized, common.ERR_UNAUTHORIZED)
			return
		}
		// 验证 JWT 是否有效
		if claims, ok = token.Claims.(jwt.MapClaims); !(ok && token.Valid) {
			common.RespAbort(c, http.StatusUnauthorized, common.ERR_UNAUTHORIZED)
			return
		}
		// 存储用户信息
		userID := int(claims["userID"].(float64))
		c.Set("userID", userID)
		c.Set("authInfo", &auth.Info{
			UserID: userID,
		})
		c.Next()
	}
}
