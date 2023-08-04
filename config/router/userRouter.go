package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/userControllers"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create/student/wechat", userControllers.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userControllers.CreateStudentUser)

		user.POST("/login/wechat", userControllers.WeChatLogin)
		user.POST("/login", userControllers.AuthByPassword)
		user.POST("/login/session", userControllers.AuthBySession)
	}
}
