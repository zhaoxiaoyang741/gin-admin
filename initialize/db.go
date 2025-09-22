package initialize

import (
	"fmt"

	"github.com/zhaoxiaoyang741/gin-admin/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresql() *gorm.DB {
	dsnTemplate := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	dsn := fmt.Sprintf(dsnTemplate,
		global.Config.Db.Host,
		global.Config.Db.Hostname,
		global.Config.Db.Password,
		global.Config.Db.Database,
		global.Config.Db.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database %s", err.Error()))
	}
	sqlDb, err := db.DB()

	if err != nil {
		panic(fmt.Sprintf("failed to connect database %s", err.Error()))
	}

	sqlDb.SetMaxIdleConns(int(global.Config.Db.MaxIdleConns))
	sqlDb.SetMaxOpenConns(int(global.Config.Db.MaxIdleConns))
	fmt.Println("connect to postgresql success")
	return db
}
