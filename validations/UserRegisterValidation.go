package validations

import (
    "github.com/astaxie/beego/validation"
)

type UserRegisterValidation struct {
    Nickname string `form:"nickname" valid:"Required; MinSize(2); MaxSize(16)"`
    Email string `form:"username" valid:"Required; MaxSize(64)"`
    Password string `form:"password" valid:"Required; Email; MinSize(6); MaxSize(16)"`
    PasswordConfirmation string `form:"password_confirmation"`
}

func (urv *UserRegisterValidation) valid(valid *validation.Validation)  {
    if urv.Password != urv.PasswordConfirmation {
        valid.SetError("password", "两次输入密码不一致")
    }
}