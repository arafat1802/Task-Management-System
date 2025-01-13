package models

import (
	"database/sql"
	"log"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetUsers retrieves all users from the database
func GetUsers(db *sql.DB) ([]User, error) {
	// Query to get all users
	query := `SELECT id, username, password, email FROM users`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Slice to hold all users
	var users []User

	// Loop through the rows and scan each user into the users slice
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Check for any row iteration error
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	// Return the slice of users
	return users, nil
}
