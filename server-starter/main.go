package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/end-date/server"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const SERVICE_PORT string = "3000"

func initServer() server.Server {
	username := os.Getenv("USERNAME")
	pass := os.Getenv("PASS")
	dbName := os.Getenv("DBNAME")
	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, pass, dbName)
	db, err := gorm.Open("postgres", str)
	if err != nil {
		panic("failed to connect database")
	}

	server := server.NewServer(db)
	return server
}

func listen(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{} // use default options
	fmt.Println("Connected user to the server on port: " + SERVICE_PORT)

	s := initServer()
	defer s.Close()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go server.UserWork(conn, s, wg)
	wg.Wait()
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/listen", listen)

	err := http.ListenAndServe(":"+SERVICE_PORT, nil)
	if err != nil {
		log.Fatalf("The server doesn't listen and serve: %s\n", err)
	}
}
