package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

func StartWsServer() {
	http.HandleFunc("/", socketHandler)
	port := viper.GetString("wsPort")
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

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		fmt.Printf("Received: %s\n", messageBytes)

		// json parse message
		var msg Message
		err = json.Unmarshal(messageBytes, &msg)
		if err != nil {
			log.Println("Error during parse reading:", err)
			continue
		}

		// resp := fmt.Sprintf("hi %s", message)
		// 向当前这个连接写入
		err = conn.WriteMessage(messageType, messageBytes)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}
