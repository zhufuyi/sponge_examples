// Package main is the grpc gateway server of the application.
package main

import (
	"github.com/zhufuyi/sponge/pkg/app"

	"eshop/eshop_gw/cmd/eshop_gw/initial"
)

func main() {
	initial.InitApp()
	services := initial.CreateServices()
	closes := initial.Close(services)

	a := app.New(services, closes)
	a.Run()
}
