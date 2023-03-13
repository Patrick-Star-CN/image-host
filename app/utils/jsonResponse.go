package utils

import (
	"github.com/gin-gonic/gin"
	"image-host/app/utils/stateCode"
	"net/http"
)

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": stateCode.OK,
		"msg":  "OK",
	})
}
