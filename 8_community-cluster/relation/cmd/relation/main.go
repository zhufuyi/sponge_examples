// Package main is the grpc server of the application.
package main

import (
	"relation/cmd/relation/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.Config()
	servers := initial.RegisterServers()
	closes := initial.RegisterClose(servers)

	a := app.New(servers, closes)
	a.Run()
}
