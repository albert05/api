package controllers

import (
	"github.com/astaxie/beego"
	"api/services"
	"fmt"
	"api/common"
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
	f, _, _ := this.GetFile("image")
	defer f.Close()

	path := fmt.Sprintf("%s/%s/", beego.AppConfig.String("fileRootPath"), this.UserKey)

	if !common.IsExist(path) {
		if err := common.MakeDir(path); err != nil {
			this.Abort666("failed", map[string]interface{}{
				"err": err.Error(),
			})
		}
	}

	fileName := services.GenerateFileName()
	this.SaveToFile("image", path + fileName)

	this.JsonSucc("success")
}
