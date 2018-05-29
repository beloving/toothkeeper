package help

import "github.com/astaxie/beego"

func ReturnSucc(con * beego.Controller)  {
	con.Data["json"] = map[string]string{"status": "succ"}
	con.ServeJSON()
}
func ReturnFalse(con * beego.Controller,err error)  {
	con.Data["json"] = map[string]string{"error": err.Error()}
	con.ServeJSON()
}
func ReturnErrorMsg(con * beego.Controller,msg string)  {
	con.Data["json"] = map[string]string{"error": msg}
	con.ServeJSON()
}