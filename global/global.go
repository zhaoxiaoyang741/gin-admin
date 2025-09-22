package global

import (
	"github.com/zhaoxiaoyang741/gin-admin/config"
	"go.uber.org/zap"
)

var (
	Config config.Config
	Logger *zap.Logger
)
