package server

import (
	"log"
	"strings"

	"github.com/end-date/user"
	"github.com/gorilla/websocket"
)

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
	userName := s.getUserName(conn)
	pass := s.getPass(conn)
	user := user.NewUser(string(userName), string(pass))
	s.db.Create(&user)
	return user
}

// ------------------------------------- Login -------------------------------------

func (s *Server) isUserInDataBase(conn *websocket.Conn, possUser user.User) bool {
	u := user.User{}
	log.Print("The name of the possUser is: ")
	log.Print(possUser.Name)
	name := strings.Trim(string(possUser.Name), "\n")

	s.db.Where("name = ?", name).First(&u)
	log.Print("The name of the found user is: ")
	log.Print(u.Name)
	if u.Name != possUser.Name {
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
		//log.Println(count)
		userName := s.getUserName(conn)
		pass := s.getPass(conn)
		log.Print(userName)
		log.Print(pass)
		user := user.NewUser(userName, pass)
		isUserInDBalready := s.isUserInDataBase(conn, user)
		if isUserInDBalready {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("Logged in!")); err != nil {
				panic(err)
			}
			return user
		} else {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("Wrong username or password!")); err != nil {
				panic(err)
			}
		}
		count++
	}

}

// ------------------------------------- Helper function -------------------------------------

func (s *Server) getUserName(conn *websocket.Conn) string {
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Username: ")); err != nil {
		panic(err)
	}
	_, userName, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	return string(userName)
}

func (s *Server) getPass(conn *websocket.Conn) string {
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Password: ")); err != nil {
		panic(err)
	}

	_, pass, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	return string(pass)
}
