package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["qrpc_log/controllers:QlogsController"] = append(beego.GlobalControllerRouter["qrpc_log/controllers:QlogsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qrpc_log/controllers:QlogsController"] = append(beego.GlobalControllerRouter["qrpc_log/controllers:QlogsController"],
		beego.ControllerComments{
			Method: "SocketTime",
			Router: `/get_socket_time`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
