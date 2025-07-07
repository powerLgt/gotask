package database

import (
	"blog/config"
	"blog/module/common/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error

	// 配置GORM日志
	gormConfig := &gorm.Config{}
	gormConfig.Logger = logger.Default.LogMode(logger.Info)

	// 连接数据库
	switch config.AppConfig.Database.Driver {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(config.AppConfig.Database.DSN), gormConfig)
	default:
		log.Fatal("不支持的数据库驱动")
	}

	if err != nil {
		return err
	}

	// 自动迁移数据库表
	err = DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		return err
	}

	log.Println("数据库连接成功并完成表迁移")
	return nil
}
