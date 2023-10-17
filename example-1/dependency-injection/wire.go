//go:build wireinject
// +build wireinject

package dependencyinjection

import (
	"github.com/charlieLin101274/golang-example-dependency-injection/example-1/config"
	"github.com/google/wire"
)

func provideServe(cfg config.Config, c Chatroom) Serve {
	return NewServe(cfg.RoomName, c)
}

func InitializeServe(s string, cfg config.Config) Serve {
	wire.Build(provideServe, NewChatroom, NewMessage)
	return Serve{}
}
