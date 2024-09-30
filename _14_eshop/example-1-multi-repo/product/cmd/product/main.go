// Package main is the http and grpc server of the application.
package main

import (
	"github.com/zhufuyi/sponge/pkg/app"

	"product/cmd/product/initial"
)

func main() {
	initial.InitApp()
	services := initial.CreateServices()
	closes := initial.Close(services)

	a := app.New(services, closes)
	a.Run()
}
