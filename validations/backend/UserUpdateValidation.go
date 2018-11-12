package backend

type UserUpdateValidation struct {
	Password string `json:"password"`
	IsLock   int    `json:"is_lock"`
}
