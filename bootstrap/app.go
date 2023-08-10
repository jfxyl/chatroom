package bootstrap

import (
	"chatroom/internal/config"
	"chatroom/internal/db"
	"chatroom/internal/mgr"
	"chatroom/internal/mq"
	"chatroom/internal/oss"
	"fmt"
)

func BootStrap() (err error) {
	//初始化命令行参数
	InitArgs()
	//初始化配置
	if err = config.InitConfig(); err != nil {
		return
	}
	fmt.Println("初始化config成功")
	//初始化数据库
	if err = db.InitDatabase(); err != nil {
		return
	}
	fmt.Println("初始化db成功")
	//初始化oss
	if err = oss.InitOss(); err != nil {
		return
	}
	fmt.Println("初始化oss成功")
	//初始化rocketmq
	if err = mq.InitMQ(); err != nil {
		return
	}
	fmt.Println("初始化mq成功")
	mgr.InitWsMgr()
	fmt.Println("初始化ws mgr成功")
	mgr.InitMessageMgr()
	fmt.Println("初始化message mgr成功")
	return
}
