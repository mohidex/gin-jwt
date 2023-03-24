package main

import (
	"github.com/mohidex/identity-service/server"
	"github.com/mohidex/identity-service/settings"
)

func main() {
	settings.ConnectDB()
	server.Init()
}
