package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) List() {
	fmt.Println(c.Ctx.Input.Params())
	category := c.GetString("c")
	fmt.Println(category)
	lastImgName := c.GetString("filename")
	fmt.Println(lastImgName)
	c.Data["json"] = "[" + category + lastImgName + "]"
	c.ServeJSON()
}
