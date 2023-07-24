package handlers

import (
	"fmt"
	"net/http"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/dtos"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var eventRequest dtos.EventDTO

	if err := c.ShouldBindJSON(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := dtos.ValidateEventDTO(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.eventService.CreateEvent(eventRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (h *EventHandler) AddParticipant(c *gin.Context) {
	eventID := c.Param("eventId")

	eventUUID, err := uuid.Parse(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format"})
		return
	}

	var participantRequest dtos.ParticipantDTO
	if err := c.ShouldBindJSON(&participantRequest); err != nil {
		fmt.Println("err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	participantRequest.EventId = eventUUID

	event, err := h.eventService.AddParticipant(participantRequest)
	if err != nil {
		fmt.Println("err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, event)
}
