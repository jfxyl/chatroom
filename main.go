package main

import (
	"chatroom/bootstrap"
	"chatroom/internal/config"
	"chatroom/routes"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//go:embed public/*
var publicFS embed.FS

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
	routes.InitRouter(r, publicFS)
	if err = r.Run(fmt.Sprintf(":%d", config.G_Config.Port)); err != nil {
		log.Fatalln(err)
	}
}
