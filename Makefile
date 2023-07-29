.PHONY: build

build:
	if [ ! -d log ]; then mkdir log; fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openim main.go
	@echo "build successfully!"