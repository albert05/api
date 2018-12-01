package controllers

// Operations about File
type FileController struct {
	BaseController
}

// @Title Upload
// @Description File Upload
// @Param userCode query string true
// @Success 200 {int}
// @Failure 403
// @router /upload [post]
func (f *FileController) Upload() {
	f.MustParams("userCode")

	f.JsonSucc("success", map[string]interface{}{
		"status": 1,
	})
}
