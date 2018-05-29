package help

import "github.com/astaxie/beego"

func ReturnSucc(con * beego.Controller)  {
	con.Data["json"] = map[string]string{"status": "succ"}
	con.ServeJSON()
}
