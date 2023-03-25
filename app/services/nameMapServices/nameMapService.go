package nameMapServices

import (
	"gorm.io/gorm"
	"image-host/app/models"
	"image-host/config/database"
)

func Insert(nameMap models.NameMap) error {
	res := database.DB.Create(&nameMap)
	return res.Error
}

func QueryByUUID(uuid string) (models.NameMap, error) {
	nameMap := models.NameMap{}
	result := database.DB.Where(
		&models.NameMap{
			UUID: uuid,
		},
	).First(&nameMap)
	if result.Error != nil {
		return models.NameMap{}, result.Error
	}
	return nameMap, nil
}

func QueryBySrc(src string) (models.NameMap, error) {
	nameMap := models.NameMap{}
	result := database.DB.Where(
		&models.NameMap{
			Src: src,
		},
	).First(&nameMap)
	if result.Error != nil {
		return models.NameMap{}, result.Error
	}
	return nameMap, nil
}

func Delete(src string) error {
	result := database.DB.Delete(models.NameMap{
		Src: src,
	})
	return result.Error
}

func FileDownloadCountIncrement(src string) error {
	return database.DB.Model(models.NameMap{}).Where(models.NameMap{
		Src: src,
	}).Update("download_count", gorm.Expr("download_count + ?", 1)).Error
}
