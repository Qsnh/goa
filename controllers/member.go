package controllers

import (
	"github.com/Qsnh/goa/libs"
	"github.com/Qsnh/goa/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberController struct {
	Base
}

// @router /member [get]
func (this *MemberController) Index() {
	this.Layout = "layout/member.tpl"
}

// @router /member/password [get]
func (this *MemberController) ChangePassword() {
	this.Layout = "layout/member.tpl"
}

// @router /member/password [post]
func (this *MemberController) ChangePasswordHandler() {
	this.redirectUrl = beego.URLFor("MemberController.ChangePassword")
	passwordData := validations.MemberChangePasswordValidation{}
	this.ValidatorAuto(&passwordData)

	if this.CurrentLoginUser.Password != libs.SHA256Encode(passwordData.OldPassword) {
		this.FlashError("原密码不正确")
		this.RedirectTo(this.redirectUrl)
	}

	this.CurrentLoginUser.Password = libs.SHA256Encode(passwordData.NewPassword)
	if result, err := orm.NewOrm().Update(this.CurrentLoginUser); err != nil || result == 0 {
		this.FlashError("修改失败")
		this.RedirectTo(this.redirectUrl)
	}

	this.FlashSuccess("密码修改成功")
	this.RedirectTo(this.redirectUrl)
}
