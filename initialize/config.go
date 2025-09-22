package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zhaoxiaoyang741/gin-admin/config"
)

func InitConfig() config.Config {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}

	// 将配置文件中的内容映射到结构体中
	conf := config.Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("fail to decode into struct : %s", err.Error()))
	}

	return conf
}
