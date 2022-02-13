package routers

import (
	"beego-web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由
	beego.Router("/", &controllers.MainController{})
}
