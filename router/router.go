package router

import (
	"fmt"
	"log"
	"net/http"
	"openim/services"
	"openim/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	corsMiddleWare := cors.Default()
	r.Use(corsMiddleWare)
	r.GET("/ping", pong)

	api := r.Group("/api")
	// apis
	api.POST("/register", services.UserRegister) // 用户注册
	// todo 检查邮箱是否重复的接口
	api.POST("/login", services.LoginHandler) //用户登录接口
	r.POST("/refreshToken", services.refreshTokenHandler)

	api.GET("/userinfo", services.UserInfoHandler) //需要通过token验证

	// group
	r.GET("/groups", services.OpenGroups)           // 查看所有开放的组列表
	r.POST("/group/apply", services.ApplyJoinGroup) // 申请入群
	r.GET("/my/group/list", services.MyGroupList)   // 查看自己的组列表
	r.GET("/group/info", services.GroupInfo)        // 查看某个组的信息
	r.GET("/group/members", services.GroupMembers)
	r.POST("/group/create", services.CreateGroup)
	r.POST("/group/join", services.JoinGroup)
	r.POST("/group/quit", services.QuitGroup)

	// friends
	// r.GET("/my/friends")

	r.LoadHTMLGlob("public/*")
	r.GET("/manage", managePage)
	r.GET("/chat", chatPage)
	r.GET("/login", loginPage)

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
	return r
}

func pong(c *gin.Context) {
	// c.String(200, "pong")
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

// todo this page could use ws update number instead of refresh webpage by hand
func managePage(c *gin.Context) {
	onlineUserNumber := ws.GetOnlineUserNumber()
	log.Println("online UserNumber:", onlineUserNumber)
	c.HTML(http.StatusOK, "manage.html", gin.H{"onlineUserNumber": onlineUserNumber})
}

func chatPage(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", nil)
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
