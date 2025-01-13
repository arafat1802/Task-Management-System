package main

import (
	routes "github.com/arafat1802/Task-Management-System/api"
	"github.com/arafat1802/Task-Management-System/config"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
