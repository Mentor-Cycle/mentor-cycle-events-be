package routes

import (
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, eventHandler *handlers.EventHandler) {
	r.GET("/events", eventHandler.FindAllEvents)
	r.POST("/events", eventHandler.CreateEvent)
	r.PUT("/events/:eventId", eventHandler.AddParticipant)
}
