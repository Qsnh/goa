package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("go", "mysql", "root:root@tcp(127.0.0.1:33060)/?charset=utf8")
}

func main() {
	beego.Run()
}