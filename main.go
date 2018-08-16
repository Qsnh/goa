package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
    "goa/models"
    _ "goa/routers"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    db_host := beego.AppConfig.String("database_host")
    db_port := beego.AppConfig.String("database_port")
    db_user := beego.AppConfig.String("database_user")
    db_pass := beego.AppConfig.String("database_pass")
    db_db := beego.AppConfig.String("database_db")

    orm.RegisterDataBase(db_db, "mysql", db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/?charset=utf8")
    orm.RegisterModel(new(models.Category), new(models.User), new(models.Question))
}

func main() {
	beego.Run()
}