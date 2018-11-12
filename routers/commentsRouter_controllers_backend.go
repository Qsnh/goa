package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:AnswerContoller"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:AnswerContoller"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/backend/answer/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:AnswerContoller"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:AnswerContoller"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/backend/answers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/backend/categories`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/backend/category`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/backend/category/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/backend/category/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/backend/category/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:DashboardController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:DashboardController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/backend/dashboard`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:LoginController"],
		beego.ControllerComments{
			Method: "LoginHandler",
			Router: `/backend/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:QuestionController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/backend/question/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:QuestionController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/backend/questions`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/backend/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/backend/user/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/backend/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
