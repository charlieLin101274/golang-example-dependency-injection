package dependencyinjection

import (
	"fmt"

	"github.com/charlieLin101274/golang-example-dependency-injection/example-1/models"
)

type Message struct {
	message string
}

type Chatroom struct {
	Message Message
}
type Serve struct {
	RoomName models.RoomName
	Chatroom Chatroom
}

func NewMessage(s string) Message {
	return Message{s}
}

func NewChatroom(m Message) Chatroom {
	return Chatroom{m}
}

func NewServe(roomName models.RoomName, c Chatroom) Serve {
	return Serve{RoomName: roomName, Chatroom: c}
}

func (s Serve) Start() {
	fmt.Println("start serve...")
	fmt.Println("show the message: ", s.Chatroom.Message.message)
	fmt.Println("show the roomName: ", s.RoomName)
}
