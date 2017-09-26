package routers

import (
	"qrpc_log/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.QlogsController{})
	ns := beego.NewNamespace("qrpc_log/",
		beego.NSNamespace("/qlogs",
			beego.NSInclude(
				&controllers.QlogsController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
