package webpUrlKeyControllers

import (
	"github.com/gin-gonic/gin"
	"image-host/app/config"
	"image-host/app/utils"
)

// SubmitKey 上传路径 key
func SubmitKey(c *gin.Context) {
	key := c.Query("key")
	err := config.SetWebpUrlKey(key)
	if err != nil {
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// GetKey 获取路径 key
func GetKey(c *gin.Context) {
	key := config.GetWebpUrlKey()
	utils.JsonSuccessResponse(c, key)
}

// DeleteKey 删除路径 key
func DeleteKey(c *gin.Context) {
	err := config.DelWebpUrlKey()
	if err != nil {
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
