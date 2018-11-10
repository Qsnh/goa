package fronted

type UserLoginValidation struct {
	Email string `form:"username" valid:"Required; Email; MaxSize(64)"`
	Password string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
}