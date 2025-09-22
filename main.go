package main

import (
	"github.com/zhaoxiaoyang741/gin-admin/core"
	"github.com/zhaoxiaoyang741/gin-admin/global"
	"github.com/zhaoxiaoyang741/gin-admin/initialize"
)

func main() {
	// 初始化读取后台配置
	global.Config = initialize.InitConfig()
	// initialize.InitServer()
	global.Logger = core.Zap()
	global.Logger.Info("gin-admin start success----------------------")
}
