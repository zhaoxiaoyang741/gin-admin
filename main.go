package main

import (
	"github.com/zhaoxiaoyang741/gin-admin/global"
	"github.com/zhaoxiaoyang741/gin-admin/initialize"
)

func main() {
	// 初始化读取后台配置
	global.BaseSystem = initialize.InitConfig()
	initialize.InitServer()
}
