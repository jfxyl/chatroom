package db

import (
	"chatroom/app/models"
	"chatroom/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	G_DB *gorm.DB
)

func InitDatabase() (err error) {
	var (
		dsn string
	)
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.G_Config.Mysql.Username,
		config.G_Config.Mysql.Password,
		config.G_Config.Mysql.Host,
		config.G_Config.Mysql.Port,
		config.G_Config.Mysql.Database,
	)
	G_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
	if err != nil {
		return err
	}
	err = G_DB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.UserRoom{},
		&models.Message{},
		&models.MessageRead{},
	)
	return err
}
