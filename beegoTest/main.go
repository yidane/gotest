package main

import (
	"github.com/astaxie/beego"
	beegorm "github.com/yidane/gotest/beegoTest/models/gorm"
	_ "github.com/yidane/gotest/beegoTest/routers"
)

func main() {
	beegorm.Init()
	beego.Run()
}
