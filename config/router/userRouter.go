package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/userController"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userController.CreateStudentUser)

		user.POST("/login/wechat", userController.WeChatLogin)
		user.POST("/login", userController.AuthByPassword)
		user.POST("/login/session", userController.AuthBySession)
	}
}
