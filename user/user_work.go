package user

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

func Work(conn *websocket.Conn, done chan struct{}, wg sync.WaitGroup) {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You have to give an operation: ")
		msg, _, err := in.ReadLine()
		if err != nil {
			panic(err)
		}
		realMsg := Trim(string(msg))
		if realMsg == "bye" {
			done <- struct{}{}
			return
		}

		defineOperation(conn, realMsg)
	}
	wg.Wait()
}
