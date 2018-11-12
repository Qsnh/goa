package frontend

import (
	"bytes"
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/Qsnh/goa/validations/fronted"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dchest/captcha"
	template2 "html/template"
)

type MemberController struct {
	Base
}

// @router /member [get]
func (this *MemberController) Index() {
	this.Layout = "layout/member.tpl"

	user := this.CurrentLoginUser
	questionCount, _ := orm.NewOrm().QueryTable("questions").Filter("user_id", user.Id).Count()
	answerCount, _ := orm.NewOrm().QueryTable("Answers").Filter("user_id", user.Id).Count()
	rate := 0

	this.Data["QuestionCount"] = questionCount
	this.Data["AnswerCount"] = answerCount
	this.Data["Rate"] = rate
	this.Data["PageTitle"] = user.Nickname+"的个人中心"
}

// @router /member/password [get]
func (this *MemberController) ChangePassword() {
	this.Layout = "layout/member.tpl"
	this.Data["PageTitle"] = this.CurrentLoginUser.Nickname+" - 修改密码"
}

// @router /member/password [post]
func (this *MemberController) ChangePasswordHandler() {
	this.redirectUrl = beego.URLFor("MemberController.ChangePassword")
	passwordData := fronted.MemberChangePasswordValidation{}
	this.ValidatorAuto(&passwordData)

	if this.CurrentLoginUser.Password != utils.SHA256Encode(passwordData.OldPassword) {
		this.FlashError("原密码不正确")
		this.RedirectTo(this.redirectUrl)
	}

	this.CurrentLoginUser.Password = utils.SHA256Encode(passwordData.NewPassword)
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
	this.Data["xsrf"] = this.XSRFToken()
	this.Data["PageTitle"] = this.CurrentLoginUser.Nickname+" - 修改头像"
}

// @router /member/avatar [post]
func (this *MemberController) ChangeAvatarHandler() {
	this.CurrentLoginUser.Avatar = this.GetString("avatar")
	if result, err := orm.NewOrm().Update(this.CurrentLoginUser); err != nil || result == 0 {
		this.FlashSuccess("头像保存失败")
	} else {
		this.FlashSuccess("头像更换成功")
	}
	this.Back()
}

// @router /member/logout [get]
func (this *MemberController) Logout() {
	this.SetSession("login_user_id", 0)
	this.FlashSuccess("已安全退出")
	this.RedirectTo("/")
}

// @router /member/profile [get]
func (this *MemberController) Profile() {
	this.Layout = "layout/member.tpl"
	this.Data["PageTitle"] = this.CurrentLoginUser.Nickname+" - 修改我的资料"
}

// @router /member/profile [post]
func (this *MemberController) SaveProfileHandler() {
	this.redirectUrl = beego.URLFor("MemberController.Profile")
	profileData := fronted.MemberProfileValidation{}
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

// @router /member/questions [get]
func (this *MemberController) Questions() {
	this.Layout = "layout/member.tpl"

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	var pageSize int64
	pageSize = 8
	startPos := (page - 1) * pageSize

	questions := []models.Questions{}
	_, _ = orm.NewOrm().QueryTable("questions").Filter("user_id", this.CurrentLoginUser.Id).RelatedSel().OrderBy("-created_at", "-id").Limit(pageSize, startPos).All(&questions)
	count, _ := orm.NewOrm().QueryTable("questions").Filter("user_id", this.CurrentLoginUser.Id).Count()

	paginator := new(utils.BootstrapPaginator)
	paginator.Instance(count, page, pageSize, beego.URLFor("MemberController.Questions"))

	this.Data["Questions"] = questions
	this.Data["Paginator"] = paginator.Render()
	this.Data["PageTitle"] = this.CurrentLoginUser.Nickname+"提出的问题"
}

// @router /member/answers [get]
func (this *MemberController) Answers() {
	this.Layout = "layout/member.tpl"

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	var pageSize int64
	pageSize = 8
	startPos := (page - 1) * pageSize

	answers := []models.Answers{}
	_, _ = orm.NewOrm().QueryTable("answers").Filter("user_id", this.CurrentLoginUser.Id).RelatedSel().OrderBy("-created_at", "-id").Limit(pageSize, startPos).All(&answers)
	count, _ := orm.NewOrm().QueryTable("answers").Filter("user_id", this.CurrentLoginUser.Id).Count()

	paginator := new(utils.BootstrapPaginator)
	paginator.Instance(count, page, pageSize, beego.URLFor("MemberController.Answers"))

	this.Data["Answers"] = answers
	this.Data["Paginator"] = paginator.Render()
	this.Data["PageTitle"] = this.CurrentLoginUser.Nickname+"回答的问题"
}

// @router /member/active/mail/send [get]
func (this *MemberController) SendActiveMail() {
	this.Layout = "layout/member.tpl"
	this.Data["user"] = this.CurrentLoginUser
	this.Data["PageTitle"] = "激活"
}

// @router /member/active/mail/send [post]
func (this *MemberController) SendActiveMailHandler() {
	if this.CurrentLoginUser.IsLock != models.IS_LOCK_YES {
		this.FlashError("您的账户已经激活啦")
		this.Back()
	}
	if captcha.VerifyString(this.GetSession("captcha_id").(string), this.GetString("captcha_code")) == false {
		this.FlashError("验证码错误")
		this.Back()
	}
	template, err := template2.ParseFiles("./views/emails/verify.tpl")
	if err != nil {
		this.ErrorHandler(err)
	}
	data := map[string]string{
		"Url": this.CurrentLoginUser.GenerateHashedUrl(beego.URLFor("UserController.ActiveHandler")),
	}
	html := new(bytes.Buffer)
	err = template.Execute(html, data)
	if err != nil {
		this.ErrorHandler(err)
	}
	go utils.SendMail(this.CurrentLoginUser.Email, "密码重置链接", html.String())
	this.FlashSuccess("激活邮件发送成功，有效期一个小时，请尽快操作")
	this.RedirectTo("/")
}