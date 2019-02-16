package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const PORT = 3540

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: " + err.String())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}

func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=mvrachev sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//db.CreateTable(&components.Task{})

	//db.DropTableIfExists(&components.Elem{})

	//db.CreateTable(&components.Elem{})

	// server := server.NewServer(db)

	//elements := make([]Elem, 0)

	// t := time.Date(2018, 11, 24, 13, 35, 21, 0, time.UTC)
	// task := components.NewTask("Stefan Yatanski", t, 10, "First Task!")

	// server.AddTask(task)
	// //server.UpdatePriority(task, 10)
	// //server.FinishTask(task)

	// // //elements = append(elements, elem)

	// secT := time.Date(2019, 05, 24, 13, 35, 21, 0, time.UTC)
	// secTask := components.NewTask("Stefan Yatanski", secT, 3, "Second Task")
	// server.Delete(secTask)

	// server.AddTask(secTask)
	// //elements = append(elements, secElem)

	// thirdT := time.Date(2020, 12, 24, 13, 35, 21, 0, time.UTC)
	// thirdTask := components.NewTask("Stefan Yatanski", thirdT, 2, "First task")
	// server.AddTask(thirdTask)

	//elems := server.ListElementsByPriority()

	//fmt.Println(elems)
	//print(elems)
	//elements = append(elements, thirdElem)

	server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil {
		panic("couldn't start listening: " + err.String())
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}

}
