package backend

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	Base
}

// @router /backend/users [get]
func (this *UserController) Index()  {
	// 过滤
	keywords := this.GetString("keywords")

	page, _ := this.GetInt64("page")
	if page <= 0 {
		page = 1
	}
	pageSize := int64(10)
	startPos := (page - 1) * pageSize
	users := []models.Users{}

	db := orm.NewOrm().QueryTable("users")
	if keywords != "" {
		db = db.Filter("nickname__icontains", keywords).Filter("email__icontains", keywords)
	}

	count, err := db.Count()
	if err != nil {
		logs.Info(err)
		this.StopRun()
	}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&users)

	data := make(map[string]interface{})
	data["users"] = users
	data["total"] = count
	data["page"] = page
	data["page_size"] = pageSize
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/user/:id [put]
func (this *UserController) Update()  {
	userId := this.Ctx.Input.Param(":id")
	user := models.Users{}
	if err := orm.NewOrm().QueryTable("users").Filter("id", userId).One(&user); err != nil {
		this.errorHandler(err)
	}

	password := this.GetString("password")
	isLock, _ := this.GetInt("is_lock")
	user.IsLock = isLock
	if password != "" {
		user.Password = utils.SHA256Encode(password)
	}

	if _, err := orm.NewOrm().Update(&user); err != nil {
		this.errorHandler(err)
	}

	this.successResponse()
}