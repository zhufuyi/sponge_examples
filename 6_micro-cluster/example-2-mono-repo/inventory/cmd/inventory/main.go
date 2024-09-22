// Package main is the grpc server of the application.
package main

import (
	"github.com/zhufuyi/sponge/pkg/app"

	"eshop/inventory/cmd/inventory/initial"
)

func main() {
	initial.InitApp()
	services := initial.CreateServices()
	closes := initial.Close(services)

	a := app.New(services, closes)
	a.Run()
}
