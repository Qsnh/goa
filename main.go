package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
    "github.com/Qsnh/goa/models"
    _ "github.com/Qsnh/goa/routers"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    dbHost := beego.AppConfig.String("database_host")
    dbPort := beego.AppConfig.String("database_port")
    dbUser := beego.AppConfig.String("database_user")
    dbPass := beego.AppConfig.String("database_pass")
    dbDb := beego.AppConfig.String("database_db")

    orm.RegisterDataBase(dbDb, "mysql", dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8")
    orm.RegisterModel(new(models.Category), new(models.User), new(models.Question))
}

func main() {
	beego.Run()
}