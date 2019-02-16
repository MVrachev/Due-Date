package server

import (
	"sync"

	"github.com/jinzhu/gorm"
)

// Server will be the server doing the backend work
// The main work of the server will be to handle user requests
type Server struct {
	db    *gorm.DB
	mutex sync.Mutex
}

// NewServer creates a new server instance
func NewServer(db *gorm.DB) Server {
	return Server{
		db: db,
	}
}

func (s *Server) Close() {
	s.db.Close()
}
