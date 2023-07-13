package handlers

import (
	"net/http"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/services"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService services.EventService
}

func NewEventHandler(eventService services.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

func (h *EventHandler) FindAllEvents(c *gin.Context) {
	events, err := h.eventService.FindAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
