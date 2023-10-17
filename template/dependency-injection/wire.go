//go:build wireinject
// +build wireinject

package dependencyinjection

import "github.com/google/wire"

func InitializeServe(s string) Serve {
	wire.Build(NewServe, NewChatroom, NewMessage)
	return Serve{}
}
