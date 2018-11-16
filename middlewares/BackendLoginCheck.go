package middlewares

import (
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
)

var BackendLoginCheckFilter = func(ctx *context.Context) {
	tokenHeader := ctx.Request.Header["token"]
	if len(tokenHeader) <= 0 {
		ctx.Abort(403, "NO AUTH")
	}
	sign := utils.SHA256Encode(os.Getenv("BACKEND_USER") + os.Getenv("BACKEND_PASS"))
	if sign != tokenHeader[0] {
		ctx.Abort(403, "NO AUTH")
	}
}

func BackendLoginCheck() {
	beego.InsertFilter("/backend/*", beego.BeforeRouter, BackendLoginCheckFilter)
}
