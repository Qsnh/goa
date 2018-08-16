package validations

type UserRegisterValidation struct {
    Nickname string `form:"nickname";valid:"Required;Min(2);Max(16)"`
    Email string `form:"email";valid:"Required;Max(64)"`
    Password string `form:"password";valid:"Required;Email;Min(6);Max(16)"`
}