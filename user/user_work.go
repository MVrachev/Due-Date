package user

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

func Work(conn *websocket.Conn, wg sync.WaitGroup) {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You have to give an operation: ")
		msg, _, err := in.ReadLine()
		if err != nil {
			wg.Done()
			panic(err)
		}
		realMsg := Trim(string(msg))
		if realMsg == "bye" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("bye")); err != nil {
				panic(err)
			}
			conn.Close()
			wg.Done()
			return
		}

		defineOperation(conn, realMsg)
	}
}
