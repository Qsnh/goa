package routers

import (
	"github.com/Qsnh/goa/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
    beego.Include(&controllers.UserController{})
	beego.Include(&controllers.MemberController{})
	beego.Include(&controllers.QuestionController{})
	beego.Include(&controllers.UploadController{})
}