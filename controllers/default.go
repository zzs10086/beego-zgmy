package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Name"] = "张正山"
	c.Layout = "layouts/base.tpl"
	c.TplName = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Foucs"] = "foucs.tpl"

}
