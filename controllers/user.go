package controllers

import (
    "goa/validations"
    "github.com/astaxie/beego"
	"goa/models"
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
    this.redirectUrl = beego.URLFor("UserController.Register")
    userData := validations.UserRegisterValidation{}
    this.ValidatorAuto(&userData)

    _, err := models.CreateUser(userData.Nickname, userData.Email, userData.Password)
    if err != nil {
    	this.FlashError("注册失败")
    	this.RedirectTo(this.redirectUrl)
	}

    this.FlashSuccess("注册成功")
    this.RedirectTo(beego.URLFor("UserController.Login"))
}