package controllers

import (
	"net/http"

	"github.com/arafat1802/Task-Management-System/config"
	"github.com/arafat1802/Task-Management-System/models"
	"github.com/arafat1802/Task-Management-System/schema"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func CreateUser(c *gin.Context) {
	var user schema.User
	// Bind JSON body to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(models.CreateUser, user.Username, user.Password, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user into the database"})
		return
	}

	// For now, just simulate success
	c.JSON(http.StatusOK, gin.H{
		"message":  "User created successfully!",
		"username": user.Username,
		"email":    user.Email,
	})

}
