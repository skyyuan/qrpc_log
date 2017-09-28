package main

import (
	_"qrpc_log/routers"
	"github.com/astaxie/beego"
	"net/http"
	"golang.org/x/net/websocket"
	"qrpc_log/socket"
)

func main() {
	beego.SetStaticPath("/qrpc_log", "static")
	go beego.Run()
	go http.Handle("/echo", websocket.Handler(socket.EchoHandler))
	err := http.ListenAndServe(":8100", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}