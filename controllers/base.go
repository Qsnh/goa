package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type Base struct {
	 beego.Controller
	 FlashBag *beego.FlashData
}

func (Base *Base) Prepare() {
	// 初始化读取Flash
	beego.ReadFromRequest(&Base.Controller)

	// 初始化Flash
	Base.FlashBag = beego.NewFlash()

	// 初始化XSRF
	Base.Data["xsrfdata"] = template.HTML(Base.XSRFFormHTML())
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