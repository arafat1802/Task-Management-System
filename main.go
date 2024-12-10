package main

import (
	"fmt"
	"github.com/arafat1802/Task-Management-System/config"
	"github.com/arafat1802/Task-Management-System/models"
)

func main() {
	fmt.Println(models.createTableSQL)
	config.ConnectDB()
}
