package user

import (
	"bufio"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func (u *User) LoginOrRegister(conn *websocket.Conn) {
	_, logOrRegMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	log.Println(string(logOrRegMsg))
	in := bufio.NewReader(os.Stdin)

	for {
		ans, _, err := in.ReadLine()
		//ans, err := in.ReadString('\n')

		finalAnsw := strings.Trim(string(ans), "\n")
		if err != nil {
			panic(err)
		}
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
			u.login(conn)
		} else if strRes == "You chose register." {
			u.registerUser(conn)
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

func (u *User) registerUser(conn *websocket.Conn) {
	u.sendUserName(conn)
	u.sendPass(conn)
}

// ------------------------------------- Login -------------------------------------

func (u *User) isUserInDataBase(conn *websocket.Conn, pass []byte) {
	_, possHash, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	log.Println(string(possHash))
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

func (u *User) login(conn *websocket.Conn) {
	for {
		u.sendUserName(conn)
		pass := u.sendPass(conn)
		u.isUserInDataBase(conn, pass)

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

func (u *User) sendUserName(conn *websocket.Conn) {
	_, userNameMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}

	log.Print(string(userNameMsg))
	in := bufio.NewReader(os.Stdin)

	//userName, err := in.ReadString('\n')
	//in.ReadLine
	userName, _, err := in.ReadLine()
	finalUserName := strings.Trim(string(userName), "\n")

	if err != nil {
		panic(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(finalUserName)); err != nil {
		panic(err)
	}
}

func (u *User) sendPass(conn *websocket.Conn) []byte {
	_, passMsg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	log.Print(string(passMsg))
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	finalPass := strings.Trim(string(bytePassword), "\n")

	//log.Println()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(finalPass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(hashedPassword)); err != nil {
		panic(err)
	}
	return bytePassword
}
