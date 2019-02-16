package server

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// func (s *Server) defineOperation(string operation, userName string) {
// 	switch operation {
// 	case "add":
// 		s.AddTask()
// 	}
// 	case ""
// }

func UserWork(conn *websocket.Conn, server Server) {
	//user := server.LoginOrRegister(conn)
	defer conn.Close()
	server.LoginOrRegister(conn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		//defineOperation(string(message), user.Name)
		log.Printf("recv: %s", message)
		if string(message) == "bye" {
			fmt.Println("The client left!")
			break
		}
	}
}
