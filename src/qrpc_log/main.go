package main

import (
	_"qrpc_log/routers"
	"github.com/astaxie/beego"
	"net/http"
	"golang.org/x/net/websocket"
	"qrpc_log/socket"
	"fmt"
)

func main() {
	go beego.Run()

	http.Handle("/echo", websocket.Handler(socket.EchoHandler))
	fmt.Println(1111)
	err := http.ListenAndServe(":8100", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

