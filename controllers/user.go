package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
    "github.com/Qsnh/goa/validations"
)

type UserController struct {
    beego.Controller
}

// @router /login [get]
func (this *UserController) Login() {
    this.Layout = "layout/app.tpl"
}

// @router /login [post]
func (this *UserController) LoginHandler() {

}

// @router /logout [get]
func (this *UserController) Logout() {
    
}

// @router /register [get]
func (this *UserController) Register() {
    beego.ReadFromRequest(&this.Controller)
    this.Layout = "layout/app.tpl"
}

// @router /register [post]
func (this *UserController) RegisterHandler() {
    flash := beego.NewFlash()
    
    userData := validations.UserRegisterValidation{}
    if err := this.ParseForm(&userData); err != nil {
        flash.Error("参数解析错误")
        flash.Store(&this.Controller)
        this.Redirect("/register", 302)
        return
    }

    validate := validation.Validation{}
    isValid, err := validate.Valid(&userData)
    if err != nil {
        flash.Error("服务器出错")
        flash.Store(&this.Controller)
        this.Redirect("/register", 302)
        return
    }

    if !isValid {
        for _, err := range validate.Errors {
            flash.Error(err.Key + err.Message)
            break
        }
        flash.Store(&this.Controller)
        this.Redirect("/register", 302)
        return
    }

    flash.Notice("注册成功")
    flash.Store(&this.Controller)
    this.Redirect("/login", 302)
}