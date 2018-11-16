package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"os"
)

func CorsHandler() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGINAL")},
		AllowMethods:     []string{"GET", "POST", "DELETE", "HEAD", "PUT"},
		AllowHeaders:     []string{"token", "Content-Type"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
	}))
}
