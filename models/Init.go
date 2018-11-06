package models

import (
	"github.com/astaxie/beego/orm"
	"net/url"
	"os"
)

func Init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dbTimeZone := os.Getenv("DATABASE_TIMEZONE")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4"
	if dbTimeZone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbTimeZone)
	}

	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Categories), new(Users), new(Answers), new(Questions))

	if os.Getenv("APP_DEBUG") == "dev" {
		orm.Debug = true
	}
}