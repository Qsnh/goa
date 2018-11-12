package backend

import (
	"encoding/json"
	"github.com/Qsnh/goa/request/backend"
	"github.com/Qsnh/goa/utils"
	"os"
)

type LoginController struct {
	Base
}

// @router /backend/login [post]
func (this *LoginController) LoginHandler()  {
	loginRequest := backend.LoginRequest{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginRequest); err != nil {
		this.errorHandler(err)
	}
	if loginRequest.Username != os.Getenv("BACKEND_USER") || loginRequest.Password != os.Getenv("BACKEND_PASS") {
		this.warningResponse("用户名或密码错误")
	}
	token := utils.SHA256Encode(loginRequest.Username + loginRequest.Password)
	data := map[string]string{"token": token, "name": "管理员", "id": "1"}
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}