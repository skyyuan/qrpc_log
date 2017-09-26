package socket

import (
	"golang.org/x/net/websocket"
	"time"
	"fmt"
)

type Client struct {
	ws     *websocket.Conn
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{ws}
}

func  (c *Client)Start() (err error) {
	timeout := time.After(3 * time.Second)
	fmt.Println("进去client")
	for {
		fmt.Println("循环client")
		select {
		case <-timeout:
			if err = websocket.Message.Send(c.ws, "ping"); err != nil {
				fmt.Println("关闭client")
				c.ws.Close();
				return err
			}
		//case msg := <-c.ch:
		//	log.Println("Send:", msg)
		//	websocket.JSON.Send(c.ws, msg)
		}
	}
}