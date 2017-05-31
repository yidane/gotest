package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yidane/gotest/beegoTest/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	user := models.New()
	c.Data["ID"] = user.ID
	c.Data["Name"] = user.Name
	c.Data["Gender"] = user.Gender
	c.Data["Age"] = user.Age
	c.Data["Time"] = user.Time
	c.TplName = "user.html"
}
