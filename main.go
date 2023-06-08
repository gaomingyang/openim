package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openim/config"
	"openim/dao"
	"openim/im"
	"openim/services"
	"os"
)

var http_addr = flag.String("http_addr", ":8001", "http service address")

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
	go im.StartSocketServer()

	r := gin.Default()
	corsMiddleWare := cors.Default()
	r.Use(corsMiddleWare)
	r.GET("/ping", pong)
	r.POST("/register", services.UserRegister)
	r.POST("/login", services.UserLogin)
	r.GET("/", indexPage)
	r.Run(":8888") // listen and serve on 0.0.0.0:8080

	// log.Println("Start web server with port number", *http_addr)
	// http.HandleFunc("/version", version)
	// // http.Handle("/", http.FileServer(http.Dir("./public")))  //根目录指向public
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	//
	// // 开启web 聊天页面
	// http.HandleFunc("/", home)

	// err := http.ListenAndServe(PORT, nil)
	// flag获取的addr是一个指针，对指针使用*进行取值
	// err := http.ListenAndServe(*http_addr, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.0.1")
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving", r.Host, "for", r.URL.Path)
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./public/home.html")
}

func indexPage(c *gin.Context) {
	// c.String(200, "Success")
	c.HTML(http.StatusOK, "/public/home.html", nil)
}
