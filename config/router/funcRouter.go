package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/funcControllers/imageController"
	"image-host/app/midwares"
)

func funcRouterInit(r *gin.RouterGroup) {
	fun := r.Group("/func", midwares.CheckLogin)
	{
		fun.POST("/upload_img", imageController.UploadImg)
	}
}
