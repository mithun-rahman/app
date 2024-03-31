package sendmessage

import (
	"fmt"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) StartAndListen() {
	for {
		msg := <-s.msgch
		fmt.Printf("Message Received from: %s payload: %s\n", msg.From, msg.Payload)
	}
}

func SendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Mithun",
		Payload: payload,
	}
	msgch <- msg
}

func Main() {
	s := &Server{
		msgch: make(chan Message),
	}
	go s.StartAndListen()

	SendMessageToServer(s.msgch, "hello")

	SendMessageToServer(s.msgch, "welcome")
}
