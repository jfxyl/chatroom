package routes

import (
	"chatroom/app/http/controllers"
	"chatroom/app/http/middleware"
	"chatroom/internal/common"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	var (
		v1Group   *gin.RouterGroup
		wsGroup   *gin.RouterGroup
		userGroup *gin.RouterGroup
		msgGroup  *gin.RouterGroup
		roomGroup *gin.RouterGroup
		chatGroup *gin.RouterGroup
		ossGroup  *gin.RouterGroup

		roomController    *controllers.RoomController
		chatController    *controllers.ChatController
		userController    *controllers.UserController
		ossController     *controllers.OssController
		wsController      *controllers.WsController
		messageController *controllers.MessageController
	)

	v1Group = router.Group("/v1")
	wsGroup = v1Group.Group("/ws", middleware.AuthMiddleware())
	userGroup = v1Group.Group("/users")
	msgGroup = v1Group.Group("/messages", middleware.AuthMiddleware())
	roomGroup = v1Group.Group("/rooms", middleware.AuthMiddleware())
	chatGroup = v1Group.Group("/chats", middleware.AuthMiddleware())
	ossGroup = v1Group.Group("/oss")

	wsController = controllers.NewWsController()
	{
		wsGroup.GET("", wsController.Conn)
	}

	userController = controllers.NewUserController()
	{
		userGroup.POST("/register", userController.Create)                      //用户注册
		userGroup.POST("/login", userController.Login)                          //用户登录
		userGroup.POST("/logout", userController.Logout)                        //用户登出
		userGroup.GET("", middleware.AuthMiddleware(), userController.Info)     //用户信息
		userGroup.GET("/:id", middleware.AuthMiddleware(), userController.Info) //用户信息
		userGroup.PUT("", middleware.AuthMiddleware(), userController.Update)   //用户信息修改
	}

	messageController = controllers.NewMessageController()
	{
		msgGroup.POST("", messageController.Send)                 //发送消息
		msgGroup.GET("/:id", messageController.List)              //消息列表
		msgGroup.GET("/:id/readinfo", messageController.ReadInfo) //查看已读状态
	}

	roomController = controllers.NewRoomController()
	{
		roomGroup.GET("", roomController.List)                 //用户拥有的聊天室
		roomGroup.POST("", roomController.Create)              //创建聊天室
		roomGroup.PUT("/:id", nil)                             //修改聊天室
		roomGroup.GET("/:id", roomController.Info)             //聊天室信息
		roomGroup.GET("/:id/messages", messageController.List) //聊天室聊天记录
		roomGroup.POST("/:id/quit", roomController.Quit)       //退出聊天室
		roomGroup.DELETE("/:id", roomController.Delete)        //删除聊天室
		roomGroup.POST("/:id/join", roomController.Join)       //加入聊天室
	}

	chatController = controllers.NewChatController()
	{
		chatGroup.GET("", chatController.List) //oss签名
	}

	ossController = new(controllers.OssController)
	{
		ossGroup.GET("/signature", ossController.Signature) //oss签名
	}
	router.NoRoute(func(c *gin.Context) {
		acceptHeader := c.GetHeader("Accept")
		if acceptHeader == "application/json" {
			common.RespAbort(c, common.StatusNotFound, common.ERR_NOT_FOUND)
		} else {
			c.String(404, "404!!!")
		}
		return
	})
}
