package backend

import (
	"github.com/Qsnh/goa/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type QuestionController struct {
	Base
}

// @router /backend/questions [get]
func (this *QuestionController) Index()  {
	// 过滤
	categoryId, _ := this.GetInt64("category_id")
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
		if err := orm.NewOrm().QueryTable("categories").Filter("id", categoryId).One(&category); err != nil {
			this.errorHandler(err)
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

	data := make(map[string]interface{})
	data["questions"] = questions
	data["total"] = count
	data["page"] = page
	data["page_size"] = pageSize
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/question/:id [delete]
func (this *QuestionController) Destroy()  {
	questionId := this.Ctx.Input.Param(":id")
	question := models.Questions{}
	if err := orm.NewOrm().QueryTable("questions").Filter("id", questionId).One(&question); err != nil {
		this.errorHandler(err)
	}
	if _, err := orm.NewOrm().Delete(&question); err != nil {
		this.errorHandler(err)
	}
	this.successResponse()
}