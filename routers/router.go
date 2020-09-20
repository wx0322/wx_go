package routers

import (
	"wx_go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/product", &controllers.ProductController{})
}
