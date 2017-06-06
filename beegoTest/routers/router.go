package routers

import (
	"github.com/astaxie/beego"
	"github.com/yidane/gotest/beegoTest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/User", &controllers.UserController{})
	beego.Router("/Define", &controllers.DefineController{})
	beego.Router("/Sale", &controllers.SaleController{})
	beego.Router("/Sale/GetDefine", &controllers.SaleController{}, "Get:GetDefine")
}
