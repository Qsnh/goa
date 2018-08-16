package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/astaxie/beego/validation"
)

type Base struct {
	 beego.Controller
	 FlashBag *beego.FlashData
	 redirectUrl string
	 redirectCode int
}

func (Base *Base) Prepare() {
	// 初始化读取Flash
	beego.ReadFromRequest(&Base.Controller)

	// 初始化Flash
	Base.FlashBag = beego.NewFlash()

	// 初始化XSRF
	Base.Data["xsrfdata"] = template.HTML(Base.XSRFFormHTML())

	Base.redirectCode = 302
}

// 保存成功的Flash信息
func (Base *Base) FlashSuccess(message string) {
	Base.FlashBag.Notice(message)
	Base.FlashBag.Store(&Base.Controller)
}

// 保存失败的Flash信息
func (Base *Base) FlashError(message string)  {
	Base.FlashBag.Error(message)
	Base.FlashBag.Store(&Base.Controller)
}

// 自动化的表单验证器
func (Base *Base) ValidatorAuto(frontendData interface{})  {

	if err := Base.ParseForm(frontendData); err != nil {
		Base.FlashBag.Error("参数解析错误")
		Base.FlashBag.Store(&Base.Controller)
		Base.Redirect(Base.redirectUrl, Base.redirectCode)
		return
	}

	validate := validation.Validation{}
	isValid, err := validate.Valid(frontendData)
	if err != nil {
		Base.FlashError("服务器出错")
		Base.Redirect(Base.redirectUrl, Base.redirectCode)
		return
	}

	if !isValid {
		Base.FlashError(validate.Errors[0].Key + validate.Errors[0].Message)
		Base.Redirect(Base.redirectUrl, Base.redirectCode)
		return
	}
}