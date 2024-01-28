package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/userControllers"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create", userControllers.CreateUser)

		user.POST("/login", userControllers.AuthByPassword)
	}
}
