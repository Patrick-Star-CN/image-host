package userControllers

import (
	"crypto/sha256"
	"image-host/app/apiException"
	"image-host/app/services/userServices"
	"image-host/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type autoLoginForm struct {
	Code      string `json:"code" binding:"required"`
	LoginType string `json:"type"`
}
type passwordLoginForm struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	LoginType string `json:"type"`
}

func AuthByPassword(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)

	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := userServices.GetUserByUsername(postForm.Username)
	if err == gorm.ErrRecordNotFound {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	h := sha256.New()
	h.Write([]byte(postForm.Password))

	token := utils.GenerateStandardJwt(&utils.JwtData{
		ID: strconv.Itoa(user.ID),
	})
	utils.JsonSuccessResponse(c, gin.H{"token": token})
}
