package routers

import (
	"bee-web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/news/", &controllers.CategoryController{})
	beego.Router("/news/?:time([0-9]+)/?:id([0-9]+).html", &controllers.ArticleController{})
}
