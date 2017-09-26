package socket

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
)

func EchoHandler(ws *websocket.Conn) {
	msg := make([]byte, 32*1024)
	for {
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read: ", msg[:n])
		if n > 0 {
			send_msg := "[" + string(msg[0:n]) + "]"
			m, err := ws.Write([]byte(send_msg))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Write:", msg[:m])
		}
	}
}
