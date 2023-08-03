package routes

import (
	"fmt"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the request header or query parameters
		token, _ := c.Cookie("token")

		fmt.Println("token: ", token)
		// Your token validation logic goes here
		// For example, you can check if the token is valid or not

		// If the token is valid, proceed to the next handler
		c.Next()
		// If the token is invalid, you can stop the request and return an error
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
}

func SetupRoutes(r *gin.Engine, eventHandler *handlers.EventHandler) {
	r.GET("/events", eventHandler.FindAllEvents)
	r.POST("/events", AuthMiddleware(), eventHandler.CreateEvent)
	r.PUT("/events/:eventId", eventHandler.AddParticipant)
}
