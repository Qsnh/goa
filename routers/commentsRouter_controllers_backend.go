package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

}
