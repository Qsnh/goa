package controllers

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/validations"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"time"
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
	this.Data["PageTitle"] = "我要提问"
}

// @router /member/questions/create [post]
func (this *QuestionController) Store() {
	this.redirectUrl = beego.URLFor("QuestionController.Create")
	questionData := validations.QuestionStoreValidation{}
	this.ValidatorAuto(&questionData)

	_, err := models.CreateQuestion(questionData.CategoryId, questionData.Title, questionData.Description, this.CurrentLoginUser)
	if err != nil {
		logs.Info(err)
		this.FlashError("问题创建失败")
		this.RedirectTo(this.redirectUrl)
	}

	this.FlashSuccess("问题创建成功")
	this.RedirectTo("/")
}

// @router  /member/questions/:question_id/edit [get]
func (this *QuestionController) Edit() {
	questionId := this.Ctx.Input.Param(":question_id")
	question, err := models.FindQuestionById(questionId)
	if err != nil {
		this.FlashError("问题不存在")
		this.RedirectTo("/")
	}
	if this.CurrentLoginUser.Id != question.User.Id {
		this.Abort("401")
		this.StopRun()
	}

	categories, err := models.AllCategories()
	if err != nil {
		this.FlashError("读取分类失败")
		this.RedirectTo("/")
	}

	this.Data["Question"] = question
	this.Data["categories"] = categories
	this.Data["PageTitle"] = "问题编辑："+question.Title
}

// @router /member/questions/:question_id/edit [post]
func (this *QuestionController) Update() {
	questionId := this.Ctx.Input.Param(":question_id")
	question, err := models.FindQuestionById(questionId)
	if err != nil {
		this.FlashError("问题不存在")
		this.RedirectTo("/")
	}
	if this.CurrentLoginUser.Id != question.User.Id {
		this.Abort("401")
		this.StopRun()
	}

	this.redirectUrl = beego.URLFor("QuestionController.Edit", ":question_id", questionId)
	questionData := validations.QuestionStoreValidation{}
	this.ValidatorAuto(&questionData)

	category := models.Categories{}
	err = orm.NewOrm().QueryTable("categories").Filter("id", questionData.CategoryId).One(&category)
	if err != nil {
		logs.Info(err)
		this.Abort("404")
		this.StopRun()
	}

	question.Category = &category
	question.Title = questionData.Title
	question.Description = questionData.Description

	if _, err := orm.NewOrm().Update(question); err != nil {
		this.FlashError("问题修改失败")
	} else {
		this.FlashSuccess("问题修改成功")
	}
	this.RedirectTo(this.redirectUrl)
}

// @router /questions/:id [get]
func (this *QuestionController) Show() {
	questionId := this.Ctx.Input.Param(":id")
	question, err := models.FindQuestionById(questionId)
	if err != nil {
		this.FlashError("问题不存在")
		this.RedirectTo("/")
	}
	if question.IsBan == 1 {
		this.FlashError("该问题已被禁止")
		this.RedirectTo("/")
	}

	question.Description = string(blackfriday.MarkdownCommon([]byte(question.Description)))

	// 回答
	page, _ := this.GetInt64("page")
	pageSize := int64(8)

	answers, paginator, err := models.AnswerPaginate(questionId, page, pageSize)
	if err != nil {
		logs.Info(err)
		this.FlashError("系统错误")
		this.RedirectTo("/")
	}

	this.Data["question"] = question
	this.Data["Answers"] = answers
	this.Data["Paginator"] = paginator.Render()
	this.Data["PageTitle"] = question.Title
	this.Data["PageKeywords"] = question.Title
	this.Data["PageDescription"] = question.Description
	this.Layout = "layout/app.tpl"
}

// @router /member/questions/:id [post]
func (this *QuestionController) AnswerHandler() {
	questionId := this.Ctx.Input.Param(":id")
	this.redirectUrl = beego.URLFor("QuestionController.Show", ":id", questionId)
	questionData := validations.AnswerValidation{}
	this.ValidatorAuto(&questionData)

	question, err := models.FindQuestionById(questionId)
	if err != nil {
		this.FlashError("问题不存在");
		this.RedirectTo(this.redirectUrl)
	}

	orm := orm.NewOrm()
	orm.Begin()

	// 创建回答
	_, err = models.AnswerCreate(this.CurrentLoginUser, question, questionData.Description, &orm)
	if err != nil {
		orm.Rollback()
		logs.Info(err)
		this.FlashError("系统发生错误了哦")
		this.RedirectTo(this.redirectUrl)
	}

	// 更新问题
	question.AnswerUser = this.CurrentLoginUser
	question.AnswerAt = time.Now()
	question.AnswerCount += 1
	question.UpdatedAt = time.Now()
	if _, err := orm.Update(question); err != nil {
		orm.Rollback()
		logs.Info(err)
		this.FlashError("系统发生错误了哦")
		this.RedirectTo(this.redirectUrl)
	}
	orm.Commit()

	this.FlashSuccess("回答成功")
	this.RedirectTo(this.redirectUrl)
}