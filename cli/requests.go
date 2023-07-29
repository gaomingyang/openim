package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"openim/services"
	"openim/ws"
)

var (
	loginAPI    = "http://10.44.22.19:8888/login"
	registerAPI = "http://10.44.22.19:8888/register"
	wsAPI       = "ws://10.44.22.19:8899/ws"
)

func regiserRequest(username string, password string) {
	req := services.RegisterRequest{
		UserName: username,
		Password: password,
	}
	data, err := json.Marshal(req)
	errorCheck(err)

	resp, err := http.Post(registerAPI, "application/json", bytes.NewReader(data))
	errorCheck(err)
	respCheck(resp)
}

func loginRequest(username string, password string) {
	req := services.LoginRequest{
		UserName: username,
		Password: password,
	}
	data, err := json.Marshal(req)
	errorCheck(err)

	resp, err := http.Post(loginAPI, "application/json", bytes.NewReader(data))
	errorCheck(err)
	respCheck(resp)
}

func websocketConnection() (chan<- ws.Message, <-chan ws.Message) {
	conn, _, err := websocket.DefaultDialer.Dial(wsAPI, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}

	return sendMessages(conn), receiveMessages(conn)
}

func sendMessages(conn *websocket.Conn) chan<- ws.Message {
	ch := make(chan ws.Message)
	go func() {
		for msg := range ch {
			message := fmt.Sprintf(`{"user_name":"%s","content":"%s"}`, msg.UserName, msg.Content)
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Error sending message:", err)
				return
			}

		}
	}()
	return ch
}

func receiveMessages(conn *websocket.Conn) <-chan ws.Message {
	ch := make(chan ws.Message)
	go func() {
		for {
			// Read message from the server
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			var msg ws.Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				fmt.Printf("Error unmarshalling message %q: %v\n", message, err)
				continue
			}
			ch <- msg
		}
	}()
	return ch
}
