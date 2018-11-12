package frontend

import (
	"bytes"
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/Qsnh/goa/validations/fronted"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	template2 "text/template"
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
	this.Data["PageTitle"] = "登录"
	this.Data["PageKeywords"] = os.Getenv("SEO_INDEX_KEYWORDS")
	this.Data["PageDescription"] = os.Getenv("SEO_INDEX_DESCRIPTION")
}

// @router /login [post]
func (this *UserController) LoginHandler() {
	this.redirectUrl = beego.URLFor("UserController.Login")
	loginData := fronted.UserLoginValidation{}
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
	this.SetSession("login_user_id", 0)
	this.CurrentLoginUser = nil
	this.FlashSuccess("已安全退出")
	this.RedirectTo("/")
}

// @router /register [get]
func (this *UserController) Register() {
	this.Layout = "layout/app.tpl"
	this.Data["PageTitle"] = "注册"
	this.Data["PageKeywords"] = os.Getenv("SEO_INDEX_KEYWORDS")
	this.Data["PageDescription"] = os.Getenv("SEO_INDEX_DESCRIPTION")
}

// @router /register [post]
func (this *UserController) RegisterHandler() {
	this.redirectUrl = beego.URLFor("UserController.Register")
	userData := fronted.UserRegisterValidation{}
	this.ValidatorAuto(&userData)

	_, err := models.CreateUser(userData.Nickname, userData.Email, userData.Password)
	if err != nil {
		this.FlashError("注册失败")
		this.RedirectTo(this.redirectUrl)
	}

	this.FlashSuccess("注册成功")
	this.RedirectTo(beego.URLFor("UserController.Login"))
}

// @router /password/find [get]
func (this *UserController) FindPassword() {
	if this.CurrentLoginUser != nil {
		this.RedirectTo("/")
	}
	this.Layout = "layout/app.tpl"
	this.Data["PageTitle"] = "找回密码"
	this.Data["PageKeywords"] = os.Getenv("SEO_INDEX_KEYWORDS")
	this.Data["PageDescription"] = os.Getenv("SEO_INDEX_DESCRIPTION")
}

// @router /password/find [post]
func (this *UserController) FindPasswordHandler() {
	this.redirectUrl = beego.URLFor("UserController.FindPassword")
	if this.CurrentLoginUser != nil {
		this.RedirectTo("/")
	}
	email := this.GetString("username")
	if email == "" {
		this.FlashError("请输入邮箱")
		this.RedirectTo(this.redirectUrl)
	}
	user := models.Users{}
	err := orm.NewOrm().QueryTable("users").Filter("email", email).One(&user)
	if err != nil {
		this.FlashError("邮箱不存在")
		this.RedirectTo(this.redirectUrl)
	}
	template, err := template2.ParseFiles("./views/emails/findpassword.tpl")
	if err != nil {
		this.ErrorHandler(err)
	}
	data := map[string]string{
		"Url": user.GenerateHashedUrl(beego.URLFor("UserController.PasswordReset")),
	}
	html := new(bytes.Buffer)
	err = template.Execute(html, data)
	if err != nil {
		this.ErrorHandler(err)
	}
	go utils.SendMail(email, "密码重置链接", html.String())
	this.FlashSuccess("密码重置邮件发送成功")
	this.RedirectTo(this.redirectUrl)
}

// @router /password/reset [get]
func (this *UserController) PasswordReset() {
	if this.CurrentLoginUser != nil {
		this.RedirectTo("/")
	}
	userId, _ := this.GetInt("id")
	user, err := models.FindUserById(userId)
	if err != nil {
		this.ErrorHandler(err)
	}
	sign := this.GetString("sign")
	time := this.GetString("time")
	if user.CheckHash(sign, time) == false {
		this.FlashError("参数错误2")
		this.RedirectTo("/")
	}
	this.Layout = "layout/app.tpl"
	this.Data["PageTitle"] = "重置密码"
	this.Data["sign"] = sign
	this.Data["time"] = time
	this.Data["user"] = user
}

// @router /password/reset [post]
func (this *UserController) PasswordResetHandler() {
	passwordResetData := fronted.PasswordResetValidation{}
	this.ValidatorAuto(&passwordResetData)

	userId, _ := this.GetInt("id")
	user, err := models.FindUserById(userId)
	if err != nil {
		this.ErrorHandler(err)
	}
	sign := this.GetString("sign")
	time := this.GetString("time")
	if user.CheckHash(sign, time) == false {
		this.ErrorHandler(err)
	}

	user.Password = utils.SHA256Encode(passwordResetData.Password)
	if _, err := orm.NewOrm().Update(user); err != nil {
		this.FlashError("系统错误")
		this.Back()
	}
	this.FlashSuccess("修改成功，请重新登录")
	this.RedirectTo(beego.URLFor("UserController.Login"))
}

// @router /user/active [get]
func (this *UserController) ActiveHandler() {
	userId, _ := this.GetInt("id")
	user, err := models.FindUserById(userId)
	if err != nil {
		this.ErrorHandler(err)
	}
	sign := this.GetString("sign")
	time := this.GetString("time")
	if user.CheckHash(sign, time) == false {
		this.ErrorHandler(err)
	}

	user.IsLock = models.IS_LOCK_NO
	if _, err := orm.NewOrm().Update(user); err != nil {
		this.ErrorHandler(err)
	}

	this.FlashSuccess("激活成功")
	this.RedirectTo("/")
}