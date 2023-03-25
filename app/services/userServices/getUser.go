package userServices

import (
	"image-host/app/models"
	"image-host/config/database"
)

func GetUserByWechatOpenID(openid string) *models.User {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			WechatOpenID: openid,
		},
	).First(&user)
	if result.Error != nil {
		return nil
	}

	return &user
}

func GetUserID(id int) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			Username: username,
		},
	).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
