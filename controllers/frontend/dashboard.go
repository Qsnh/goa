package frontend

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type DashboardController struct {
	Base
}

// @router /user/:user_id [get]
func (this *DashboardController) MemberQuestions() {
	this.Layout = "layout/app.tpl"
	userIdString := this.Ctx.Input.Param(":user_id")
	userId, _ := strconv.Atoi(userIdString)
	user, err := models.FindUserById(userId)
	if err != nil {
		this.ErrorHandler(err)
	}

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	pageSize := int64(10)
	startPos := (page - 1) * pageSize

	db := orm.NewOrm().QueryTable("questions").Filter("user_id", user.Id)
	total, err := db.Count()
	if err != nil {
		this.ErrorHandler(err)
	}
	questions := []models.Questions{}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&questions)

	paginator := new(utils.BootstrapPaginator)
	paginator.Instance(total, page, pageSize, beego.URLFor("DashboardController.MemberQuestion"))

	IndexUrl := beego.URLFor("IndexController.Index")

	this.Data["Paginator"] = paginator.Render()
	this.Data["Questions"] = questions
	this.Data["Indexurl"] = IndexUrl
	this.Data["User"] = user
	this.Data["PageTitle"] = user.Nickname + "的主页"
}

// @router /user/:user_id/answers [get]
func (this *DashboardController) MemberAnswers() {
	this.Layout = "layout/app.tpl"
	userIdString := this.Ctx.Input.Param(":user_id")
	userId, _ := strconv.Atoi(userIdString)
	user, err := models.FindUserById(userId)
	if err != nil {
		this.ErrorHandler(err)
	}

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	pageSize := int64(10)
	startPos := (page - 1) * pageSize

	db := orm.NewOrm().QueryTable("answers").Filter("user_id", user.Id)
	total, err := db.Count()
	if err != nil {
		this.ErrorHandler(err)
	}
	answers := []models.Answers{}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&answers)

	paginator := new(utils.BootstrapPaginator)
	paginator.Instance(total, page, pageSize, beego.URLFor("DashboardController.MemberQuestion"))

	IndexUrl := beego.URLFor("IndexController.Index")

	this.Data["Paginator"] = paginator.Render()
	this.Data["Answers"] = answers
	this.Data["Indexurl"] = IndexUrl
	this.Data["User"] = user
	this.Data["PageTitle"] = user.Nickname + "的回答"
}
