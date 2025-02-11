package login

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/arafat1802/Task-Management-System/config"
    "github.com/arafat1802/Task-Management-System/schema"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key_here") // Replace with a secure secret key

func LoginUser(c *gin.Context) {
    var user schema.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB
    var hashedPassword string
    query := "SELECT password FROM users WHERE username = $1"
    err := db.QueryRow(query, user.Username).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query user"})
        }
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
    if err != nil { 
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 3 days
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": tokenString})
}
