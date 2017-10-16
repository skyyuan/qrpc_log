<!DOCTYPE html>
<html>
<link rel="stylesheet" href="/qrpc_log/bootstrap/css/bootstrap.min.css"/>
<style>
    .table {
        background: #FFF;
        font-size: 12px;
        border-top: 1px solid #e1e6eb;
        margin-top: 8px;
        border: 1px solid #e1e6eb;
    }
    .table {
        width: 100%;
        margin-bottom: 20px;
    }
    table {
        border-collapse: collapse;
        border-spacing: 0;
    }
    tr {
        display: table-row;
        vertical-align: inherit;
        border-color: inherit;
    }
    thead {
        display: table-header-group;
        vertical-align: middle;
        border-color: inherit;
    }
    tbody {
        display: table-row-group;
        vertical-align: middle;
        border-color: inherit;
    }
    .table thead tr th {
        padding: 8px 8px;
        font-weight: normal;
        color: #999;
        border-bottom: 1px solid #e1e6eb;
        background-color: #F5F6FA;
    }
    .table tbody tr td {
        padding: 12px 8px;
        border-top: 0px;
        border-bottom: 1px solid #e1e6eb;
        text-align: -internal-center;

    }
    td{
        text-align: -webkit-center;
    }
    th {
            text-align: -webkit-center;
    }
    .margin-top, .margin-top-1 {
        margin-top: 8px !important;
        text-align: -webkit-center;
    }
</style>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="/qrpc_log/js/jquery-2.1.4.min.js"></script>
    {{if .showNewest }}
    <script>
        var dataa;
        $(function() {
            var ws = new WebSocket("ws://lws.qychbb.com/echo");
            var $tables = $('#tbody');
            ws.onmessage = function(e) {
                html = ""
                if(event.data != "null"){
                    results = JSON.parse(event.data)
                    $.each(results, function (i, item) {
                        html = html + "<tr><td>" + item['flag'] +  "</td><td>" + item['type'] +  "</td><td>" + item['content'] +  "</td><td>" + item['level'] + "</td><td>" + item['trace_id'] +  "</td><td data-time='" + item['correct_time'] +  "'>" + item['time'] +  "</td></tr>"
                    })
                    $tables.prepend(html)
                }
            };

            setInterval(function(){
                time = $("table").find("tbody tr").first().find('td').last().data("time")
                log_level = $("#log_level").val()
                log_type = $("#log_type").val()
                trace_id =  $("#trace_id").val()
                content = $("#content").val()
                var params = 'time=' + time + '&level=' + log_level + '&type=' + log_type + '&trace_id=' + trace_id+ '&content=' + content;
                ws.send(params);
            }, 5000)
        });
    </script>
   {{end }}
</head>
<body>
<form action="/qrpc_log/qlogs" class="form-inline ng-pristine ng-valid margin-top" role="form">
    <div class="form-group">
        <label class="control-label">日志类型：</label>
        <input type="text" class="form-control ng-pristine ng-valid" size="15" id="log_type" name="log_type" value="{{ .log_type }}"/>
    </div>
    <div class="form-group">
        <label class="control-label">日志级别：</label>
        <select id="log_level" name="log_level"  class="form-control">
            <option value="info" {{if eq .log_level "info"}} selected{{end}}>info</option>
            <option value="error" {{if eq .log_level "error"}} selected{{end}}>error</option>
            <option value="fatal" {{if eq .log_level "fatal"}} selected{{end}}>fatal</option>
        </select>
    </div>
    <div class="form-group">
        <label class="control-label">traceid：</label>
        <input type="text" class="form-control ng-pristine ng-valid" size="15" id="trace_id" name="trace_id" value="{{ .trace_id }}"/>
    </div>
    <div class="form-group">
        <label class="control-label">详情：</label>
        <input type="text" class="form-control ng-pristine ng-valid" size="15" id="content" name="content" value="{{ .content }}"/>
    </div>
    <div class="form-group">
        <button type="submit" class="btn btn-default">搜索</button>
    </div>

</form>
<div class="gridSection">
    <table class="table table-hover">
        <thead>
            <tr>
                <th>BFlag</th>
                <th>日志类型</th>
                <th>日志详情</th>
                <th>日志级别</th>
                <th>traceid</th>
                <th>操作日期</th>
            </tr>
        </thead>
        <tbody id="tbody">
            {{range .qlogs}}
                <tr>
                    <td>
                        {{ .flag }}
                    </td>
                    <td>
                        {{ .type }}
                    </td>
                    <td>
                        {{ .content }}
                    </td>
                    <td>
                        {{ .level }}
                    </td>
                    <td>
                        {{if ne .trace_id ""}}
                        <a href="/qrpc_log/qlogs?trace_id={{ .trace_id }}">{{ .trace_id }}</a>
                        {{end}}
                    </td>
                    <td data-time="{{ .correct_time }}">
                        {{ .time }}
                    </td>
                </tr>
            {{end}}
        </tbody>
    </table>
</div>
{{if .paginator.HasPages}}
<ul class="pagination pagination">
    {{if .paginator.HasPrev}}
        <li><a href="{{.paginator.PageLinkFirst}}">首页</a></li>
        <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
    {{else}}
        <li class="disabled"><a>首页</a></li>
        <li class="disabled"><a>&laquo;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
    <li{{if $.paginator.IsActive .}} class="active"{{end}}>
    <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
    </li>
{{end}}
{{if .paginator.HasNext}}
    <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
    <li><a href="{{.paginator.PageLinkLast}}">最后一页</a></li>
{{else}}
    <li class="disabled"><a>&raquo;</a></li>
    <li class="disabled"><a>最后一页</a></li>
{{end}}
</ul>
{{end}}
</body>
</html>