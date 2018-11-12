package routers

import (
	"github.com/Qsnh/goa/controllers/backend"
	"github.com/Qsnh/goa/controllers/frontend"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&frontend.IndexController{})
    beego.Include(&frontend.UserController{})
	beego.Include(&frontend.MemberController{})
	beego.Include(&frontend.QuestionController{})
	beego.Include(&frontend.UploadController{})
	beego.Include(&frontend.DashboardController{})

	beego.Include(&backend.CategoryController{})
	beego.Include(&backend.QuestionController{})
	beego.Include(&backend.UserController{})
	beego.Include(&backend.AnswerContoller{})
	beego.Include(&backend.LoginController{})
	beego.Include(&backend.DashboardController{})
}