
package routers

import (
	"github.com/astaxie/beego"
	"toothkeeper/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/tooth",
				beego.NSInclude(&controllers.ToothController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
