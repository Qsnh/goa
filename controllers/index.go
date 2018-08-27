package controllers

import "goa/models"

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"

	page, _ := this.GetInt64("page")
	pageSize := int64(20)

	questions, paginator, _ := models.Paginate(page, pageSize)

	this.Data["paginator"] = paginator.Render()
	this.Data["questions"] = questions
}