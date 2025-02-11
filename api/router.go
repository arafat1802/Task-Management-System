package routes

import (
	"net/http"
	"os"

	"github.com/arafat1802/Task-Management-System/controllers"
	"github.com/arafat1802/Task-Management-System/internal/database/login"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Assuming the AuthMiddleware function is defined in another file or package
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSecret := os.Getenv("jwtSecret")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Token is valid, continue to the next handler
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/createUser", controllers.CreateUser)
	r.POST("/login", login.LoginUser)
	r.GET("/hello", AuthMiddleware(), controllers.HelloWorld)

	return r
}
