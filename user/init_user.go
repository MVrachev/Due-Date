package user

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func LoginOrRegister(conn *websocket.Conn) {
	_, logOrRegMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	log.Println(string(logOrRegMsg))
	in := bufio.NewReader(os.Stdin)

	for {
		ans, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}
		finalAnsw := strings.Trim(string(ans), "\n")

		if err := conn.WriteMessage(websocket.TextMessage, []byte(finalAnsw)); err != nil {
			panic(err)
		}
		_, resMessage, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		strRes := string(resMessage)
		log.Println(strRes)
		if strRes == "You chose login." {
			login(conn)
			return
		} else if strRes == "You chose register." {
			registerUser(conn)
			return
		} else {
			_, invalidCommMsg, err := conn.ReadMessage()
			if err != nil {
				panic(err)
			}
			log.Println(string(invalidCommMsg))
		}
	}
}

// ------------------------------------- Register -------------------------------------

func registerUser(conn *websocket.Conn) {
	userName := getUserName(conn)
	pass := getPass(conn)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := NewUser(userName, string(hashedPassword))
	if err := conn.WriteJSON(&user); err != nil {
		panic(err)
	}
}

// ------------------------------------- Login -------------------------------------

func isUserInDataBase(conn *websocket.Conn, pass []byte) {
	_, possHash, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(possHash, pass)
	if err != nil {
		if err = conn.WriteMessage(websocket.TextMessage, []byte("Not in DB")); err != nil {
			panic(err)
		}
	} else {
		if err = conn.WriteMessage(websocket.TextMessage, []byte("In DB")); err != nil {
			panic(err)
		}
	}

}

func login(conn *websocket.Conn) {
	for {
		userName := getUserName(conn)
		pass := getPass(conn)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(userName)); err != nil {
			panic(err)
		}
		isUserInDataBase(conn, pass)

		_, resMessage, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		strRes := string(resMessage)
		log.Println(strRes)
		if strRes == "Logged in!" {
			break
		}
	}
}

// ------------------------------------- Helper function -------------------------------------

func getUserName(conn *websocket.Conn) string {
	fmt.Print("Username: ")

	in := bufio.NewReader(os.Stdin)
	userName, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	finalUserName := strings.Trim(string(userName), "\n")
	return finalUserName
}

func getPass(conn *websocket.Conn) []byte {
	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	log.Println()
	return bytePassword
}
