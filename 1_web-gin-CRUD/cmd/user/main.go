package main

import (
	"user/cmd/user/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

// @title user api docs
// @description http server api docs
// @schemes http https
// @version v0.0.0
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer your-jwt-token" to Value
func main() {
	initial.Config()
	servers := initial.RegisterServers()
	closes := initial.RegisterClose(servers)

	a := app.New(servers, closes)
	a.Run()
}
