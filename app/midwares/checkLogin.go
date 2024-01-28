package midwares

import (
	"github.com/gin-gonic/gin"
	"image-host/app/apiException"
	"image-host/app/utils"
	"log"
	"strconv"
)

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		_ = c.AbortWithError(200, apiException.NotLogin)
		c.Abort()
		return
	}
	id, err := utils.ParseToken(token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		c.Abort()
		return
	}
	idInt, ok := strconv.Atoi(id)
	if ok != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		log.Println(ok)
		c.Abort()
	}
	c.Set("user_id", idInt)
	c.Next()
}
