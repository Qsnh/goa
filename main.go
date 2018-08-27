package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "goa/routers"
	"goa/models"
	"goa/middlewares"
)

func main() {
	models.Init()
	middlewares.LoginCheck()
	beego.Run()
}