package controllers

import "goa/models"

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"

	page, _ := this.GetInt("page")
	pageSize, _ := this.GetInt64("page_size")

	questions, paginator, _ := models.Paginate(page, pageSize, this.Ctx.Request)

	this.Data["paginator"] = paginator.Render()
	this.Data["questions"] = questions
}