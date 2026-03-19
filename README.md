# Open IM（后端）

[![Build and Deploy Status](https://github.com/gaomingyang/openim/workflows/Build%20and%20Deploy%20To%20EC2/badge.svg)](https://github.com/gaomingyang/openim/actions)

开源即时通讯系统后端。可私有化部署以提升通信安全。

本仓库仅包含后端代码，前端已独立至单独仓库。

## 开发运行

```bash
go run main.go
```

或使用脚本：

```bash
./run.sh
```

## 部署

**编译并运行**

```bash
go build -o openim main.go
./openim
```

**交叉编译为 Linux 可执行文件**

```bash
make build
# 或
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openim main.go
```

## API

- `http://127.0.0.1:8888/api/register` 注册
- `http://127.0.0.1:8888/api/login` 登录
- `ws://127.0.0.1:8888/ws` WebSocket

更多接口见代码与 `docs/`。
