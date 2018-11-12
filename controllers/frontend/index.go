package frontend

import (
	"os"

	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dchest/captcha"
)

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"
	baseUrl := beego.URLFor("IndexController.Index")

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
			this.ErrorHandler(err)
		}
		db = db.Filter("category_id", category.Id)
	}
	if keywords != "" {
		db = db.Filter("title__icontains", keywords)
	}

	count, err := db.Count()
	if err != nil {
		this.ErrorHandler(err)
	}
	_, _ = db.OrderBy("-updated_at", "-id").RelatedSel().Limit(pageSize, startPos).All(&questions)

	paginator := new(utils.BootstrapPaginator)
	paginator.Instance(count, page, pageSize, utils.Url(baseUrl, "keywords", keywords, "category_id", categoryId))

	this.Data["Paginator"] = paginator.Render()
	this.Data["Questions"] = questions
	this.Data["Keywords"] = keywords
	this.Data["Category"] = category
	this.Data["Baseurl"] = baseUrl
	this.Data["PageTitle"] = os.Getenv("SEO_INDEX_TITLE")
	this.Data["PageKeywords"] = os.Getenv("SEO_INDEX_KEYWORDS")
	this.Data["PageDescription"] = os.Getenv("SEO_INDEX_DESCRIPTION")
	this.Layout = "layout/app.tpl"
}

// @router /captcha [get]
func (this *IndexController) CaptchaShow() {
	captchaId := captcha.New()
	this.SetSession("captcha_id", captchaId)
	err := captcha.WriteImage(this.Ctx.ResponseWriter, captchaId, 240, 80)
	if err != nil {
		this.ErrorHandler(err)
	}
	this.StopRun()
}