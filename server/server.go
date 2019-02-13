package server

import (
	"github.com/jinzhu/gorm"

	"github.com/end-date/user"
)

// Server will be the server doing the backend work
// The main work of the server will be to handle user requests
type Server struct {
	db *gorm.DB
}

// NewServer creates a new server instance
func NewServer(db *gorm.DB) Server {
	return Server{
		db: db,
	}
}

func (s *Server) registerClient(user *user.User) {
	// user.
}

// Here the server will be spawned and will wait for new users
func (s *Server) executeServer() {

}
