# im
用go语言实现的im服务，基于websocket通讯

## 部署

```
go mod init openim

```


开启web服务（为展示im页面）
```
go run webserver.go
```

开启socket服务
```
cd im
go run *.go
```