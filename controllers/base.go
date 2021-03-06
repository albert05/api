package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"api/common"
	"api/models"
)

// Operations about Users
type BaseController struct {
	beego.Controller
	Params	map[string]string
	UserKey string
}

func (ctx *BaseController) Prepare() {
	json.Unmarshal(ctx.Ctx.Input.RequestBody, &ctx.Params)

	ctx.MustAuth()
}

func (ctx *BaseController) MustParams(keys ...string) {
	if !common.CheckParams(ctx.Params, keys) {
		ctx.Abort666("failed", map[string]interface{}{
			"err": "invalid params",
		})
	}
}

func (ctx *BaseController) MustAuth() {
	key := ctx.GetString("api_key")
	secret := ctx.GetString("api_secret")

	v, ok := models.KeyList[key]
	if !ok || v != secret {
		ctx.Abort666("failed", map[string]interface{}{
			"err": "invalid params",
		})
	}

	ctx.UserKey = key
}

func (ctx *BaseController) JsonSucc(msg string, datas ...map[string]interface{}) {
	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx.Data["json"] = &Code{
		Code:   SUCCESS_CODE,
		Message:    msg,
		Data: data,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) Abort666(msg string, datas ...map[string]interface{}) {
	ctx.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}

	ctx.CustomAbort(200, map2Json(&Code{
		Code:   FAILED_CODE,
		Message:    msg,
		Data: data,
	}))
}

func map2Json(data interface{}) string {
	content, _ := json.Marshal(data)

	return string(content)
}

const SUCCESS_CODE  = 200
const FAILED_CODE  = 100

type Code struct {
	Code int						`json:"code"`
	Message string 					`json:"message"`
	Data  	map[string]interface{} 	`json:"data"`
}



