<!DOCTYPE html>
<html>
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
</style>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>

    <script>
        var dataa;
        $(function() {
            var ws = new WebSocket("ws://localhost:8100/echo");
            var $tables = $('#tbody');
            ws.onmessage = function(e) {
                $tables.append(event.data)
            };

            setInterval(function(){
                $.get('/qrpc_log/qlogs/get_socket_time').done(function(data) {
                   dataa = data
                   for(var i=0;i<data.length;i++){
                      html = "<tr><td>" + data[i]['BFlag'] +  "</td><td>" + data[i]['BType'] +  "</td><td>" + data[i]['Content'] +  "</td><td>" + data[i]['Level'] +  "</td><td>" + data[i]['created_at'] +  "</td></tr>"
                      ws.send(html);
                   }
                })
            }, 5000)
        });
    </script>
</head>
<body>
<ul id="msg-list"><input type="hidden" id="last_time" value="{{ .qlogs }}"></ul>
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
                        {{ .BFlag }}
                    </td>
                    <td>
                        {{ .BType }}
                    </td>
                    <td>
                        {{ .Content }}
                    </td>
                    <td>
                        {{ .Level }}
                    </td>
                    <td>
                        {{ .CreatedAt }}
                    </td>
                </tr>
            {{end}}
        </tbody>
    </table>
</div>
</body>
</html>