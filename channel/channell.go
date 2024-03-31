package channel

import "fmt"

type Server struct {
	Users  map[string]string
	UserCh chan string
}

func NewServer() *Server {
	return &Server{
		Users:  make(map[string]string),
		UserCh: make(chan string),
	}
}

func (s *Server) Start() {
	go s.Loop()
}

func (s *Server) Loop() {
	for {
		user := <-s.UserCh
		s.Users[user] = user
		fmt.Println("adding new user ", user)
	}
}
