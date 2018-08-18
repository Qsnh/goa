package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "goa/routers"
	"goa/models"
)

func main() {
	models.Init()
	beego.Run()
}