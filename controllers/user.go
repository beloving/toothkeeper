package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"toothkeeper/models"
	"toothkeeper/help"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	false		"注册手机号和密码"
// @Success 200 succ
// @Failure 403 fail
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	_,error := models.AddUser(user)
	if error !=nil{
		u.Data["json"] = map[string]string{"status": "fail"}
		u.ServeJSON()
	}else {
		help.ReturnSucc(&u.Controller)
	}
}

// @Title Login
// @Description Logs user into the system
// @Param	namewithpwd		body 	models.User	false		"手机号和密码登录"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if models.Login(user) {
		help.ReturnSucc(&u.Controller)
	} else {
		u.Data["json"] = "user not exist"
		u.ServeJSON()
	}
}

