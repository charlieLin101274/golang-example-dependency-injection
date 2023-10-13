package main

import dependencyinjection "github.com/charlieLin101274/golang-example-dependency-injection/dependency-injection"

// func main() {
// 	msg := dependencyinjection.NewMessage("hello world")
// 	cr := dependencyinjection.NewChatroom(msg)
// 	serve := dependencyinjection.NewServe(cr)

// 	serve.Start()
// }

func main() {
	serve := dependencyinjection.InitializeServe("hello world")

	serve.Start()
}
