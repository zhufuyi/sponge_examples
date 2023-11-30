// Package main is the grpc gateway server of the application.
package main

import (
	"order_gw/cmd/order_gw/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.Config()
	servers := initial.RegisterServers()
	closes := initial.RegisterClose(servers)

	a := app.New(servers, closes)
	a.Run()
}
