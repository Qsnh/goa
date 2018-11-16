package fronted

import "github.com/astaxie/beego/validation"

type MemberChangePasswordValidation struct {
	OldPassword             string `form:"old_password" valid:"Required; MinSize(6); MaxSize(16)"`
	NewPassword             string `form:"new_password" valid:"Required; MinSize(6); MaxSize(16)"`
	NewPasswordConfirmation string `form:"new_password_confirmation"`
}

func (validation *MemberChangePasswordValidation) Valid(valid *validation.Validation) {
	if validation.NewPassword != validation.NewPasswordConfirmation {
		valid.SetError("new_password", "两次输入的密码不一致")
		return
	}
}
