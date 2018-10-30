package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/Qsnh/goa/libs"
	"github.com/Qsnh/goa/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"path"
	"strings"
	"time"
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

// @router /member/avatar [get]
func (this *MemberController) ChangeAvatar() {
	this.Layout = "layout/member.tpl"
}

// @router /member/avatar [post]
func (this *MemberController) ChangeAvatarHandler() {
	this.redirectUrl = beego.URLFor("MemberController.ChangeAvatarHandler")
	file, header, err := this.GetFile("file")
	if err != nil {
		this.FlashError("请选择头像文件")
		this.RedirectTo(this.redirectUrl)
	}
	defer file.Close()
	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		this.FlashError("请上传有效图片文件")
		this.RedirectTo(this.redirectUrl)
	}
	// 文件后缀判断
	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		this.FlashError("请上传jpeg/png/gif图片")
		this.RedirectTo(this.redirectUrl);
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		this.FlashError("头像文件大小不能超过2MB")
		this.RedirectTo(this.redirectUrl)
	}
	// 保存文件
	rand.Seed(time.Now().Unix())
	newFilename := fmt.Sprintf("%d+%d", time.Now().Unix(), rand.Intn(100))
	newFilename = fmt.Sprintf("%x", md5.Sum([]byte(newFilename)))
	path := path.Join("static/uploads/avatar", newFilename+"."+extension)
	err = this.SaveToFile("file", path)
	if err != nil {
		logs.Info(err)
		this.FlashError("头像上传失败")
		this.RedirectTo(this.redirectUrl)
	}
	// 修改数据表
	this.CurrentLoginUser.Avatar = "/" + path
	if result, err := orm.NewOrm().Update(this.CurrentLoginUser); err != nil || result == 0 {
		this.FlashSuccess("头像保存失败")
		this.RedirectTo(this.redirectUrl)
	}
	this.FlashSuccess("头像更换成功")
	this.RedirectTo(this.redirectUrl)
}

// @router /member/logout [get]
func (this *MemberController) Logout() {
	this.SetSession("login_user_id", 0)
	this.FlashSuccess("已安全退出")
	this.RedirectTo("/");
}

// @router /member/profile [get]
func (this *MemberController) Profile() {
	this.Layout = "layout/member.tpl"
}

// @router /member/profile [post]
func (this *MemberController) SaveProfileHandler() {
	this.redirectUrl = beego.URLFor("MemberController.Profile")
	profileData := validations.MemberProfileValidation{}
	this.ValidatorAuto(&profileData)

	this.CurrentLoginUser.Company = profileData.Company
	this.CurrentLoginUser.Age = profileData.Age
	this.CurrentLoginUser.Profession = profileData.Profession
	this.CurrentLoginUser.Website = profileData.Website
	this.CurrentLoginUser.Weibo = profileData.Weibo
	this.CurrentLoginUser.Wechat = profileData.Wechat

	if _, err := orm.NewOrm().Update(this.CurrentLoginUser); err != nil {
		this.FlashError("资料保存失败")
	} else {
		this.FlashSuccess("保存成功")
	}
	this.RedirectTo(this.redirectUrl)
}
