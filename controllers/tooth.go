package controllers

import (
	"github.com/astaxie/beego"
	"toothkeeper/models"
	"encoding/json"
	"toothkeeper/help"
)

type ToothController struct {
	beego.Controller
}
// @Title CreateTooth
// @Description create tooths
// @Param	body		body 	models.Tooth	false		"添加一个牙齿模型数据"
// @Success 200 succ
// @Failure 403 fail
// @router / [post]
func (u *ToothController) Post() {
	var tooth models.Tooth
	json.Unmarshal(u.Ctx.Input.RequestBody, &tooth)

	_,error := models.CreateTooth(tooth)
	if error !=nil{
		help.ReturnFalse(&u.Controller,error)
	}else {
		help.ReturnSucc(&u.Controller)
	}
}

// @Title CreateTooth
// @Description create tooths
//@Param  id  path  int64 0 "牙齿id"
// @Success 200 succ
// @Failure 403 fail
// @router /:id [get]
func (u *ToothController) Get() {
	var tooth models.Tooth
	id,err := u.GetInt64(":id")
	if err == nil{
		tooth.Id = id
		toothData,error := models.FindOneTooth(&tooth)
		if error !=nil{
			help.ReturnErrorMsg(&u.Controller,"没有此数据")
		}else {
			u.Data["json"] = toothData
			u.ServeJSON()
		}
	}else{
		help.ReturnFalse(&u.Controller,err)
	}

}