package validations

import (
	"github.com/astaxie/beego/validation"
)

type PasswordResetValidation struct {
	Password             string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func (urv *PasswordResetValidation) Valid(valid *validation.Validation) {
	if urv.Password != urv.PasswordConfirmation {
		valid.SetError("password", "两次输入密码不一致")
		return
	}
}
