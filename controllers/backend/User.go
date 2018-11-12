package backend

import (
	"encoding/json"
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/Qsnh/goa/validations/backend"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	Base
}

// @router /backend/users [get]
func (this *UserController) Index() {
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
	data["code"] = 0
	this.Data["json"] = map[string]interface{}{"data": data}
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/user/:id [get]
func (this *UserController) Show() {
	userId := this.Ctx.Input.Param(":id")
	userIdString, _ := strconv.Atoi(userId)
	user, err := models.FindUserById(userIdString)
	if err != nil {
		this.errorHandler(err)
	}
	this.Data["json"] = map[string]interface{}{"user": user}
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/user/:id [put]
func (this *UserController) Update() {
	userId := this.Ctx.Input.Param(":id")
	user := models.Users{}
	if err := orm.NewOrm().QueryTable("users").Filter("id", userId).One(&user); err != nil {
		this.errorHandler(err)
	}

	userUpdateValidation := backend.UserUpdateValidation{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &userUpdateValidation); err != nil {
		this.errorHandler(err)
	}
	user.IsLock = userUpdateValidation.IsLock
	if userUpdateValidation.Password != "" {
		user.Password = utils.SHA256Encode(userUpdateValidation.Password)
	}

	if _, err := orm.NewOrm().Update(&user); err != nil {
		this.errorHandler(err)
	}

	this.successResponse()
}
