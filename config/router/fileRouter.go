package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/fileControllers"
	"image-host/app/midwares"
)

func fileRouterInit(r *gin.RouterGroup) {
	fun := r.Group("/file", midwares.CheckLogin)
	{
		fun.POST("/upload_img", fileControllers.UploadImg)
		fun.GET("/download/:file_name", fileControllers.GetFile)
		fun.POST("/upload", fileControllers.UploadFile)
	}
}
