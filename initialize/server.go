package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhaoxiaoyang741/gin-admin/global"
)

// 127.0.0.1:8088 只允许本地访问（localhost）其他计算机无法访问该服务 适用于开发环境或只供本地使用的服务 最安全的方式，因为限制了外部访问
// 0.0.0.0:8088 允许任何网络接口访问 其他计算机可以通过局域网IP或公网IP访问 适用于生产环境或需要被其他设备访问的服务 安全性较低，因为允许所有网络接口访问
// :8088 效果等同于 0.0.0.0:8088 是 0.0.0.0:8088 的简写形式 同样允许所有网络接口访问

func InitServer() {
	c := gin.Default()
	url := fmt.Sprintf("127.0.0.1:%s", global.BaseSystem.System.Port)
	if gin.Mode() == gin.ReleaseMode {
		url = fmt.Sprintf("0.0.0.0:%s", global.BaseSystem.System.Port)
	}
	c.Run(url)
}
