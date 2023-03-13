package userController

import (
	"github.com/gin-gonic/gin"
	"image-host/app/apiException"
	"image-host/app/services/sessionServices"
	"image-host/app/services/userServices"
	"image-host/app/utils"
	"image-host/config/wechat"
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

func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)

	if err != nil {
		_ = c.AbortWithError(200, apiException.OpenIDError)
		return
	}
	postForm.Username = strings.ToUpper(postForm.Username)
	user, err := userServices.CreateStudentUserWechat(
		postForm.Username,
		postForm.Password,
		session.OpenID)
	if err != nil && err != apiException.ReactiveError {
		_ = c.AbortWithError(200, err)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func CreateStudentUser(c *gin.Context) {
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

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
