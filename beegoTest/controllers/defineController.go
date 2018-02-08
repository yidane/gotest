package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yidane/gotest/beegoTest/models"
	"github.com/yidane/gotest/beegoTest/models/grid"
)

type DefineController struct {
	beego.Controller
}

func (c DefineController) Get() {
	c.EnableRender = false
	r := &grid.ReportResult{}
	// d := data.Define{}
	// s := style.Grid{}
	d, s := models.GetDefine()
	c.Data["json"] = r.Success(d, s)

	c.ServeJSON()
}
