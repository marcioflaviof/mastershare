package main

import (
	"project/database"
	"project/service"
)

func main() {
	database.CreateClient()
	service.StartServer()
}
