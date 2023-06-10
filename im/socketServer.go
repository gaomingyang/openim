package im

import (
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func StartSocketServer() {
	hub := NewHub()
	go hub.run() // 死循环
	// 接收连接请求
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Println("ws port:", viper.GetString("wsPort"))
	err := http.ListenAndServe(viper.GetString("wsPort"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// var header http.Header {
	// 	CheckOrigin: func() bool {
	// 		return true
	// 	}
	// }
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 每一个连接一个client实例
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	go client.writeProgress()
	go client.readProgress()
}
