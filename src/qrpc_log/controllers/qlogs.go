package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
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
	traceId := c.GetString("trace_id")
	content := c.GetString("content")
	page, _ := c.GetInt("p")
	if page < 1 {
		page = 1
	}
	qlogs, _ := models.GetQlogsByParams(mdb,level,logType,traceId,content,page)
	var results []map[string]interface{}
	for _, q := range qlogs {
		timestamp := q.CreatedAt
		result := map[string]interface{}{
			"type": q.BType,
			"flag": q.BFlag,
			"level": q.Level,
			"content": q.Content,
			"trace_id":q.TraceId,
			"time":  timestamp.Format("2006-01-02 15:04:05"),
			"correct_time": timestamp.Format("2006-01-02 15:04:05.999999999"),
		}
		results = append(results, result)
	}
	c.Data["qlogs"] = results
	c.Data["log_type"] = logType
	c.Data["log_level"] = level
	c.Data["trace_id"] = traceId
	c.Data["content"] = content
	c.Data["showNewest"] = (page == 1)
	count, _ := models.GetAllQlogsCount(mdb,level,logType,traceId,content)
	maxCount := int64(count)
	perPage := 10
	paginator := pagination.SetPaginator(c.Ctx, perPage, maxCount)
	c.Data["paginator"] = paginator
	c.TplName = "qlogs/index.tpl"
}

// @router /get_socket_time [get]
func (c *QlogsController) SocketTime() {
	str := c.GetString("time")
	logType := c.GetString("type")
	level := c.GetString("level")
	traceId := c.GetString("trace_id")
	content := c.GetString("content")
	fmt.Println(logType)
	fmt.Println(level)
	str = strings.TrimSpace(str)
	t2, _ := time.Parse("2006-01-02 15:04:05", str)
	//fmt.Println(t2.UTC()) UTC时间
	cst,_ := time.LoadLocation("Local")
	//fmt.Println(t2.In(cst))// CST时间
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	qlogs, _ := models.GetQlogsByTime(mdb, t2.In(cst), level, logType,traceId,content)
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
