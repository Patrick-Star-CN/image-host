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

func QueryFileList(page int) ([]models.NameMap, int64, error) {
	var nameMap []models.NameMap
	var count int64
	countResult := database.DB.Model(&models.NameMap{}).Count(&count)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}
	database.DB.Model(&models.NameMap{}).
		Where("type LIKE ?", "application%").
		Offset((page - 1) * 10).Limit(10).Find(&nameMap)

	return nameMap, count, nil
}

func QueryImgList(page int) ([]models.NameMap, int64, error) {
	var nameMap []models.NameMap
	var count int64
	countResult := database.DB.Model(&models.NameMap{}).Count(&count)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}
	database.DB.Model(&models.NameMap{}).
		Where("type LIKE ?", "image%").
		Offset((page - 1) * 10).Limit(10).Find(&nameMap)

	return nameMap, count, nil
}

func Delete(uuid string) error {
	result := database.DB.Delete(models.NameMap{
		UUID: uuid,
	})
	return result.Error
}

func FileDownloadCountIncrement(src string) error {
	return database.DB.Model(models.NameMap{}).Where(models.NameMap{
		Src: src,
	}).Update("download_count", gorm.Expr("download_count + ?", 1)).Error
}
