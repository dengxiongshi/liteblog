package controllers

type IndexController struct {
	BaseController
}

//@router / [get]
func (this *IndexController) Index() {
	this.TplName = "index.html"
}

//@router /about [get]
func (this *IndexController) About() {
	this.TplName = "about.html"
}

//@router /message [get]
func (this *IndexController) Message() {
	this.TplName = "message.html"
}
