package fileControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image-host/app/apiException"
	"image-host/app/config"
	"image-host/app/models"
	"image-host/app/services/nameMapServices"
	"image-host/app/utils"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

// UploadImg 上传图片
func UploadImg(c *gin.Context) {
	// 存储文件
	isTransparent := c.Query("isTransparent")
	form, _ := c.MultipartForm()
	img := form.File["img"][0]
	imgName := img.Filename
	_ = c.SaveUploadedFile(img, "./tmp/"+imgName)

	// 打开并判断文件类型
	file, _ := os.Open("./tmp/" + imgName)
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	_ = file.Close()
	if contentType != "image/png" && contentType != "image/jpeg" {
		_ = os.Remove("./tmp/" + imgName)
		_ = c.AbortWithError(200, apiException.ImgTypeError)
		return
	}

	// 重启文件并转换类型
	file, _ = os.Open("./tmp/" + imgName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	imgPrefix := strings.TrimSuffix(img.Filename, path.Ext(imgName))
	if contentType == "image/png" && isTransparent == "false" {
		// 为了处理一些仅修改了后缀而并未重新编码的图片，所有 png 文件都改为正确后缀
		newTypeName := "./tmp/" + imgPrefix + ".png"
		_ = os.Rename("./tmp/"+imgName, newTypeName)

		// png2jpg
		imgNew, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		out, err := os.Create("./tmp/" + imgPrefix + ".jpg")
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		defer func(out *os.File) {
			_ = out.Close()
		}(out)
		err = jpeg.Encode(out, imgNew, &jpeg.Options{Quality: 95})
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		_ = os.Remove(newTypeName)
		imgName = imgPrefix + ".jpg"
		_ = file.Close()
		file, _ = os.Open("./tmp/" + imgName)
	}
	fileName := uuid.NewString()

	var imgType string
	if isTransparent == "true" {
		// 处理透明图片
		if contentType == "image/jpeg" {
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		fileName += ".png"
		imgPrefix += ".png"
		imgType = "image/png"
		_ = os.Rename("./tmp/"+imgName, "./img/"+fileName)
	} else {
		// jpg2webp
		fileName += ".webp"
		imgPrefix += ".webp"
		imgType = "image/webp"
		imgNew, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		output, _ := os.Create("./img/" + fileName)
		defer func(output *os.File) {
			_ = output.Close()
		}(output)
		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		err = webp.Encode(output, imgNew, options)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		_ = os.Remove("./tmp/" + imgName)
	}

	fileInfo, err := os.Stat("./img/" + fileName)
	err = nameMapServices.Insert(models.NameMap{
		Src:         imgPrefix,
		UUID:        fileName,
		Type:        imgType,
		Size:        fileInfo.Size(),
		Temporary:   false,
		ExpireCount: -1,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, config.GetWebpUrlKey()+fileName)
}

// GetImgList 获取图片列表
func GetImgList(c *gin.Context) {
	pageStr, _ := url.QueryUnescape(c.Param("page"))
	page, _ := strconv.Atoi(pageStr)
	list, count, err := nameMapServices.QueryImgList(page)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"list":  list,
		"count": count,
	})
}
