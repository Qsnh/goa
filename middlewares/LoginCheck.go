package middlewares

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

var LoginCheckFilter = func(ctx *context.Context) {
	userId := ctx.GetCookie("login_user_id")
	userSign := ctx.GetCookie("login_user_sign")
	user := models.Users{}
	if err := orm.NewOrm().QueryTable("users").Filter("id", userId).One(&user); err != nil {
		ctx.Redirect(302, beego.URLFor("UserController.Login"))
		return
	}
	if utils.AuthSignCheck(user.Id, user.Email, user.Password, userSign) == false {
		ctx.Redirect(302, beego.URLFor("UserController.Login"))
		return
	}
}

func LoginCheck() {
	beego.InsertFilter("/member/*", beego.BeforeRouter, LoginCheckFilter)
}
