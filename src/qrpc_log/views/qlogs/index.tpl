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
    <script>
        var dataa;
        $(function() {
            var ws = new WebSocket("ws://127.0.0.1:8100/echo");
            var $tables = $('#tbody');
            ws.onmessage = function(e) {
                html = ""
                if(event.data != "null"){
                    results = JSON.parse(event.data)
                    $.each(results, function (i, item) {
                        html = html + "<tr><td>" + item['flag'] +  "</td><td>" + item['type'] +  "</td><td>" + item['content'] +  "</td><td>" + item['level'] +  "</td><td data-time='" + item['correct_time'] +  "'>" + item['time'] +  "</td></tr>"
                    })
                    $tables.prepend(html)
                }
            };

            setInterval(function(){
                time = $("table").find("tbody tr").first().find('td').last().data("time")
                log_level = $("#log_level").val()
                log_type = $("#log_type").val()
                var params = 'time=' + time + '&level=' + log_level + '&type=' + log_type;
                ws.send(params);
            }, 5000)
        });
    </script>

</head>
<body>
<form action="/qrpc_log/qlogs" class="form-inline ng-pristine ng-valid margin-top" role="form">
    <div class="form-group">
        <label class="control-label">日志类型：</label>
        <input type="text" class="form-control ng-pristine ng-valid" size="15" id="log_type" name="log_type"/>
    </div>
    <div class="form-group">
        <label class="control-label">日志级别：</label>
        <input type="text" class="form-control ng-pristine ng-valid" size="15" id="log_level" name="log_level"/>
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
                    <td data-time="{{ .correct_time }}">
                        {{ .time }}
                    </td>
                </tr>
            {{end}}
        </tbody>
    </table>
</div>
</body>
</html>