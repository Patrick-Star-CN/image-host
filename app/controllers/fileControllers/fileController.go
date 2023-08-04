package fileControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image-host/app/apiException"
	"image-host/app/models"
	"image-host/app/services/nameMapServices"
	"image-host/app/utils"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

// GetFile 获取文件
func GetFile(c *gin.Context) {
	uuidName, _ := url.QueryUnescape(c.Param("file_name"))
	file, err := nameMapServices.QueryByUUID(uuidName)
	if (err != nil || file == models.NameMap{}) {
		_ = c.AbortWithError(404, apiException.NotFound)
		return
	}
	if strings.HasPrefix(file.Type, "image") {
		c.FileAttachment("./img/"+file.UUID, file.Src)
	} else {
		c.FileAttachment("./public/"+file.UUID, file.Src)
	}
	if file.Temporary && file.ExpireCount == file.DownloadCount+1 {
		_ = nameMapServices.Delete(file.UUID)
	} else if err := nameMapServices.FileDownloadCountIncrement(file.Src); err != nil {
		log.Println(err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	temp := c.Query("temp")
	expireCount, err := strconv.Atoi(c.Query("expireCount"))
	if err != nil {
		if c.Query("expireCount") == "" {
			expireCount = -1
		} else {
			_ = c.AbortWithError(200, apiException.ParamError)
			return
		}
	}
	file, _ := c.FormFile("file")
	fileName := file.Filename
	uuidName := uuid.NewString() + path.Ext(fileName)
	_ = c.SaveUploadedFile(file, "./public/"+uuidName)
	err = nameMapServices.Insert(models.NameMap{
		Src:         fileName,
		UUID:        uuidName,
		Type:        file.Header.Get("Content-Type"),
		Size:        file.Size,
		Temporary:   temp == "true",
		ExpireCount: expireCount,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, uuidName)
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	uuidName, _ := url.QueryUnescape(c.Param("file_name"))
	file, err := nameMapServices.QueryByUUID(uuidName)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 删除一个路径为file.uuid的文件
	if strings.HasPrefix(file.Type, "image") {
		_ = os.Remove("./img/" + file.UUID)
	} else {
		_ = os.Remove("./public/" + file.UUID)
	}
	err = nameMapServices.Delete(uuidName)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// GetFileList 获取文件列表
func GetFileList(c *gin.Context) {
	pageStr, _ := url.QueryUnescape(c.Param("page"))
	page, _ := strconv.Atoi(pageStr)
	list, count, err := nameMapServices.QueryFileList(page)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"list":  list,
		"count": count,
	})
}
