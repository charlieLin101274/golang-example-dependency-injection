package main

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

func NewMessage(s string) (Message, error) {
	return Message{s}, nil
}

func NewChatroom(m Message) (Chatroom, error) {
	return Chatroom{m}, nil
}

func NewServe(c Chatroom) (Serve, error) {
	return Serve{c}, nil
}

func (s Serve) Start() {
	fmt.Println("start serve...")
	fmt.Println("show the message: ", s.Chatroom.Message.message)
}

func main() {
	msg, _ := NewMessage("hello world")
	cr, _ := NewChatroom(msg)
	serve, _ := NewServe(cr)

	serve.Start()
}
