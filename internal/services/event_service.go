package services

import (
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/models"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/repositories"
)

type EventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository repositories.EventRepository) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}

func (s *EventService) FindAllEvents() ([]models.Event, error) {
	return s.eventRepository.FindAll()
}
