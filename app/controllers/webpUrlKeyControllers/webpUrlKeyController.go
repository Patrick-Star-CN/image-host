package webpUrlKeyControllers

import (
	"github.com/gin-gonic/gin"
	"image-host/app/utils"
	"image-host/config/config"
)

// SubmitKey 上传路径 key
func SubmitKey(c *gin.Context) {
	key := c.Query("key")
	config.Config.Set("webpUrlKey", key)
	utils.JsonSuccessResponse(c, nil)
}

// GetKey 获取路径 key
func GetKey(c *gin.Context) {
	key := config.Config.GetString("webpUrlKey")
	utils.JsonSuccessResponse(c, key)
}

// DeleteKey 删除路径 key
func DeleteKey(c *gin.Context) {
	config.Config.Set("webpUrlKey", "")
	utils.JsonSuccessResponse(c, nil)
}
