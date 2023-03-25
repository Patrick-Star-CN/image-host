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
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func UploadImg(c *gin.Context) {
	// 存储文件
	form, _ := c.MultipartForm()
	img := form.File["img"][0]
	imgName := img.Filename
	_ = c.SaveUploadedFile(img, "./tmp/"+imgName)

	// 打开并判断文件类型
	file, _ := os.Open("./tmp/" + imgName)
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	file.Close()

	// 重启文件并转换类型
	file, _ = os.Open("./tmp/" + imgName)
	defer file.Close()
	imgPrefix := strings.TrimSuffix(img.Filename, path.Ext(imgName))
	if contentType == "image/png" {
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
		defer out.Close()
		err = jpeg.Encode(out, imgNew, &jpeg.Options{Quality: 95})
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		_ = os.Remove(newTypeName)
		imgName = imgPrefix + ".jpg"
		file.Close()
		file, _ = os.Open("./tmp/" + imgName)
	}

	// jpg2webp
	imgNew, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(200, apiException.ImgTypeError)
		return
	}
	fileName := uuid.NewString() + ".webp"
	output, _ := os.Create("./img/" + fileName)
	defer output.Close()
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
	fileInfo, err := os.Stat("./img/" + fileName)
	err = nameMapServices.Insert(models.NameMap{
		Src:  imgPrefix + ".webp",
		UUID: fileName,
		Type: "image/webp",
		Path: "./img/" + fileName,
		Size: fileInfo.Size(),
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, config.GetWebpUrlKey()+fileName)
}

func GetFile(c *gin.Context) {
	uuidName, _ := url.QueryUnescape(c.Param("file_name"))
	file, err := nameMapServices.QueryByUUID(uuidName)
	if (err != nil || file == models.NameMap{}) {
		_ = c.AbortWithError(404, apiException.NotFound)
		return
	}
	c.FileAttachment(file.Path, file.Src)
	if err := nameMapServices.FileDownloadCountIncrement(file.Src); err != nil {
		log.Println(err)
	}
	return
}

func UploadFile(c *gin.Context) {
	// 存储文件
	file, _ := c.FormFile("file")
	fileName := file.Filename
	uuidName := uuid.NewString() + path.Ext(fileName)
	_ = c.SaveUploadedFile(file, "./public/"+uuidName)
	err := nameMapServices.Insert(models.NameMap{
		Src:  fileName,
		UUID: uuidName,
		Type: file.Header.Get("Content-Type"),
		Path: "./public/" + uuidName,
		Size: file.Size,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
	}
	utils.JsonSuccessResponse(c, uuidName)
}
