package controllers

import (
	"github.com/Qsnh/goa/models"
)

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"

	page, _ := this.GetInt64("page")
	pageSize := int64(16)

	questions, paginator, err := models.QuestionPaginate(page, pageSize)
	if err != nil {
		this.StopRun()
	}

	this.Data["Paginator"] = paginator.Render()
	this.Data["Questions"] = questions
	this.Layout = "layout/app.tpl"
}
