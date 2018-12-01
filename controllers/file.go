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
	//f.MustParams("userCode")

	m := make(map[string]interface{})

	for k, v := range f.Params {
		m[k] = v
	}

	f.JsonSucc("success", m)
}
