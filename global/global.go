package global

import (
	"github.com/zhaoxiaoyang741/gin-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Logger *zap.Logger
	Db     *gorm.DB
)
