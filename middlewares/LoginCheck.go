package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

var LoginCheckFilter = func(ctx *context.Context) {
	_, exists := ctx.Input.Session("login_user_id").(int)
	if exists == false {
		ctx.Redirect(302, beego.URLFor("UserController.Login"))
		return
	}
}

func LoginCheck()  {
	beego.InsertFilter("/member/*", beego.BeforeRouter, LoginCheckFilter)
}