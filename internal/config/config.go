package config

type mysqlConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type rocketmqConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type config struct {
	Name     string         `mapstructure:"name"`
	Mysql    mysqlConfig    `mapstructure:"mysql"`
	Rocketmq rocketmqConfig `mapstructure:"mysql"`
}
