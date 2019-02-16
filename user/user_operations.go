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
	switch operation {
	case "add":
		add(conn, in)
	}
}

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
		Priority:    trim(priority),
		Description: trim(description),
		Year:        trim(dueDateYear),
		Month:       trim(dueDateMonth),
		Day:         trim(dueDateDate),
	}
	if err := conn.WriteJSON(&info); err != nil {
		panic(err)
	}
}

func trim(str string) string {
	return strings.Trim(str, "\n")
}
