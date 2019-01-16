package server

// This will be the server doing the backend work
// The main work of the server will be to handle user requests
type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s *Server) registerClient(user *User) {
	// user.
}

// Here the server will be spawned and will wait for new users
func (s *Server) executeServer() {

}
