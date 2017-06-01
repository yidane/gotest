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
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = err.(error).Error()
			c.ServeJSON()
		}
	}()

	r := &grid.ReportResult{}
	d, s := models.GetDefine()
	c.Data["json"] = r.Success(d, s)
	c.ServeJSON()
}
