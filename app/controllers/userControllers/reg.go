package userControllers

import (
	"github.com/gin-gonic/gin"
	"image-host/app/apiException"
	"image-host/app/services/userServices"
	"image-host/app/utils"
	"strconv"
	"strings"
)

type createStudentUserForm struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type createStudentUserWechatForm struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Code     string `json:"code"  binding:"required"`
}

func CreateUser(c *gin.Context) {
	var postForm createStudentUserForm
	errBind := c.ShouldBindJSON(&postForm)
	if errBind != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	postForm.Username = strings.ToUpper(postForm.Username)
	user, err := userServices.CreateStudentUser(
		postForm.Username,
		postForm.Password)
	if err != nil && err != apiException.ReactiveError {
		_ = c.AbortWithError(200, err)
		return
	}

	token := utils.GenerateStandardJwt(&utils.JwtData{
		ID: strconv.Itoa(user.ID),
	})
	utils.JsonSuccessResponse(c, gin.H{"token": token})
}
