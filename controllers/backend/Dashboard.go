package backend

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type DashboardController struct {
	Base
}

// @router /backend/dashboard [get]
func (this *DashboardController) Index() {
	today := time.Now().Format("2006-01-01")
	// 今日注册
	registerCount, err := orm.NewOrm().QueryTable("users").Filter("created_at__gte", today).Count()
	if err != nil {
		this.errorHandler(err)
	}
	// 今日提问
	questionCount, err := orm.NewOrm().QueryTable("questions").Filter("created_at__gte", today).Count()
	if err != nil {
		this.errorHandler(err)
	}
	// 今日回答
	answerCount, err := orm.NewOrm().QueryTable("answers").Filter("created_at__gte", today).Count()
	if err != nil {
		this.errorHandler(err)
	}

	data := make(map[string]interface{})
	data["count_register"] = registerCount
	data["count_question"] = questionCount
	data["count_answer"] = answerCount
	this.Data["json"] = map[string]interface{}{"data": data}
	this.ServeJSON()
	this.StopRun()
}
