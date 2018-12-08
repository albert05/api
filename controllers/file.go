package controllers

import (
	"api/services"
	"api/common"
	"path/filepath"
	"strings"
	"github.com/astaxie/beego"
)

// Operations about File
type FileController struct {
	BaseController
}

// @Title Upload
// @Description File Upload
// @Success 200 {int}
// @Failure 403
// @router /upload [post]
func (this *FileController) Upload() {
	f, h, err := this.GetFile("image")
	if err != nil {
		this.Abort666("upload failed", map[string]interface{}{
			"err": err.Error(),
		})
	}
	defer f.Close()

	path := services.GenerateFilePath(this.UserKey)
	if !common.IsExist(path) {
		if err := common.MakeDir(path); err != nil {
			this.Abort666("failed", map[string]interface{}{
				"err": err.Error(),
			})
		}
	}

	ext := filepath.Ext(filepath.Base(h.Filename))
	fileName := path + services.GenerateFileName() + ext
	this.SaveToFile("image", fileName)

	this.JsonSucc("success", map[string]interface{}{
		"path": filepath.Base(h.Filename),
		"filename": h.Filename,
		"url": strings.Replace(fileName, beego.AppConfig.String("fileRootPath"), beego.AppConfig.String("fileHost"), 1),
	})
}
