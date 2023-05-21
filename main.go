package main

import (
	"chatroom/bootstrap"
	"chatroom/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var (
		err error
		r   *gin.Engine
	)
	//初始化
	if err = bootstrap.BootStrap(); err != nil {
		log.Fatalln(err)
	}
	//初始化路由、http服务
	r = gin.Default()
	routes.InitRouter(r)
	r.Run()
}
