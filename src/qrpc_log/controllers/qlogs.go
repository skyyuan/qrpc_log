package controllers

import (
	"github.com/astaxie/beego"
)

type QlogsController struct {
	beego.Controller
}

// @router / [get]
func (c *QlogsController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
