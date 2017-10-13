package controllers

import (
	"github.com/astaxie/beego"
	"qrpc_log/utils"
	"qrpc_log/models"
	"time"
	"strings"
	"fmt"
)

type QlogsController struct {
	beego.Controller
}

// @router / [get]
func (c *QlogsController) Get() {
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	logType := c.GetString("log_type")
	level := c.GetString("log_level")
	qlogs, _ := models.GetQlogsByParams(mdb,level,logType)
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
	c.Data["qlogs"] = results
	c.Data["log_type"] = logType
	c.Data["log_level"] = level
	c.TplName = "qlogs/index.tpl"
}

// @router /get_socket_time [get]
func (c *QlogsController) SocketTime() {
	str := c.GetString("time")
	logType := c.GetString("type")
	level := c.GetString("level")
	fmt.Println(logType)
	fmt.Println(level)
	str = strings.TrimSpace(str)
	t2, _ := time.Parse("2006-01-02 15:04:05", str)
	//fmt.Println(t2.UTC()) UTC时间
	cst,_ := time.LoadLocation("Local")
	//fmt.Println(t2.In(cst))// CST时间
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	qlogs, _ := models.GetQlogsByTime(mdb, t2.In(cst), level, logType)
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
	c.Data["qlogs"] = results
	c.TplName = "qlogs/index.tpl"
}
