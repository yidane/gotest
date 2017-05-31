package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yidane/gotest/beegoTest/models"
	"github.com/yidane/gotest/beegoTest/models/grid"
)

type SaleController struct {
	beego.Controller
}

func (c SaleController) Get() {
	r := &grid.ReportResult{}
	d, s := models.GetDefine()
	c.Data["json"] = r.Success(d, s)
	c.ServeJSON()
}
