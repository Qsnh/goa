package routers

import (
	"goa/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}