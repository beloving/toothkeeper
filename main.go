package main

import (
	_ "toothkeeper/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"bytes"
	"toothkeeper/models"
	_"github.com/lib/pq"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	var dbtype = beego.AppConfig.String("dbtype")
	var username = beego.AppConfig.String("dbuser")
	var password = beego.AppConfig.String("dbpassword")
	var dbname=beego.AppConfig.String("dbname")
	orm.RegisterDriver(dbtype, orm.DRPostgres)
	var bf bytes.Buffer
	bf.WriteString("user=")
	bf.WriteString(username)
	bf.WriteString(" password=")
	bf.WriteString(password)
	bf.WriteString(" dbname=")
	bf.WriteString(dbname)
	bf.WriteString(" host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterDataBase("default", dbtype,bf.String(), 30)
	// register model
	orm.RegisterModel(new(models.User))


	orm.RunSyncdb("default", false, true)
	beego.Run()
}
