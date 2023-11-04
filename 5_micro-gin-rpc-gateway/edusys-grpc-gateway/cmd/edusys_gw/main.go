// Package main is the grpc gateway server of the application.
package main

import (
	"edusys_gw/cmd/edusys_gw/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.Config()
	servers := initial.RegisterServers()
	closes := initial.RegisterClose(servers)

	a := app.New(servers, closes)
	a.Run()
}
