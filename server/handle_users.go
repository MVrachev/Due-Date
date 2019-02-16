package server

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func (s *Server) defineOperation(conn *websocket.Conn, operation string, userName string) {
	switch operation {
	case "add":
		s.AddTask(conn, userName)
	}
	//case ""
}

func UserWork(conn *websocket.Conn, server Server, wg sync.WaitGroup) {
	//user := server.LoginOrRegister(conn)
	defer conn.Close()
	for {

		user := server.LoginOrRegister(conn)
		for {
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
	}
	wg.Done()

}
