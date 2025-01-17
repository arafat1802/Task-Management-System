package controllers

import (
	"log"
	"net/http"

	"github.com/arafat1802/Task-Management-System/config"
	"github.com/arafat1802/Task-Management-System/models"
	"github.com/arafat1802/Task-Management-System/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
}
func GetUsers(c *gin.Context) {
	// Call the GetUsers function from models to retrieve all users
	users, err := models.GetUsers(config.DB)
	if err != nil {
		// Log and return error if there's an issue
		log.Printf("Error retrieving users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// If no users found, return an empty array
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No users found"})
		return
	}

	// Return all users as a JSON response
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user schema.User
	// Bind JSON body to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		// Handle error from DB execution
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrpt the Password"})
		return
	}

	_, err = config.DB.Exec(models.CreateUser, user.Username, hashedPassword, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user into the database"})
		return
	}

	// For now, just simulate success
	c.JSON(http.StatusOK, gin.H{
		"message":  "User created successfully!",
		"username": user.Username,
		"password": hashedPassword,
		"email":    user.Email,
	})

}
