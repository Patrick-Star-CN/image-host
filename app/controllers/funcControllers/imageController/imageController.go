package imageController

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image-host/app/apiException"
	"image-host/app/config"
	"image-host/app/utils"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path"
	"strings"
)

func UploadImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	img := form.File["img"][0]
	imgName := img.Filename
	_ = c.SaveUploadedFile(img, "./tmp/"+imgName)
	file, _ := os.Open("./tmp/" + imgName)
	defer file.Close()
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	if contentType == "image/png" {
		imgNew, _ := png.Decode(file)
		out, _ := os.Create("./tmp/" + strings.TrimSuffix(img.Filename, path.Ext(path.Base(img.Filename))) + ".jpg")
		defer out.Close()
		_ = jpeg.Encode(out, imgNew, &jpeg.Options{Quality: 95})
		_ = os.Remove("./tmp/" + img.Filename)
		imgName = strings.TrimSuffix(img.Filename, path.Ext(path.Base(img.Filename))) + ".jpg"
		file.Close()
		file, _ = os.Open("./tmp/" + imgName)
	}
	imgNew, _ := jpeg.Decode(file)
	fileName := uuid.NewString() + ".webp"
	output, _ := os.Create("./img/" + fileName)
	defer output.Close()
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	err = webp.Encode(output, imgNew, options)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	_ = os.Remove("./tmp/" + imgName)
	utils.JsonSuccessResponse(c, config.GetWebpUrlKey()+fileName)
}
