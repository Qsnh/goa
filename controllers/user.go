package controllers

import (
    "github.com/astaxie/beego"
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
    this.Layout = "layout/app.tpl"
}

// @router /register [post]
func (this *UserController) RegisterHandler() {

}