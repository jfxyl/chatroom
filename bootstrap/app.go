package bootstrap

import (
	"chatroom/internal/config"
	"chatroom/internal/db"
)

func BootStrap() (err error) {
	//初始化命令行参数
	InitArgs()
	//初始化配置
	config.InitConfig()
	//初始化数据库
	if err = db.InitDatabase(); err != nil {
		return
	}
	//初始化rocketmq
	//if err = mq.InitMQ(); err != nil {
	//	return
	//}
	return
}
