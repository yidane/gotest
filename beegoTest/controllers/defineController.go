package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yidane/gotest/beegoTest/models/grid"
	"github.com/yidane/gotest/beegoTest/models/grid/data"
	"github.com/yidane/gotest/beegoTest/models/grid/style"
)

type DefineController struct {
	beego.Controller
}

func (c DefineController) Get() {
	c.EnableRender = false
	r := &grid.ReportResult{}
	d := data.Define{}
	s := style.Grid{}
	c.Data["json"] = r.Success(d, s)

	c.ServeJSON()
}
