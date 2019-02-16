package server

import (
	"fmt"
	"log"

	"github.com/end-date/user"
	"github.com/gorilla/websocket"
)

// LoginOrRegister is the server function which
// initializes the user
func (s *Server) LoginOrRegister(conn *websocket.Conn) user.User {
	if err := conn.WriteMessage(websocket.TextMessage,
		[]byte("Login or Register? Write l for login or r for register")); err != nil {
		panic(err)
	}

	for {
		_, answer, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		strAnswer := string(answer)
		fmt.Print(strAnswer)
		if strAnswer == "l" || strAnswer == "L" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("You chose login.")); err != nil {
				panic(err)
			}
			user := s.loginUser(conn)
			return user
		} else if strAnswer == "r" || strAnswer == "R" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("You chose register.")); err != nil {
				panic(err)
			}
			user := s.registerUser(conn)
			return user
		} else {
			if err := conn.WriteMessage(websocket.TextMessage,
				[]byte("Invalid command. Write l for login or r for register.")); err != nil {
				panic(err)
			}
		}
	}

}

// ------------------------------------- Register -------------------------------------

func (s *Server) registerUser(conn *websocket.Conn) user.User {
	user := s.readCredentials(conn)
	//user := user.NewUser(string(userName), string(pass))
	s.db.Create(&user)
	return user
}

// ------------------------------------- Login -------------------------------------

func (s *Server) isUserInDataBase(conn *websocket.Conn, possUserName string) bool {
	u := user.User{}
	log.Print("The name of the possUser is: ")
	log.Print(possUserName)
	//name := strings.Trim(string(possUserName), "\n")

	s.db.Where("name = ?", possUserName).First(&u)
	log.Print("The name of the found user is: ")
	log.Print(u.Name)
	if u.Name != possUserName {
		return false
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(u.Password)); err != nil {
		panic(err)
	}
	_, res, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	if string(res) == "Not in DB" {
		return false
	}
	return true
}

func (s *Server) loginUser(conn *websocket.Conn) user.User {
	count := 0
	for {
		userName := s.getName(conn)
		isUserInDBalready := s.isUserInDataBase(conn, userName)
		if isUserInDBalready {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("Logged in!")); err != nil {
				panic(err)
			}
			u := user.NewUser(userName, "")
			return u
		} else {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("Wrong username or password!")); err != nil {
				panic(err)
			}
		}
		count++
	}

}

// ------------------------------------- Helper function -------------------------------------

func (s *Server) readCredentials(conn *websocket.Conn) user.User {
	var u user.User
	if err := conn.ReadJSON(&u); err != nil {
		panic(err)
	}
	return u
}

func (s *Server) getName(conn *websocket.Conn) string {
	_, userName, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	return string(userName)
}
