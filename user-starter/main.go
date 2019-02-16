package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/end-date/user"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/listen"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	user := user.NewUser("", "")
	user.LoginOrRegister(conn)

	done := make(chan struct{})
	chanIn := make(chan string)

	go func(chanIn chan string, done chan struct{}) {
		in := bufio.NewReader(os.Stdin)
		for {
			msg, _, err := in.ReadLine()
			if err != nil {
				log.Fatal("write: ", err)
			}
			if string(msg) == "bye" {
				done <- struct{}{}
				return
			}
			chanIn <- string(msg)
		}
	}(chanIn, done)

	for {
		select {
		case <-done:
			conn.WriteMessage(websocket.TextMessage, []byte("bye"))
			return
		case msg := <-chanIn:
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
