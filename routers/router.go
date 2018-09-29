package routers

import (
	"github.com/Qsnh/goa/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}