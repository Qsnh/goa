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
	if this.CurrentLoginUser != nil {
		this.RedirectTo("/")
	}
    this.Layout = "layout/app.tpl"
}

// @router /login [post]
func (this *UserController) LoginHandler() {
	this.redirectUrl = beego.URLFor("UserController.Login")
	loginData := validations.UserLoginValidation{}
	this.ValidatorAuto(&loginData)

	user, err := models.UserExistsByEmailAndPassword(loginData.Email, loginData.Password)
	if err != nil {
		this.FlashError("用户名或密码错误")
		this.RedirectTo(this.redirectUrl)
	}

	this.SetSession("login_user_id", user.Id)

	this.FlashSuccess("登陆成功")
	this.RedirectTo("/")
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