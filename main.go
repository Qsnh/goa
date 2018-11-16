package main

import (
	"github.com/Qsnh/goa/middlewares"
	"github.com/Qsnh/goa/models"
	_ "github.com/Qsnh/goa/routers"
	"github.com/Qsnh/goa/tasks"
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := godotenv.Load(utils.Pwd() + string(filepath.Separator) + ".env")
	if err != nil {
		log.Fatal("can't find .env file")
		return
	}
	models.Init()

	// 中间件注册
	middlewares.LoginCheck()
	middlewares.CorsHandler()
	middlewares.BackendLoginCheck()

	// 模板函数注册
	beego.AddFuncMap("urlquery", utils.Url)
	beego.AddFuncMap("timeforhumnas", utils.TimeDiffForHumans)

	// 定时任务
	backupTask := toolbox.NewTask("backup_task", "0 0 05 * * *", tasks.WebBackupTask)
	toolbox.AddTask("backup_task", backupTask)
	toolbox.StartTask()
	defer toolbox.StopTask()

	// 运行
	beego.Run(":" + os.Getenv("APP_PORT"))
	logs.SetLogger("console")
}
