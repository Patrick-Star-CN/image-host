package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"image-host/app/apiException"
	"image-host/app/models"
	"image-host/config/database"
	"time"
)

func CreateStudentUser(username, password string) (*models.User, error) {
	if CheckUsername(username) {
		return nil, apiException.UserAlreadyExisted
	}

	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))

	user := &models.User{
		JHPassword: pass,
		Username:   username,
		CreateTime: time.Now(),
	}

	res := database.DB.Create(&user)
	return user, res.Error
}

func CreateStudentUserWechat(username, password, wechatOpenID string) (*models.User, error) {
	if !CheckWechatOpenID(wechatOpenID) {
		return nil, apiException.OpenIDError
	}
	user, err := CreateStudentUser(username, password)
	if err != nil && err != apiException.ReactiveError {
		return nil, err
	}
	user.WechatOpenID = wechatOpenID
	database.DB.Updates(user)
	database.DB.Save(user)
	return user, nil
}
