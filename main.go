package main

import (
	"github.com/mohidex/identity-service/models"
	"github.com/mohidex/identity-service/server"
	"github.com/mohidex/identity-service/settings"
)

func main() {
	settings.ConnectDB()
	settings.GetDB().AutoMigrate(&models.User{})
	server.Init()
}
