package route

import (
	"chatroom/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	var (
		v1Group   *gin.RouterGroup
		wsGroup   *gin.RouterGroup
		userGroup *gin.RouterGroup
		msgGroup  *gin.RouterGroup
		roomGroup *gin.RouterGroup
	)
	v1Group = router.Group("/v1")
	wsGroup = v1Group.Group("/ws", middleware.AuthMiddleware())
	{
		wsGroup.GET("", nil)
	}
	userGroup = v1Group.Group("/users")
	{
		userGroup.POST("/register", nil)                        //用户注册
		userGroup.POST("/login", nil)                           //用户登录
		userGroup.POST("/logout", nil)                          //用户登出
		userGroup.GET("/:id", middleware.AuthMiddleware(), nil) //用户信息
		userGroup.PUT("/:id", middleware.AuthMiddleware(), nil) //用户信息修改
	}
	msgGroup = v1Group.Group("/messages", middleware.AuthMiddleware())
	{
		msgGroup.POST("", nil)          //发送消息
		msgGroup.POST("/:id/read", nil) //发送已读状态
		msgGroup.GET("/:id/read", nil)  //查看已读状态
	}
	roomGroup = v1Group.Group("/rooms", middleware.AuthMiddleware())
	{
		roomGroup.GET("", nil)             //用户拥有的聊天室
		roomGroup.POST("", nil)            //创建聊天室
		roomGroup.PUT("/:id", nil)         //修改聊天室
		roomGroup.GET("/:id", nil)         //聊天室信息
		roomGroup.GET("/:id/records", nil) //聊天室聊天记录
		roomGroup.POST("/:id/quit", nil)   //退出聊天室
		roomGroup.DELETE("/:id", nil)      //删除聊天室
		roomGroup.DELETE("/:id/join", nil) //加入聊天室
	}
}
