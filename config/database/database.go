package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"image-host/config/config"
	"log"
)

var DB *gorm.DB

func Init() { // 初始化数据库
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.UserName,
		config.Config.DB.Password,
		config.Config.DB.Address,
		config.Config.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})

	if err != nil {
		log.Fatal("DatabaseConnectFailed", err)
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed", err)
	}
	DB = db
}
