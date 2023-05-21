package bootstrap

import (
	"chatroom/internal/config"
	"flag"
)

func InitArgs() {
	flag.StringVar(&config.G_ConfigPath, "config", "./config.yaml", "指定配置文件，默认./config.yaml")
	flag.Parse()
}
