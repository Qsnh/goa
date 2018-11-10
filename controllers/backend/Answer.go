package backend

import (
	"github.com/Qsnh/goa/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AnswerContoller struct {
	Base
}

// @router /backend/answers [get]
func (this *AnswerContoller) Index()  {
	// 过滤
	questionId := this.GetString("question_id")
	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	pageSize := int64(10)
	startPos := (page - 1) * pageSize
	answers := []models.Answers{}

	db := orm.NewOrm().QueryTable("answers")
	if questionId != "" {
		db = db.Filter("question_id", questionId)
	}

	count, err := db.Count()
	if err != nil {
		logs.Info(err)
		this.StopRun()
	}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&answers)

	data := make(map[string]interface{})
	data["answers"] = answers
	data["total"] = count
	data["page"] = page
	data["page_size"] = pageSize
	data["code"] = 0
	this.Data["json"] = map[string]interface{}{"data": data}
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/answer/:id [delete]
func (this *AnswerContoller) Destroy()  {
	answerId := this.Ctx.Input.Param(":id")
	answer := models.Answers{}
	if err := orm.NewOrm().QueryTable("answers").Filter("id", answerId).One(&answer); err != nil {
		this.errorHandler(err)
	}
	if _, err := orm.NewOrm().Delete(&answer); err != nil {
		this.errorHandler(err)
	}
	this.successResponse()
}