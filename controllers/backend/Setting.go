package backend

import (
	"encoding/json"
	"github.com/Qsnh/goa/models"
	"github.com/Qsnh/goa/validations/backend"
	"github.com/astaxie/beego/orm"
)

type SettingController struct {
	Base
}

// @router /backend/setting/save [put]
func (this *SettingController) SaveHandler() {
	settingValidation := backend.SettingValidation{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &settingValidation); err != nil {
		this.errorHandler(err)
	}
	db := orm.NewOrm()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.CORS_ORIGINAL, "CORS_ORIGINAL").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.ICP, "ICP").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.MEMBER_DEFAULT_AVATAR, "MEMBER_DEFAULT_AVATAR").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.MEMBER_DEFAULT_IS_LOCK, "MEMBER_DEFAULT_IS_LOCK").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.SEO_INDEX_DESCRIPTION, "SEO_INDEX_DESCRIPTION").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.SEO_INDEX_KEYWORDS, "SEO_INDEX_KEYWORDS").Exec()
	_, _ = db.Raw("UPDATE settings set value = ? where name = ? limit 1", settingValidation.SEO_INDEX_TITLE, "SEO_INDEX_TITLE").Exec()

	this.successResponse()
	this.StopRun()
}

// @router /backend/setting/data [get]
func (this *SettingController) SettingData() {
	settingRecords := []models.Settings{}
	if _, err := orm.NewOrm().QueryTable("settings").All(&settingRecords); err != nil {
		this.errorHandler(err)
	}
	var settingMap = map[string]string{}
	for _, item := range settingRecords {
		settingMap[item.Name] = item.Value
	}
	this.Data["json"] = settingMap
	this.ServeJSON()
	this.StopRun()
}
