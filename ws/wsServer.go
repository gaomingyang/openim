package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"openim/common"
)

var upgrader = websocket.Upgrader{} // use default options

func StartWsServer() {
	http.HandleFunc("/socket", socketHandler)
	port := viper.GetString("wsPort")
	port = ":6666"
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
		Conn:      nil,
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

	log.Printf("%+v\n", Conns)

	defer func() {
		log.Println("conn.Close()")
		conn.Close()
	}()

	// write go routine
	go Write()

	// The event loop
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			// delete(connMap.Conn, conn)
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
			for _, conn := range Conns.Conn {
				err := conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("write msg error", string(msg))
				}
				log.Println("send ")
			}

		}
	}
}
