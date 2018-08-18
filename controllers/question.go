package controllers

import (
	"goa/validations"
	"goa/models"
	"github.com/astaxie/beego"
)

type QuestionController struct {
	Base
}

// @router  /member/questions/create [get]
func (this *QuestionController) Create()  {
	this.Layout = "layout/app.tpl"

	categories, err := models.AllCategories()
	if err != nil {
		this.FlashError("读取分类失败")
		this.RedirectTo("/")
	}

	this.Data["categories"] = categories
}

// @router /member/questions/create [post]
func (this *UserController) Store()  {
	this.redirectUrl = beego.URLFor("QuestionController.Create")
	questionData := validations.QuestionStoreValidation{}
	this.ValidatorAuto(questionData)

	result, err := models.CreateQuestion(questionData.CategoryId, questionData.Title, questionData.Description, this.CurrentLoginUser)
	if err != nil || result <= 0 {
		this.FlashError("问题创建失败")
		this.RedirectTo(this.redirectUrl)
	}

	this.FlashSuccess("问题创建成功")
	this.RedirectTo("/")
}