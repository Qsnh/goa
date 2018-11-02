package controllers

import (
	"github.com/Qsnh/goa/libs"
	"github.com/Qsnh/goa/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"

	// 过滤
	categoryId, err := this.GetInt64("category_id")
	keywords := this.GetString("keywords")
	category := models.Categories{}

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	pageSize := int64(16)
	startPos := (page - 1) * pageSize
	questions := []models.Questions{}

	db := orm.NewOrm().QueryTable("questions")
	if categoryId > 0 {
		err = orm.NewOrm().QueryTable("categories").Filter("id", categoryId).One(&category)
		if err != nil {
			logs.Info(err)
			this.StopRun()
		}
		db = db.Filter("category_id", category.Id)
	}
	if keywords != "" {
		db = db.Filter("title__icontains", keywords)
	}

	count, err := db.Count()
	if err != nil {
		logs.Info(err)
		this.StopRun()
	}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&questions)

	paginator := new(libs.BootstrapPaginator)
	paginator.Instance(count, page, pageSize, beego.URLFor("IndexController.Index"))

	this.Data["Paginator"] = paginator.Render()
	this.Data["Questions"] = questions
	this.Data["Keywords"] = keywords
	this.Data["Category"] = category
	this.Layout = "layout/app.tpl"
}
