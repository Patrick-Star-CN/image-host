package database

import (
	"gorm.io/gorm"
	"image-host/app/models"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.NameMap{})
}
