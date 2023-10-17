package dependencyinjection

import "fmt"

type Message struct {
	message string
}

type Chatroom struct {
	Message Message
}
type Serve struct {
	Chatroom Chatroom
}

func NewMessage(s string) Message {
	return Message{s}
}

func NewChatroom(m Message) Chatroom {
	return Chatroom{m}
}

func NewServe(c Chatroom) Serve {
	return Serve{c}
}

func (s Serve) Start() {
	fmt.Println("start serve...")
	fmt.Println("show the message: ", s.Chatroom.Message.message)
}
