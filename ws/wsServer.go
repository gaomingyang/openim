package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"openim/common"
)

// var upgrader = websocket.Upgrader{} // use default options
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWsServer() {
	http.HandleFunc("/ws", socketHandler)
	port := viper.GetString("wsPort")
	// port = ":6666"
	log.Println("/socket at", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("ListenAndServe出错：", err.Error())
	}
}

type Message struct {
	// UserId   int    `json:"id"`
	// TargetType string `target_type`
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

// 全局公共的
type ConnMap struct {
	Conn      map[string]*websocket.Conn
	WriteChan chan []byte
}

var Conns ConnMap

func init() {
	Conns = ConnMap{
		Conn:      make(map[string]*websocket.Conn),
		WriteChan: make(chan []byte, 500),
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	connId := common.MakeUuid()
	Conns.Conn[connId] = conn
	log.Println("connid:", connId, "joined")

	log.Printf("%+v\n", Conns)

	defer func() {
		log.Println("conn.Close()")
		conn.Close()
	}()

	// write go routine, used to send message
	go Write()

	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			delete(Conns.Conn, connId)
			break
		}
		log.Println("msgTYPE:", messageType)
		log.Printf("Received: %s\n", messageBytes)

		// json parse message
		// var msg Message
		// err = json.Unmarshal(messageBytes, &msg)
		// if err != nil {
		// 	log.Println("Error during parse reading:", err)
		// 	continue
		// }

		Conns.WriteChan <- messageBytes

		// resp := fmt.Sprintf("hi %s", message)
		// 向当前这个连接写入
		// err = conn.WriteMessage(messageType, messageBytes)
		// if err != nil {
		// 	log.Println("Error during message writing:", err)
		// 	break
		// }
	}
}

func Write() {
	log.Println("go write process on")
	for {
		select {
		case msg := <-Conns.WriteChan:
			log.Println("channel get ", msg)
			for k, conn := range Conns.Conn {
				fmt.Println("send to", k, "msg:", string(msg))
				err := conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("write msg error", string(msg))
				}
			}

		}
	}
}
