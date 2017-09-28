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
                $tables.prepend(event.data)
            };

            setInterval(function(){
                time = $("table").find("tbody tr").first().find('td').last().text()

                $.get('/qrpc_log/qlogs/get_socket_time?time=' + time).done(function(data) {
                   if(data != null) {
                       for(i=data.length-1; i>=0; i--){
                          html = "<tr><td>" + data[i]['flag'] +  "</td><td>" + data[i]['type'] +  "</td><td>" + data[i]['content'] +  "</td><td>" + data[i]['level'] +  "</td><td>" + data[i]['time'] +  "</td></tr>"
                          ws.send(html);
                       }
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
                        {{ .time }}
                    </td>
                </tr>
            {{end}}
        </tbody>
    </table>
</div>
</body>
</html>