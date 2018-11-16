package frontend

import (
	"github.com/Qsnh/goa/goaio"
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"html/template"
	"strconv"
)

type Base struct {
	beego.Controller
	FlashBag         *beego.FlashData
	redirectUrl      string
	CurrentLoginUser *models.Users
	SettingData      map[string]string
}

func (Base *Base) Prepare() {
	// 读取配置
	settings := []models.Settings{}
	if _, err := orm.NewOrm().QueryTable("settings").All(&settings); err != nil {
		Base.ErrorHandler(err)
	}
	settingData := make(map[string]string)
	for _, item := range settings {
		settingData[item.Name] = item.Value
	}
	Base.SettingData = settingData

	// 初始化读取Flash
	beego.ReadFromRequest(&Base.Controller)

	// 初始化Flash
	Base.FlashBag = beego.NewFlash()

	// 初始化XSRF
	Base.Data["xsrfdata"] = template.HTML(Base.XSRFFormHTML())

	// 自动读取当前登陆用户
	isLogin := false
	loginUserId, _ := strconv.Atoi(Base.Ctx.GetCookie("login_user_id"))
	loginUserSign := Base.Ctx.GetCookie("login_user_sign")
	if loginUserId > 0 {
		user, err := models.FindUserById(loginUserId)
		if err == nil && utils.AuthSignCheck(user.Id, user.Email, user.Password, loginUserSign) == true {
			Base.CurrentLoginUser = user
			isLogin = true
		}
	}
	Base.Data["IsLogin"] = isLogin
	Base.Data["user"] = Base.CurrentLoginUser

	// 是否激活
	isActive := true
	if isLogin {
		if Base.CurrentLoginUser.IsLock == models.IS_LOCK_YES {
			isActive = false
		}
	}
	Base.Data["IsActive"] = isActive

	// SEO
	Base.Data["PageTitle"] = ""
	Base.Data["PageKeywords"] = ""
	Base.Data["PageDescription"] = ""
	Base.Data["AppName"] = Base.SettingData["APP_NAME"]
	Base.Data["AppIcp"] = Base.SettingData["ICP"]
}

// 保存成功的Flash信息
func (Base *Base) FlashSuccess(message string) {
	Base.FlashBag.Notice(message)
	Base.FlashBag.Store(&Base.Controller)
}

// 保存失败的Flash信息
func (Base *Base) FlashError(message string) {
	Base.FlashBag.Error(message)
	Base.FlashBag.Store(&Base.Controller)
}

// 自动化的表单验证器
func (Base *Base) ValidatorAuto(frontendData interface{}) {

	if err := Base.ParseForm(frontendData); err != nil {
		Base.FlashBag.Error("参数解析错误")
		Base.RedirectTo(Base.redirectUrl)
	}

	defaultMessage := map[string]string{
		"Required":     "不能为空",
		"Min":          "不能小于%d",
		"Max":          "不能大于%d",
		"Range":        "取值必须在%d到%d之间",
		"MinSize":      "长度不能小于%d",
		"MaxSize":      "长度不能大于%d",
		"Length":       "长度必须等于%d",
		"Alpha":        "必须是字母",
		"Numeric":      "必须是数字",
		"AlphaNumeric": "必须是字母或者数字",
		"Match":        "必须出现 %s 关键字",
		"NoMatch":      "不能出现 %s 关键字",
		"AlphaDash":    "必须是字母，数组或者横线(-)",
		"Email":        "不合法的邮箱地址",
		"IP":           "不合法的IP",
		"Base64":       "不合法的Base64编码格式",
		"Mobile":       "不合法的手机号",
		"Tel":          "不合法的电话号码",
		"Phone":        "不合法的手机号",
		"ZipCode":      "不合法的邮编",
	}
	validation.SetDefaultMessage(defaultMessage)
	validate := validation.Validation{}

	isValid, err := validate.Valid(frontendData)
	if err != nil {
		Base.ErrorHandler(err)
	}

	if !isValid {
		Base.FlashError(validate.Errors[0].Message)
		Base.RedirectTo(Base.redirectUrl)
	}
}

// 重定向
func (Base *Base) RedirectTo(url string) {
	Base.Redirect(url, 302)
	Base.StopRun()
}

// ajax错误返回
func (Base *Base) AjaxError(message string, code uint16) {
	res := goaio.ErrorResponseJson{message, code}
	Base.Data["json"] = res
	Base.ServeJSON()
	Base.StopRun()
}

// ajax成功返回
func (Base *Base) AjaxSuccess(message string, data interface{}) {
	res := goaio.SuccessResponseJson{message, 0, data}
	Base.Data["json"] = res
	Base.ServeJSON()
	Base.StopRun()
}

// 抛出500
func (Base *Base) ErrorHandler(err error) {
	logs.Info(err)
	Base.Abort("500")
	Base.StopRun()
}

// 跳转到前一页
func (Base *Base) Back() {
	Base.RedirectTo(Base.Ctx.Request.Referer())
	Base.StopRun()
}
