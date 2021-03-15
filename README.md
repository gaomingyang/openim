# im
用go语言实现的im服务，基于websocket通信

## 部署

```
go mod init openim
go build -o openim main.go
./openim
```
测试可执行`go run main.go`

访问http://127.0.0.1:8001/可以在示例页面进行测试