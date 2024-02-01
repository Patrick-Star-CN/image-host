package router

import (
	"github.com/gin-gonic/gin"
	"image-host/app/controllers/fileControllers"
	"image-host/app/controllers/webpUrlKeyControllers"
	"image-host/app/midwares"
)

func fileRouterInit(r *gin.RouterGroup) {
	fileFun := r.Group("/file", midwares.CheckLogin)
	{
		fileFun.POST("/upload", fileControllers.UploadFile)

		fileFun.GET("/list/:page", fileControllers.GetFileList)

		fileFun.DELETE("/:file_name", fileControllers.DeleteFile)

	}

	r.GET("/file/download/:file_name", fileControllers.GetFile)

	imgFun := r.Group("/img", midwares.CheckLogin)
	{
		imgFun.POST("/upload", fileControllers.UploadImg)

		imgFun.GET("/download/:file_name", fileControllers.GetFile)

		imgFun.GET("/list/:page", fileControllers.GetImgList)

		imgFun.DELETE("/:file_name", fileControllers.DeleteFile)
	}

	webpUrlKeyFun := r.Group("/key", midwares.CheckLogin)
	{
		webpUrlKeyFun.POST("", webpUrlKeyControllers.SubmitKey)

		webpUrlKeyFun.GET("", webpUrlKeyControllers.GetKey)

		webpUrlKeyFun.DELETE("", webpUrlKeyControllers.DeleteKey)
	}
}
