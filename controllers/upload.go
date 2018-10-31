package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"path"
	"strings"
	"time"
)

type UploadController struct {
	Base
}

// @router /upload/image [post]
func (this *UploadController) Image() {
	file, header, err := this.GetFile("file")
	if err != nil {
		this.AjaxError("请选择头像文件", 1)
	}
	defer file.Close()
	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		this.AjaxError("请上传有效图片文件", 2)
	}
	// 文件后缀判断
	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		this.AjaxError("请上传jpeg/png/gif图片", 3)
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		this.AjaxError("头像文件大小不能超过2MB", 4)
	}
	// 保存文件
	rand.Seed(time.Now().Unix())
	newFilename := fmt.Sprintf("%d+%d", time.Now().Unix(), rand.Intn(100))
	newFilename = fmt.Sprintf("%x", md5.Sum([]byte(newFilename)))
	path := path.Join("static/uploads/avatar", newFilename+"."+extension)
	err = this.SaveToFile("file", path)
	if err != nil {
		logs.Info(err)
		this.AjaxError("头像上传失败", 5)
	}
	res := make(map[string]string)
	res["path"] = "/" + path
	this.AjaxSuccess("上传成功", res)
}

func (this *UploadController) Prepare()  {
	this.EnableXSRF = false
}