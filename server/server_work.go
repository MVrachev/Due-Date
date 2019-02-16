package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func (s *Server) defineOperation(conn *websocket.Conn, operation string, userName string) {
	fmt.Print("operation is")
	fmt.Print(operation)
	switch operation {
	case "add":
		s.AddTask(conn, userName)
	case "list by date":
		s.ListTasksByDueDate(conn, userName)
	case "list by priority":
		s.ListTasksByPriority(conn, userName)
	}
}

// UserWork logs the user and waits for new operation
func UserWork(conn *websocket.Conn, server Server, wg sync.WaitGroup) {
	//user := server.LoginOrRegister(conn)
	defer conn.Close()
	user := server.LoginOrRegister(conn)
	for {
		fmt.Println("Waiting for a new message")
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		server.defineOperation(conn, string(message), user.Name)
		log.Printf("recv: %s", message)
		if string(message) == "bye" {
			log.Println("The client left!")
			break
		}
	}
	wg.Done()

}
