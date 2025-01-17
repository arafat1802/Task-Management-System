package main

import (
	routes "github.com/arafat1802/Task-Management-System/api"
	"github.com/arafat1802/Task-Management-System/config"
	"github.com/arafat1802/Task-Management-System/initializers"
)

func Init(){
	initializers.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	Init()
	r := routes.SetupRouter()
	r.Run(":8080")
}
