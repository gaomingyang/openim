package main

import (
	"flag"
	"log"
	"openim/config"
	"openim/dao"
	"openim/router"
	"openim/ws"
	"os"

	"github.com/spf13/viper"
)

// var http_addr = flag.String("http_addr", ":8001", "http service address")

func init() {
	log.SetPrefix("[LOG] ")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	env := flag.String("env", "", "environment")
	flag.Parse()
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(workPath, *env)
	dao.InitDB()
}

func main() {

	// 另一种方式，从命令行参数获取
	// PORT := ":8001"
	// arguments := os.Args
	// if len(arguments) != 1 {
	// 	PORT = ":" + arguments[1]
	// }

	// 开启sockert server服务
	// go im.StartSocketServer()
	go ws.StartWsServer()

	r := router.SetupRouter()
	log.Fatal(r.Run(viper.GetString("apiPort"))) // listen and serve on 0.0.0.0:8080
}
