package routers

import (
	"goa/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
    beego.Include(&controllers.UserController{})
	beego.Include(&controllers.MemberController{})
}