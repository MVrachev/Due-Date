package main

import (
	"flag"
	"log"
	"net/url"
	"sync"

	"github.com/end-date/user"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/listen"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	user.LoginOrRegister(conn)

	var wg sync.WaitGroup
	wg.Add(1)
	go user.Work(conn, wg)
	wg.Wait()
}
