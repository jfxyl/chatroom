package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type jwtConfig struct {
	Expiration int    `mapstructure:"expiration"`
	Secret     string `mapstructure:"secret"`
}

type mysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type rocketmqConfig struct {
	Endpoints  []string `mapstructure:"endpoints"`
	BrokerAddr string   `mapstructure:"brokeraddr"`
}

type ossConfig struct {
	Entpoint        string `mapstructure:"entpoint"`
	BucketName      string `mapstructure:"bucket_name"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	PrefixDir       string `mapstructure:"prefix_dir"`
}

type config struct {
	Name     string         `mapstructure:"name"`
	Port     int            `mapstructure:"port"`
	Jwt      jwtConfig      `mapstructure:"jwt"`
	Mysql    mysqlConfig    `mapstructure:"mysql"`
	Rocketmq rocketmqConfig `mapstructure:"rocketmq"`
	Oss      ossConfig      `mapstructure:"oss"`
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

// 读取配置文件
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
	fmt.Printf("G_Config:%+v", *G_Config)
	return
}
