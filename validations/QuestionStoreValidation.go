package validations

import (
	"github.com/astaxie/beego/validation"
	"goa/models"
)

type QuestionStoreValidation struct {
	CategoryId int64 `form:"category_id" valid:"Required"`
	Title string `form:"title" valid:"Required; MinSize(10); MaxSize(255)"`
	Description string `form:"description" valid:"Required"`
}

func (validation *QuestionStoreValidation) Valid(valid *validation.Validation)  {
	if _, err := models.FindCategoryById(validation.CategoryId); err != nil {
		valid.SetError("category_id", "分类不存在")
		return
	}
}

