# Open IM
An Opensource Instance Message System.

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