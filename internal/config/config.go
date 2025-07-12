package config

import (
	"github.com/spf13/viper"
)

// Config 存储应用配置
type Config struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
