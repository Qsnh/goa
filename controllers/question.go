package controllers

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/validations"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

type QuestionController struct {
	Base
}

// @router  /member/questions/create [get]
func (this *QuestionController) Create() {
	this.Layout = "layout/app.tpl"

	categories, err := models.AllCategories()
	if err != nil {
		this.FlashError("读取分类失败")
		this.RedirectTo("/")
	}

	this.Data["categories"] = categories
}

// @router /member/questions/create [post]
func (this *QuestionController) Store() {
	this.redirectUrl = beego.URLFor("QuestionController.Create")
	questionData := validations.QuestionStoreValidation{}
	this.ValidatorAuto(&questionData)

	id, err := models.CreateQuestion(questionData.CategoryId, questionData.Title, questionData.Description, this.CurrentLoginUser)
	if err != nil || id <= 0 {
		this.FlashError("问题创建失败")
		this.RedirectTo(this.redirectUrl)
	}

	this.FlashSuccess("问题创建成功")
	this.RedirectTo("/")
}

// @router /questions/:id [get]
func (this *QuestionController) Show() {
	question, err := models.FindQuestionById(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.FlashError("问题不存在")
		this.RedirectTo("/")
	}
	if question.IsBan == 1 {
		this.FlashError("该问题已被禁止")
		this.RedirectTo("/")
	}

	question.Description = string(blackfriday.MarkdownCommon([]byte(question.Description)))

	this.Data["question"] = question
	this.Layout = "layout/app.tpl"
}
