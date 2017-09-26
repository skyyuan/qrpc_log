<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>

    <script>
      $(function() {
        setTimeout(function() {
            var ws = new WebSocket("ws://localhost:8100/echo");
            ws.onmessage = function(e) {
              $('<li>').text(event.data).appendTo($ul);
            };
            var $ul = $('#msg-list');
            $('#sendBtn').click(function(){

              var data = $('#name').val();
              ws.send(data);

            });
        }, 10)
      });
    </script>
</head>
<body>
<input id="name" type="text"/>
<input type="button" id="sendBtn" value="send"/>
<ul id="msg-list"></ul>
</body>
</html>