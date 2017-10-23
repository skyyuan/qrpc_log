# qrpc_log

##部署
配置启动项目
```aidl
mkdir -p /etc/sv/qrpc_log

cat > /etc/sv/qrpc_log/run  << EOT
#!/bin/sh
export GO_ENV=production
cd  /srv/services/qrpc_log/src/qrpc_log
exec 2>&1
exec /srv/services/qrpc_log/src/qrpc_log/qrpc_log
EOT

chmod u+x /etc/sv/qrpc_log/run
```

创建日志
```aidl
mkdir -p /srv/services/qrpc_log/src/qrpc_log/logs
mkdir -p /etc/sv/qrpc_log/log
cat > /etc/sv/qrpc_log/log/run << EOT
#!/bin/sh
exec svlogd -tt /srv/services/qrpc_log/src/qrpc_log/logs
EOT
chmod u+x /etc/sv/caihongbaby/log/run
```

# 创建日志配置,
```aidl
cat > /srv/services/qrpc_log/src/qrpc_log/logs/config << EOT
s100000000
n5
EOT
```

上面的配置下，日志每99M滚动一次，保存5次滚动。详细的日志配置参考 svlogd

挂载服务

```
ln -s /etc/sv/qrpc_log /service/
```
在挂载后，服务会自动启动。所以在没有上传执行程序前，不要挂载服务。

启动和停止

```sv start|stop|restart|status qrpc_log
sv start|stop|restart|status qrpc_log/log
```
自启动
```
cat >>/etc/rc.local <<EOT
sv start qrpc_log
EOT
```