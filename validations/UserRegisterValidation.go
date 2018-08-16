package validations

type UserRegisterValidation struct {
    Nickname string `form:"nickname" valid:"Required; MinSize(2); MaxSize(16)"`
    Email string `form:"username" valid:"Required; MaxSize(64)"`
    Password string `form:"password" valid:"Required; Email; MinSize(6); MaxSize(16)"`
}