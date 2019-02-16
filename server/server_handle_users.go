package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func (s *Server) defineOperation(conn *websocket.Conn, operation string, userName string, done chan struct{}) {
	fmt.Print("operation is")
	fmt.Println(operation)
	switch operation {
	case "add":
		s.AddTask(conn, userName)

	}
	//case ""
	fmt.Println("Finished with operation")
	done <- struct{}{}
}

// UserWork logs the user and waits for new operation
func UserWork(conn *websocket.Conn, server Server, wg sync.WaitGroup) {
	//user := server.LoginOrRegister(conn)
	done := make(chan struct{})
	defer conn.Close()
	user := server.LoginOrRegister(conn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		server.defineOperation(conn, string(message), user.Name, done)
		log.Printf("recv: %s", message)
		<-done
		if string(message) == "bye" {
			log.Println("The client left!")
			break
		}
	}
	wg.Done()

}
