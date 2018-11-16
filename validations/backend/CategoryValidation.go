package backend

import (
	"github.com/Qsnh/goa/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type CategoryValidation struct {
	Name string `form:"name" valid:"Required; MaxSize(64)"`
}

func (this *CategoryValidation) Valid(validation *validation.Validation) {
	category := models.Categories{}
	err := orm.NewOrm().QueryTable("categories").Filter("name", this.Name).One(&category)
	if err == nil {
		validation.SetError("name", "分类名已经存在")
	}
}
