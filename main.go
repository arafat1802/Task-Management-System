package main

import (
	"example.com/movie-app/db"
	"example.com/movie-app/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
