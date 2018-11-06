package main

import (
	"github.com/Qsnh/goa/middlewares"
	"github.com/Qsnh/goa/models"
	_ "github.com/Qsnh/goa/routers"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't find .env file")
		return
	}
	models.Init()
	middlewares.LoginCheck()
	beego.AddFuncMap("urlquery", utils.Url)
	beego.Run(":"+os.Getenv("APP_PORT"))
	logs.SetLogger("console")
}
