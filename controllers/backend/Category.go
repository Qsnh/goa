package backend

import (
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/validations/backend"
	"github.com/astaxie/beego/orm"
	"time"
)

type CategoryController struct {
	Base
}

// @router /backend/categories [get]
func (this *CategoryController) Index() {
	var categories []models.Categories
	if _, err := orm.NewOrm().QueryTable("categories").All(&categories); err != nil {
		this.errorHandler(err)
	}
	data := map[string]interface{}{"categories": categories}
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/category [post]
func (this *CategoryController) Store() {
	categoryValidation := backend.CategoryValidation{}
	this.ValidatorAuto(&categoryValidation)
	category := models.Categories{}
	category.Name = categoryValidation.Name
	category.CreatedAt = time.Now()
	if _, err := orm.NewOrm().Insert(&category); err != nil {
		this.errorHandler(err)
	}
	this.successResponse()
}

// @router /backend/category/:id [get]
func (this *CategoryController) Detail() {
	categoryId := this.Ctx.Input.Param(":id")
	category := models.Categories{}
	if err := orm.NewOrm().QueryTable("categories").Filter("id", categoryId).One(&category); err != nil {
		this.errorHandler(err)
	}
	data := map[string]interface{}{"category": category}
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

// @router /backend/category/:id [put]
func (this *CategoryController) Update() {
	categoryId := this.Ctx.Input.Param(":id")
	category := models.Categories{}
	if err := orm.NewOrm().QueryTable("categories").Filter("id", categoryId).One(&category); err != nil {
		this.errorHandler(err)
	}
	name := this.GetString("name")
	exists := orm.NewOrm().QueryTable("categories").Filter("name", name).Exist()
	if exists {
		this.warningResponse("分类名已经存在")
	}
	category.Name = name
	if _, err := orm.NewOrm().Update(&category); err != nil {
		this.errorHandler(err)
	}
	this.successResponse()
}

// @router /backend/category/:id [delete]
func (this *CategoryController) Destroy() {
	categoryId := this.Ctx.Input.Param(":id")
	category := models.Categories{}
	if err := orm.NewOrm().QueryTable("categories").Filter("id", categoryId).One(&category); err != nil {
		this.errorHandler(err)
	}
	questionCount, err := orm.NewOrm().QueryTable("questions").Filter("category_id", category.Id).Count()
	if err != nil {
		this.errorHandler(err)
	}
	if questionCount > 0 {
		this.warningResponse("该分类下有问题，无法删除")
	}
	if _, err := orm.NewOrm().Delete(&category); err != nil {
		this.errorHandler(err)
	}
	this.successResponse()
}
