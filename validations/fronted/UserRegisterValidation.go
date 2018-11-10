package validations

import (
    "github.com/Qsnh/goa/models"
    "github.com/astaxie/beego/validation"
)

type UserRegisterValidation struct {
    Nickname             string `form:"nickname" valid:"Required; MinSize(2); MaxSize(16)"`
    Email                string `form:"username" valid:"Required; Email; MaxSize(64)"`
    Password             string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
    PasswordConfirmation string `form:"password_confirmation"`
}

func (urv *UserRegisterValidation) Valid(valid *validation.Validation) {
    if urv.Password != urv.PasswordConfirmation {
        valid.SetError("password", "两次输入密码不一致")
        return
    }

    if models.UserNicknameExists(urv.Nickname) == true {
        valid.SetError("nickname", "昵称已经存在")
        return
    }

    if models.UserEmailExists(urv.Email) == true {
        valid.SetError("nickname", "邮箱已经存在")
        return
    }
}
