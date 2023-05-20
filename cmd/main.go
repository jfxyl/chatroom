package main

import (
	"chatroom/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		r *gin.Engine
	)
	//初始化配置
	//初始化mysql
	//初始化rocketmq

	//初始化路由、http服务
	r = gin.Default()
	route.InitRouter(r)
	r.Run()
}
