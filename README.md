

# Open IM
[![Build Status](https://github.com/gaomingyang/openim/workflows/Build%20and%20Deploy%20To%20Aws%20EC2/badge.svg?branch=master)](https://github.com/gaomingyang/openim/actions?query=branch%3Amaster)

An Opensource Instant Messaging System.  
You can deploy it privately to enhance communication security.

## Deploy

**Compile and run**
```
go mod init openim
go build -o openim main.go
./openim
```

**development run**
```
go run main.go
```

**Compile it into an executable file for a Linux server.**
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openim main.go
```

**APIs**
* http://127.0.0.1:8888/register
* http://127.0.0.1:8888/login
* ws://127.0.0.1:8888/ws