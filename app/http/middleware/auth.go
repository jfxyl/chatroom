package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	// 从请求头中获取用户的身份信息
	//	token := c.GetHeader("Authorization")
	//	if token == "" {
	//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	//			"message": "Authorization required",
	//		})
	//		return
	//	}
	//
	//	// 验证用户的身份信息是否有效
	//	userID, err := auth.VerifyToken(token)
	//	if err != nil {
	//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	//			"message": "Invalid authorization token",
	//		})
	//		return
	//	}
	//
	//	// 将用户的身份信息存储到上下文中，方便后续处理
	//	c.Set("userID", userID)
	//
	//	c.Next()
	//}
	return nil
}
