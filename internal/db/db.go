package db

import (
	"chatroom/app/models"
	"chatroom/internal/config"
	"chatroom/internal/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	err = global.DB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.UserRoom{},
		&models.Message{},
		&models.MessageRead{},
	)
	return err
}
