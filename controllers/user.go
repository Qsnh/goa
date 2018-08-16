package controllers

import (
    "github.com/astaxie/beego/validation"
    "goa/validations"
)

type UserController struct {
    Base
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
    this.Layout = "layout/app.tpl"
}

// @router /register [post]
func (this *UserController) RegisterHandler() {
    userData := validations.UserRegisterValidation{}
    if err := this.ParseForm(&userData); err != nil {
        this.FlashBag.Error("参数解析错误")
        this.FlashBag.Store(&this.Controller)
        this.Redirect("/register", 302)
        return
    }

    validate := validation.Validation{}
    isValid, err := validate.Valid(&userData)
    if err != nil {
        this.FlashError("服务器出错")
        this.Redirect("/register", 302)
        return
    }

    if !isValid {
        for _, err := range validate.Errors {
            this.FlashError(err.Key + err.Message)
            break
        }
        this.Redirect("/register", 302)
        return
    }

    this.FlashSuccess("注册成功")
    this.Redirect("/login", 302)
}