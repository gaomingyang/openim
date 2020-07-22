package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8001", "http service address")

func main() {
	flag.Parse()

	// 另一种方式，从命令行参数获取
	// PORT := ":8001"
	// arguments := os.Args
	// if len(arguments) != 1 {
	// 	PORT = ":" + arguments[1]
	// }

	fmt.Println("Start web server with port number", *addr)
	http.HandleFunc("/test", test)
	// http.Handle("/", http.FileServer(http.Dir("./public")))  //根目录指向public
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))) // /public指向 /public
	http.HandleFunc("/", home)

	// err := http.ListenAndServe(PORT, nil)
	//flag获取的addr是一个指针，对指针使用*进行取值
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
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
