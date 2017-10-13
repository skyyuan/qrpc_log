package socket

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
	"net/url"
	"strings"
	"time"
	"qrpc_log/utils"
	"qrpc_log/models"
	"encoding/json"
)

func EchoHandler(ws *websocket.Conn) {
	msg := make([]byte, 32*1024)
	for {
		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			return
		}
		queryString := string(msg[:n])
		fmt.Printf("Read: ", queryString)

		if n > 0 {
			params,_ := url.ParseQuery(queryString)
			str := params.Get("time")
			logType := params.Get("type")
			level := params.Get("level")
			fmt.Println("type:%s", logType)
			fmt.Println("level:%s", level)
			str = strings.TrimSpace(str)
			cst,_ := time.LoadLocation("Local")
			t2, _ := time.ParseInLocation("2006-01-02 15:04:05.999999999", str,cst)
			//fmt.Println(t2.UTC()) UTC时间

			//fmt.Println(t2.In(cst))// CST时间
			mdb, mSession := utils.GetMgoDbSession()
			defer mSession.Close()
			fmt.Println(t2)
			qlogs, _ := models.GetQlogsByTime(mdb, t2, level, logType)
			fmt.Println(len(qlogs))
			var results []map[string]interface{}
			for _, q := range qlogs {
				timestamp := q.CreatedAt
				result := map[string]interface{}{
					"type": q.BType,
					"flag": q.BFlag,
					"level": q.Level,
					"content": q.Content,
					"time":  timestamp.Format("2006-01-02 15:04:05"),
					"correct_time": timestamp.Format("2006-01-02 15:04:05.999999999"),
				}
				results = append(results, result)
			}
			data, _ := json.Marshal(results)
			_, err := ws.Write(data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("Write:", data)
		}
	}
}
