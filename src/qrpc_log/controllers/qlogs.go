package controllers

import (
	"github.com/astaxie/beego"
	"qrpc_log/utils"
	"qrpc_log/models"
)

type QlogsController struct {
	beego.Controller
}

// @router / [get]
func (c *QlogsController) Get() {
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	qlogs, _ := models.GetQlogs(mdb)
	c.Data["qlogs"] = qlogs
	c.TplName = "index.tpl"
}

// @router /get_socket_time [get]
func (c *QlogsController) SocketTime() {
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	qlogs, _ := models.GetQlogs(mdb)

	c.Data["json"] = qlogs
	c.ServeJSON()
	return
}
