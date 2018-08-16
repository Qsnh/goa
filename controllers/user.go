package controllers

import "goa/validations"

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
    this.redirectUrl = "/register"
    userData := validations.UserRegisterValidation{}
    this.ValidatorAuto(&userData)

    this.FlashSuccess("注册成功")
    this.Redirect("/login", this.redirectCode)
}