package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"openim/internal/handlers"
	"openim/internal/middleware"
	"openim/internal/services"
	"openim/internal/ws"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	corsMiddleWare := cors.Default()

	// 通用中间件
	r.Use(corsMiddleWare)
	r.Use(middleware.RequestIDMiddleware())

	r.GET("/ping", pong)

	api := r.Group("/api")
	api.POST("/register", services.UserRegister) // 用户注册
	// todo 检查邮箱是否重复的接口
	api.POST("/login", handlers.LoginHandler) // 用户登录接口

	// 需要验证token的
	authApi := api.Use(middleware.AuthMiddleware())
	authApi.POST("/refreshToken", services.RefreshTokenHandler)
	authApi.GET("/userinfo", handlers.UserInfoHandler) // 需要通过token验证

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

	r.LoadHTMLGlob("web/templates/*.html")
	r.GET("/manage", managePage)
	r.GET("/chat", chatPage)
	r.GET("/login", loginPage)
	r.GET("/home", homePage)
	r.GET("/", indexPage)

	// log.Println("Start web server with port number", *http_addr)
	// http.HandleFunc("/version", version)
	// // http.Handle("/", http.FileServer(http.Dir("./web")))  //根目录指向public
	// http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
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
func LogrusLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()

		// 记录请求日志
		log.WithFields(logrus.Fields{
			"request_id": c.MustGet("request_id"),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"duration":   end.Sub(start),
		}).Info("Request processed")
	}
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
	http.ServeFile(w, r, "./web/home.html")
}

func homePage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", nil)
}

func indexPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
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
