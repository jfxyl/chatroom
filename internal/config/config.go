package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type mysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type rocketmqConfig struct {
	Endpoints []string `mapstructure:"endpoints"`
}

type config struct {
	Name     string         `mapstructure:"name"`
	Mysql    mysqlConfig    `mapstructure:"mysql"`
	Rocketmq rocketmqConfig `mapstructure:"db"`
}

var (
	G_Config     *config
	G_ConfigPath = "./config.yaml"
)

func InitConfig() (err error) {
	viper.New()

	var (
		v *viper.Viper
	)
	G_Config = &config{}
	v = viper.New()
	v.SetConfigFile(G_ConfigPath)
	//读取配置文件
	if err = readConfig(v); err != nil {
		return
	}
	//监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		//编辑器可能会触发两次事件
		fmt.Printf("监听到文件变化：%s：%s", e.Name, e.Op)
		//读取配置文件
		if err = readConfig(v); err != nil {
			return
		}
		//do something
		fmt.Println("do something")
	})
	return
}

//读取配置文件
func readConfig(v *viper.Viper) (err error) {
	var (
		config config
	)
	//读取配置文件
	if err = v.ReadInConfig(); err != nil {
		return
	}
	//解析配置文件(不直接使用Config去接收解析结果，在解析成功后再赋值，避免配置错误影响原来的配置)
	if err = v.Unmarshal(&config); err != nil {
		return
	}
	G_Config = &config
	return
}
