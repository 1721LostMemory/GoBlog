package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// 使用 viper 加载配置文件
	viper.SetConfigName("config") // 配置文件名 (不需要扩展名)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // 配置文件所在路径

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
func GetSessionkey() string {
	LoadConfig()
	fmt.Println("session key:", viper.GetString("session.key"))
	return viper.GetString("session.key")
}
