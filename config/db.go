package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/arafat1802/Task-Management-System/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	fmt.Println(dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// createTableSQL := `
	// 	CREATE TABLE IF NOT EXISTS "users" (
	// 		id SERIAL PRIMARY KEY,
	// 		username VARCHAR(50) NOT NULL,
	// 		email VARCHAR(100) NOT NULL UNIQUE,
	// 		password VARCHAR(100) NOT NULL,
	// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// 	);`

	_, err = db.Exec(models.createTableSQL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table user created successfully!")
}
