package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"openim/internal/common"
	"openim/internal/logger"
	"time"
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
	logger.Logger.Info("start ws server")
	http.HandleFunc("/ws", socketHandler)
	port := viper.GetString("wsPort")
	// port = ":6666"
	log.Println("/socket at", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("ListenAndServe出错：", err.Error())
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

	var connId string
	// todo in the future ,this id parameter could be change to token or key string,which is a encryption string of unique user.
	// connId should be get by an api, after user login,and has pair relation with userKey, for safety consideration

	// get parameters
	id := r.URL.Query().Get("id")
	if id != "" {
		log.Println("get parameter ws conn id:", id)
		connId = id
	} else {
		connId = common.MakeUuid()
	}
	Conns.Conn[connId] = conn
	log.Println("connid:", connId, "joined")

	// in the future could consider send a message to notice this user is online

	log.Printf("%+v\n", Conns.Conn)

	defer func() {
		log.Println("conn.Close()")
		delete(Conns.Conn, connId)
		conn.Close()
	}()

	// write go routine, used to send message
	go write()

	for {
		// get user send to server message
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Println("msgTYPE:", messageType)
		log.Printf("Received: %s\n", messageBytes)

		// todo 这里需要完善，ping的立即返回pong响应 ping9 pong10
		if messageType == websocket.PingMessage || messageType == websocket.PongMessage {
			continue
		}

		// 收到的消息通过server的writeChan发送出去
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

func write() {
	log.Println("go write process on")
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case msg := <-Conns.WriteChan:

			// json parse message
			var message Message
			err := json.Unmarshal(msg, &message)
			if err != nil {
				if string(msg) != "ping" {
					log.Println("Error during parse reading:", err)
				}
			} else {
				log.Println(message.UserName, "send to to all, content:", message.Content)
			}

			// 这里是特殊逻辑，用户发的消息，群发给目前连接上来的所有人，后面可以改成指定组和指定人
			for k, conn := range Conns.Conn {
				log.Println("send to", k, "msg:", string(msg))

				msgType := websocket.TextMessage
				if string(msg) == "ping" {
					msgType = websocket.PingMessage
				}
				err := conn.WriteMessage(msgType, msg)
				if err != nil {
					log.Println("write to", k, " msg error,msg:", string(msg), " err:", err.Error())
				}
			}
		case <-ticker.C:
			Conns.WriteChan <- []byte("ping")
			// for _, conn := range Conns.Conn {
			// conn.SetWriteDeadline(time.Now().Add(25 * time.Second))
			// if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			// 	log.Println("send ping error to", k)
			// 	log.Println("delete it from conns")
			// 	delete(Conns.Conn, k)
			// 	return
			// }
			// }

		}
	}
}

func GetOnlineUserNumber() int {
	return len(Conns.Conn)
}
