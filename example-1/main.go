package main

import (
	"github.com/charlieLin101274/golang-example-dependency-injection/example-1/config"
	dependencyinjection "github.com/charlieLin101274/golang-example-dependency-injection/example-1/dependency-injection"
)

// func main() {
// 	msg := dependencyinjection.NewMessage("hello world")
// 	cr := dependencyinjection.NewChatroom(msg)
// 	serve := dependencyinjection.NewServe(cr)

// 	serve.Start()
// }

func main() {
	cfg := config.Config{RoomName: "TestRoom"}
	serve := dependencyinjection.InitializeServe("hello world", cfg)

	serve.Start()
}
