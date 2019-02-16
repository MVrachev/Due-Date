package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/end-date/components"
	"github.com/gorilla/websocket"
)

func defineOperation(conn *websocket.Conn, operation string) {
	in := bufio.NewReader(os.Stdin)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(operation)); err != nil {
		panic(err)
	}
	switch operation {
	case "add":
		add(conn, in)
	case "list by date":
		list(conn)
	case "list by priority":
		list(conn)
	case "finish":
		sendID(conn, in)
	default:
		fmt.Println("Unrecognized command!")
	}

}

// ------------------------------------- Add -------------------------------------

func add(conn *websocket.Conn, in *bufio.Reader) {
	fmt.Println("You will add a new task.")
	fmt.Print("Description: ")
	description, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print("Priority with a number: ")
	priority, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Due date day: ")
	dueDateDate, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Due date month: ")
	dueDateMonth, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Due date year: ")
	dueDateYear, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	info := components.Information{
		Priority:    Trim(priority),
		Description: Trim(description),
		Year:        Trim(dueDateYear),
		Month:       Trim(dueDateMonth),
		Day:         Trim(dueDateDate),
	}
	if err := conn.WriteJSON(&info); err != nil {
		panic(err)
	}
}

func Trim(str string) string {
	return strings.Trim(str, "\n")
}

// ------------------------------------- Lists -------------------------------------

func list(conn *websocket.Conn) {
	info := components.InfoForTasks{}
	if err := conn.ReadJSON(&info); err != nil {
		panic(err)
	}
	for _, task := range info.InfoTasks {
		fmt.Println(task)
	}
}

// ------------------------------------- Updates -------------------------------------

func sendID(conn *websocket.Conn, in *bufio.Reader) {
	_, idMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(idMsg)
	info, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	myID := Trim(info)
	if err := conn.WriteMessage(websocket.TextMessage,
		[]byte(myID)); err != nil {
		panic(err)
	}
	_, resMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(resMsg)
}
