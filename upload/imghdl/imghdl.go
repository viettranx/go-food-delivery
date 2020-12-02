package imghdl

import (
	"fmt"
	"fooddlv/common"
	"fooddlv/upload/imgrepo"
	"fooddlv/upload/imgstorage"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

var supportImageExtNames = []string{".jpg", ".jpeg", ".png", ".ico", ".svg", ".bmp", ".gif"}

func isImage(extName string) bool {
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == extName {
			return true
		}
	}
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func UploadImg(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			form     *multipart.Form
			err      error
			distPath string
		)

		form, err = c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		files := form.File["files"]
		var imgs = make([]common.Image, len(files))
		distPath, err = os.Getwd()
		for index, file := range files {
			extname := strings.ToLower(path.Ext(file.Filename))

			if ok := isImage(extname); !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "unsuport file type",
				})
				return
			}

			saveDir := fmt.Sprintf("%s/public/%s", distPath, strings.TrimSpace(file.Filename))
			if fileExists(saveDir) {
				continue
			}

			if err := c.SaveUploadedFile(file, saveDir); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
			var img = common.Image{
				Url:    fmt.Sprintf("%s/%s", "http://localhost:8080/v1/file", strings.TrimSpace(file.Filename)),
				Width:  100,
				Height: 100,
			}
			imgs[index] = img
		}

		db := appCtx.GetDBConnection()
		store := imgstorage.NewImgSqlStorage(db)
		repo := imgrepo.NewCreateImgRepo(store)

		err = repo.Create(c.Request.Context(), imgs)

		if err != nil {

		}

		var ids []int
		for _, img := range imgs {
			ids = append(ids, img.Id)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]interface{}{"ids": ids}))
		return
	}
}
