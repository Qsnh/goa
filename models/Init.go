package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"net/url"
)

func Init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbHost := beego.AppConfig.String("database_host")
	dbPort := beego.AppConfig.String("database_port")
	dbUser := beego.AppConfig.String("database_user")
	dbPass := beego.AppConfig.String("database_pass")
	dbName := beego.AppConfig.String("database_name")
	dbTimeZone := beego.AppConfig.String("database_timezone")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4"
	if dbTimeZone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbTimeZone)
	}

	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Categories), new(Users), new(Answers), new(Questions))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}