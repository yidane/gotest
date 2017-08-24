package routers

import (
	"github.com/yidane/weiPhoto/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/List",&controllers.MainController{},"GET:List")
}
