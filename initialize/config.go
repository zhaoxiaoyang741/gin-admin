package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zhaoxiaoyang741/gin-admin/config"
)

func InitConfig() config.Config {
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名称(不需要带后缀)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要指定配置文件的类型（可选）
	viper.AddConfigPath("./")     // 配置文件所在的路径
	err := viper.ReadInConfig()   // 读取配置文件
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
