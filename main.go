package main

import (
	"github.com/Qsnh/goa/libs"
	"github.com/Qsnh/goa/middlewares"
	"github.com/Qsnh/goa/models"
	_ "github.com/Qsnh/goa/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.Init()
	middlewares.LoginCheck()
	beego.AddFuncMap("urlquery", libs.Url)
	beego.Run()
	logs.SetLogger("console")
}
