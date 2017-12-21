package controllers

import "github.com/astaxie/beego"

type ImageController struct {
	beego.Controller
}

func (c ImageController) Get() {

}

func (c ImageController) Thumb() {
	c.Ctx.ResponseWriter.Write(nil)
	c.Ctx.ResponseWriter.Header().Add("","")
}
