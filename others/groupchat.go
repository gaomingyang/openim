// 下面是一个参考实现的简单的群聊服务器，使用Go语言编写：

// 代码文件：chat_server.go
package main

import (
	"fmt"
	"net"
	"strings"
)

// 定义一个全局的映射表，
// 存储每个活动的连接以及大家的昵称
var activeConnections = make(map[net.Conn]string)

func main() {
	fmt.Println("Starting chat server...")
	listener, err := net.Listen("tcp", "8888")
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			break
		}
		go handleConnection(conn)
	}
}

// 处理新接入的客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()
	// 处理客户端设置昵称
	handleClient(conn)
	// 等待客户端发送消息
	for {
		message := make([]byte, 4096)
		length, err := conn.Read(message)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			// 移除映射表中 ，断开连接的客户端
			delete(activeConnections, conn)
			break
		}
		// 群发消息
		for conn, name := range activeConnections {
			if conn != conn {
				sendMessage(message[:length], conn, name)
			}
		}
	}
}

// 等待客户端设置昵称
func handleClient(conn net.Conn) {
	// 获取客户端发送的昵称
	message := make([]byte, 4096)
	length, err := conn.Read(message)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	nickName := string(message[length])
	// 把 昵称 和 连接保存到全局映射表
	activeConnections[conn] = nickName

	// 广播用户加入
	for conn, name := range activeConnections {
		if conn != conn {
			broadcastMessage(name+" join the chat\n", conn)
		}
	}
}

// 群发消息
func sendMessage(message []byte, conn net.Conn, name string) {
	// 保存发送者的昵称
	msgStr := name + " " + string(message)
	msg := []byte(msgStr)
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		return
	}
}

// 广播消息
func broadcastMessage(message string, conn net.Conn) {
	for conn, _ := range activeConnections {
		if conn != conn {
			msg := []byte(strings.TrimSpace(message) + "\n")
			_, err := conn.Write(msg)
			if err != nil {
				fmt.Println("Error broadcasting message", err.Error())
				return
			}
		}
	}
}
