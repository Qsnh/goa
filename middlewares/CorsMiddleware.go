package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func CorsHandler()  {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "HEAD", "PUT"},
		AllowHeaders:     []string{"token", "Content-Type"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
	}))
}