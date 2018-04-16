package controllers

import (
	"bee-web/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	var article = models.GetArcInfo(5)
	beego.Info(article)
	c.Data["Article"] = article
	c.Layout = "layouts/base.tpl"
	c.TplName = "article.tpl"

}
