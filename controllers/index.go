package controllers

type IndexController struct {
	Base
}

// @router / [get]
func (this *IndexController) Index() {
	this.Layout = "layout/app.tpl"
}