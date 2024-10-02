package main

import (
	"flag"
	"log"
	"openim/internal/common/config"
	"openim/internal/dao"
	"openim/internal/logger"
	"openim/internal/routes"
	"openim/internal/ws"
	"os"

	"github.com/spf13/viper"
)

// var http_addr = flag.String("http_addr", ":8001", "http service address")

func init() {
	log.SetPrefix("[LOG] ")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	// init config
	env := flag.String("env", "", "environment")
	flag.Parse()
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(workPath, *env)

	// inititial database connection
	dao.InitDB()
}

func main() {
	defer logger.Sync()

	logger.Logger.Info("Application started")

	// 另一种方式，从命令行参数获取
	// PORT := ":8001"
	// arguments := os.Args
	// if len(arguments) != 1 {
	// 	PORT = ":" + arguments[1]
	// }

	// 开启socket server服务
	// go im.StartSocketServer()
	go ws.StartWsServer()

	//go testLog()

	r := routes.SetupRouter()
	log.Fatal(r.Run(viper.GetString("apiPort"))) // listen and serve on 0.0.0.0:8080
}

// func testLog() {
// 	for i := 0; i < 10000; i++ {
// 		logger.Logger.Info("testLogASDFASDFASDASDFASDFASDFASDFASDFASDFDDDDDDDDDDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASDFASFDSASDFASDFASDAFDA")
// 	}
// }
